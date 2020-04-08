package linkerd_translator_test

import (
	"context"
	"time"

	"github.com/solo-io/service-mesh-hub/test/fakes"

	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-multierror"
	linkerd_config "github.com/linkerd/linkerd2/controller/gen/apis/serviceprofile/v1alpha2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	networking_types "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1/types"
	"k8s.io/apimachinery/pkg/api/resource"

	core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	discovery_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	discovery_types "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1/types"
	"github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"
	linkerd_networking "github.com/solo-io/service-mesh-hub/pkg/clients/linkerd/v1alpha2"
	mock_core "github.com/solo-io/service-mesh-hub/pkg/clients/zephyr/discovery/mocks"
	mock_mc_manager "github.com/solo-io/service-mesh-hub/services/common/multicluster/manager/mocks"
	linkerd_translator "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/routing/traffic-policy-translator/linkerd-translator"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("LinkerdTranslator", func() {
	var (
		ctrl                           *gomock.Controller
		linkerdTrafficPolicyTranslator linkerd_translator.LinkerdTranslator
		ctx                            context.Context
		mockDynamicClientGetter        *mock_mc_manager.MockDynamicClientGetter
		mockMeshClient                 *mock_core.MockMeshClient

		clusterName       = "clusterName"
		meshObjKey        = client.ObjectKey{Name: "mesh-name", Namespace: "mesh-namespace"}
		meshServiceObjKey = client.ObjectKey{Name: "mesh-service-name", Namespace: "mesh-service-namespace"}
		kubeServiceObjKey = client.ObjectKey{Name: "kube-service-name", Namespace: "kube-service-namespace"}
		meshService       = &discovery_v1alpha1.MeshService{
			ObjectMeta: v1.ObjectMeta{
				Name:        meshServiceObjKey.Name,
				Namespace:   meshServiceObjKey.Namespace,
				ClusterName: clusterName,
			},
			Spec: discovery_types.MeshServiceSpec{
				Mesh: &core_types.ResourceRef{
					Name:      meshObjKey.Name,
					Namespace: meshObjKey.Namespace,
				},
				KubeService: &discovery_types.MeshServiceSpec_KubeService{
					Ref: &core_types.ResourceRef{
						Name:      kubeServiceObjKey.Name,
						Namespace: kubeServiceObjKey.Namespace,
						Cluster:   clusterName,
					},
					Ports: []*discovery_types.MeshServiceSpec_KubeService_KubeServicePort{
						{
							Port: 9080,
							Name: "http",
						},
					},
				},
			},
		}
		mesh = &discovery_v1alpha1.Mesh{
			Spec: discovery_types.MeshSpec{
				Cluster: &core_types.ResourceRef{
					Name: clusterName,
				},
				MeshType: &discovery_types.MeshSpec_Linkerd{
					Linkerd: &discovery_types.MeshSpec_LinkerdMesh{
						ClusterDomain: "cluster.domain",
					},
				},
			},
		}
		serviceProfileClient linkerd_networking.ServiceProfileClient
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.TODO()
		mockDynamicClientGetter = mock_mc_manager.NewMockDynamicClientGetter(ctrl)
		mockMeshClient = mock_core.NewMockMeshClient(ctrl)
		serviceProfileClient = linkerd_networking.NewServiceProfileClient(fakes.InMemoryClient())
		linkerdTrafficPolicyTranslator = linkerd_translator.NewLinkerdTrafficPolicyTranslator(
			mockDynamicClientGetter,
			mockMeshClient,
			func(client client.Client) linkerd_networking.ServiceProfileClient {
				return serviceProfileClient
			},
		)
		mockMeshClient.EXPECT().Get(ctx, meshObjKey).Return(mesh, nil)
		mockDynamicClientGetter.EXPECT().GetClientForCluster(clusterName).Return(nil, nil)

	})
	AfterEach(func() {
		ctrl.Finish()
	})
	Context("no relevant config provided", func() {

		trafficPolicy := []*v1alpha1.TrafficPolicy{{
			Spec: networking_types.TrafficPolicySpec{
				FaultInjection: &networking_types.TrafficPolicySpec_FaultInjection{
					Percentage: 100,
				},
				HeaderManipulation: &networking_types.TrafficPolicySpec_HeaderManipulation{
					AppendResponseHeaders: map[string]string{"foo": "bar"},
				},
			}},
		}

		It("does not create a service profile", func() {
			translatorError := linkerdTrafficPolicyTranslator.TranslateTrafficPolicy(
				ctx,
				meshService,
				mesh,
				trafficPolicy)
			Expect(translatorError).To(BeNil())

			serviceProfiles, err := serviceProfileClient.List(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(serviceProfiles.Items).To(BeEmpty())
		})
	})

	Context("basic traffic policy", func() {
		trafficPolicy := []*v1alpha1.TrafficPolicy{
			{
				Spec: networking_types.TrafficPolicySpec{
					HttpRequestMatchers: []*networking_types.TrafficPolicySpec_HttpMatcher{
						{}, // one default matcher
					},
				},
			},
		}

		It("creates sp with the name and namespace matching the kube DNS name and namespace of the backing service, respectively", func() {
			translatorError := linkerdTrafficPolicyTranslator.TranslateTrafficPolicy(
				ctx,
				meshService,
				mesh,
				trafficPolicy)
			Expect(translatorError).To(BeNil())

			serviceProfiles, err := serviceProfileClient.List(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(serviceProfiles.Items).To(HaveLen(1))

			sp := &serviceProfiles.Items[0]

			Expect(sp.Name).To(Equal("kube-service-name.kube-service-namespace.cluster.domain"))
			Expect(sp.Namespace).To(Equal(kubeServiceObjKey.Namespace))
		})
	})

	Context("prefix matcher provided", func() {

		trafficPolicy := []*v1alpha1.TrafficPolicy{
			{
				ObjectMeta: v1.ObjectMeta{Namespace: "ns", Name: "tp"},
				Spec: networking_types.TrafficPolicySpec{
					HttpRequestMatchers: []*networking_types.TrafficPolicySpec_HttpMatcher{
						{
							PathSpecifier: &networking_types.TrafficPolicySpec_HttpMatcher_Prefix{
								Prefix: "/prefix/",
							},
							Method: &networking_types.TrafficPolicySpec_HttpMethod{Method: core_types.HttpMethodValue_GET},
						},
					},
				},
			},
		}

		It("creates sp with the paths converted to regex", func() {
			translatorError := linkerdTrafficPolicyTranslator.TranslateTrafficPolicy(
				ctx,
				meshService,
				mesh,
				trafficPolicy)
			Expect(translatorError).To(BeNil())

			serviceProfiles, err := serviceProfileClient.List(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(serviceProfiles.Items).To(HaveLen(1))

			sp := &serviceProfiles.Items[0]

			Expect(sp.Spec.Routes).To(Equal([]*linkerd_config.RouteSpec{
				{
					Name: "tp.ns",
					Condition: &linkerd_config.RequestMatch{
						Any: []*linkerd_config.RequestMatch{
							{
								PathRegex: "/prefix/.*",
								Method:    "GET",
							},
						},
					},
				},
			}))
		})
	})

	Context("traffic shift provided", func() {

		trafficPolicy := []*v1alpha1.TrafficPolicy{
			{
				ObjectMeta: v1.ObjectMeta{Namespace: "ns", Name: "tp"},
				Spec: networking_types.TrafficPolicySpec{
					TrafficShift: &networking_types.TrafficPolicySpec_MultiDestination{
						Destinations: []*networking_types.TrafficPolicySpec_MultiDestination_WeightedDestination{
							{
								Destination: &core_types.ResourceRef{Name: "foo-svc", Namespace: "foo-ns"},
								Weight:      5,
								Port:        1234,
							},
							{
								Destination: &core_types.ResourceRef{Name: "bar-svc", Namespace: "bar-ns", Cluster: "bar-cluster"},
								Weight:      15,
								Port:        3456,
							},
						},
					},
				},
			},
		}

		It("creates sp with the corresponding traffic split", func() {

			translatorError := linkerdTrafficPolicyTranslator.TranslateTrafficPolicy(
				ctx,
				meshService,
				mesh,
				trafficPolicy)
			Expect(translatorError).To(BeNil())

			serviceProfiles, err := serviceProfileClient.List(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(serviceProfiles.Items).To(HaveLen(1))

			sp := &serviceProfiles.Items[0]

			Expect(sp.Spec.DstOverrides).To(Equal([]*linkerd_config.WeightedDst{
				{
					Authority: "foo-svc.foo-ns.svc.cluster.domain:1234",
					Weight:    resource.MustParse("250m"),
				},
				{
					Authority: "bar-svc.bar-ns.svc.cluster.domain:3456",
					Weight:    resource.MustParse("750m"),
				},
			}))

		})
	})

	Context("timeout provided", func() {

		trafficPolicy := []*v1alpha1.TrafficPolicy{
			{
				ObjectMeta: v1.ObjectMeta{Namespace: "ns", Name: "tp"},
				Spec: networking_types.TrafficPolicySpec{
					RequestTimeout: types.DurationProto(time.Minute),
				},
			},
		}

		It("creates sp with the corresponding timeout", func() {
			translatorError := linkerdTrafficPolicyTranslator.TranslateTrafficPolicy(
				ctx,
				meshService,
				mesh,
				trafficPolicy)
			Expect(translatorError).To(BeNil())

			serviceProfiles, err := serviceProfileClient.List(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(serviceProfiles.Items).To(HaveLen(1))

			sp := &serviceProfiles.Items[0]

			Expect(sp.Spec.Routes).To(Equal([]*linkerd_config.RouteSpec{
				{
					Name: "tp.ns",
					Condition: &linkerd_config.RequestMatch{
						Any: []*linkerd_config.RequestMatch{
							{
								PathRegex: "/.*",
								Method:    "",
							},
						},
					},
					Timeout: time.Minute.String(),
				},
			}))
		})
	})

	Context("multiple policies defined", func() {

		trafficPolicy := []*v1alpha1.TrafficPolicy{
			{
				ObjectMeta: v1.ObjectMeta{Namespace: "ns", Name: "tp1"},
				Spec: networking_types.TrafficPolicySpec{
					HttpRequestMatchers: []*networking_types.TrafficPolicySpec_HttpMatcher{
						{
							PathSpecifier: &networking_types.TrafficPolicySpec_HttpMatcher_Prefix{
								Prefix: "/short",
							},
							Method: &networking_types.TrafficPolicySpec_HttpMethod{Method: core_types.HttpMethodValue_GET},
						},
					},
				},
			},
			{
				ObjectMeta: v1.ObjectMeta{Namespace: "ns", Name: "tp2"},
				Spec: networking_types.TrafficPolicySpec{
					HttpRequestMatchers: []*networking_types.TrafficPolicySpec_HttpMatcher{
						{
							PathSpecifier: &networking_types.TrafficPolicySpec_HttpMatcher_Prefix{
								Prefix: "/longer",
							},
							Method: &networking_types.TrafficPolicySpec_HttpMethod{Method: core_types.HttpMethodValue_GET},
						},
					},
				},
			},
		}

		It("sorts the routes in the sp byt the length of the first path specifier", func() {
			translatorError := linkerdTrafficPolicyTranslator.TranslateTrafficPolicy(
				ctx,
				meshService,
				mesh,
				trafficPolicy)
			Expect(translatorError).To(BeNil())

			serviceProfiles, err := serviceProfileClient.List(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(serviceProfiles.Items).To(HaveLen(1))

			sp := &serviceProfiles.Items[0]

			Expect(sp.Spec.Routes).To(Equal([]*linkerd_config.RouteSpec{
				{
					Name: "tp2.ns",
					Condition: &linkerd_config.RequestMatch{
						Any: []*linkerd_config.RequestMatch{
							{
								PathRegex: "/longer.*",
								Method:    "GET",
							},
						},
					},
				},
				{
					Name: "tp1.ns",
					Condition: &linkerd_config.RequestMatch{
						Any: []*linkerd_config.RequestMatch{
							{
								PathRegex: "/short.*",
								Method:    "GET",
							},
						},
					},
				},
			}))
		})
	})
	Context("subsets defined in destination", func() {

		dest := &networking_types.TrafficPolicySpec_MultiDestination_WeightedDestination{Subset: map[string]string{}}
		trafficPolicy := []*v1alpha1.TrafficPolicy{
			{
				ObjectMeta: v1.ObjectMeta{Namespace: "ns", Name: "tp"},
				Spec: networking_types.TrafficPolicySpec{
					TrafficShift: &networking_types.TrafficPolicySpec_MultiDestination{
						Destinations: []*networking_types.TrafficPolicySpec_MultiDestination_WeightedDestination{dest},
					},
				},
			},
		}

		It("returns a translator error", func() {
			translatorError := linkerdTrafficPolicyTranslator.TranslateTrafficPolicy(
				ctx,
				meshService,
				mesh,
				trafficPolicy)
			Expect(translatorError).To(Equal(&networking_types.TrafficPolicyStatus_TranslatorError{
				TranslatorId: linkerd_translator.TranslatorId,
				ErrorMessage: multierror.Append(nil, linkerd_translator.SubsetsNotSupportedErr(dest)).Error(),
			}))
		})
	})
	Context("multiple policies defined with traffic shift", func() {

		trafficPolicy := []*v1alpha1.TrafficPolicy{
			{
				ObjectMeta: v1.ObjectMeta{Namespace: "ns", Name: "tp1"},
				Spec: networking_types.TrafficPolicySpec{
					TrafficShift: &networking_types.TrafficPolicySpec_MultiDestination{},
				},
			},
			{
				ObjectMeta: v1.ObjectMeta{Namespace: "ns", Name: "tp2"},
				Spec: networking_types.TrafficPolicySpec{
					TrafficShift: &networking_types.TrafficPolicySpec_MultiDestination{},
				},
			},
		}

		It("returns a translator error", func() {
			translatorError := linkerdTrafficPolicyTranslator.TranslateTrafficPolicy(
				ctx,
				meshService,
				mesh,
				trafficPolicy)
			Expect(translatorError).To(Equal(&networking_types.TrafficPolicyStatus_TranslatorError{
				TranslatorId: linkerd_translator.TranslatorId,
				ErrorMessage: multierror.Append(nil, linkerd_translator.TrafficShiftRedefinedErr(meshService, []core_types.ResourceRef{
					{Namespace: "ns", Name: "tp1"},
					{Namespace: "ns", Name: "tp2"},
				})).Error(),
			}))
		})
	})

})
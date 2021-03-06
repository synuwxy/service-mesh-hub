// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wire

import (
	"context"

	v1alpha1_3 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	"github.com/solo-io/service-mesh-hub/pkg/api/istio/networking/v1alpha3"
	"github.com/solo-io/service-mesh-hub/pkg/api/istio/security/v1beta1"
	v1_2 "github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/apps/v1"
	v1 "github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/core/v1"
	"github.com/solo-io/service-mesh-hub/pkg/api/linkerd/v1alpha2"
	v1alpha1_2 "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"
	"github.com/solo-io/service-mesh-hub/pkg/api/security.zephyr.solo.io/v1alpha1"
	v1alpha1_4 "github.com/solo-io/service-mesh-hub/pkg/api/smi/split/v1alpha1"
	"github.com/solo-io/service-mesh-hub/pkg/csr/certgen"
	"github.com/solo-io/service-mesh-hub/pkg/filesystem/files"
	"github.com/solo-io/service-mesh-hub/pkg/kube/kubeconfig"
	"github.com/solo-io/service-mesh-hub/pkg/kube/selection"
	mc_wire "github.com/solo-io/service-mesh-hub/services/common/compute-target/wire"
	csr_generator "github.com/solo-io/service-mesh-hub/services/csr-agent/pkg/csr-generator"
	access_policy_enforcer "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/access/access-control-enforcer"
	istio_enforcer "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/access/access-control-enforcer/istio-enforcer"
	acp_translator "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/access/access-control-policy-translator"
	istio_translator2 "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/access/access-control-policy-translator/istio-translator"
	"github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/federation/decider"
	"github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/federation/decider/strategies"
	"github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/federation/dns"
	"github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/federation/resolver"
	istio_federation "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/federation/resolver/meshes/istio"
	networking_multicluster "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/multicluster"
	controller_factories "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/multicluster/controllers"
	traffic_policy_translator "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/routing/traffic-policy-translator"
	istio_translator "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/routing/traffic-policy-translator/istio-translator"
	linkerd_translator "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/routing/traffic-policy-translator/linkerd-translator"
	"github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/routing/traffic-policy-translator/preprocess"
	cert_manager "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/security/cert-manager"
	cert_signer "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/security/cert-signer"
	vm_validation "github.com/solo-io/service-mesh-hub/services/mesh-networking/pkg/validation"
)

// Injectors from wire.go:

func InitializeMeshNetworking(ctx context.Context) (MeshNetworkingContext, error) {
	config, err := mc_wire.LocalKubeConfigProvider()
	if err != nil {
		return MeshNetworkingContext{}, err
	}
	asyncManager, err := mc_wire.LocalManagerProvider(ctx, config)
	if err != nil {
		return MeshNetworkingContext{}, err
	}
	fileReader := files.NewDefaultFileReader()
	converter := kubeconfig.NewConverter(fileReader)
	asyncManagerController := mc_wire.KubeClusterCredentialsHandlerProvider(converter)
	awsCredsHandler := NewNetworkingAwsCredsHandler()
	v := ComputeTargetCredentialsHandlersProvider(asyncManagerController, awsCredsHandler)
	asyncManagerStartOptionsFunc := mc_wire.LocalManagerStarterProvider(v)
	multiClusterDependencies := mc_wire.MulticlusterDependenciesProvider(ctx, asyncManager, asyncManagerController, asyncManagerStartOptionsFunc)
	virtualMeshCSRControllerFactory := controller_factories.NewVirtualMeshCSRControllerFactory()
	controllerFactories := NewControllerFactories(virtualMeshCSRControllerFactory)
	virtualMeshCertificateSigningRequestClientFactory := v1alpha1.VirtualMeshCertificateSigningRequestClientFactoryProvider()
	clientFactories := NewClientFactories(virtualMeshCertificateSigningRequestClientFactory)
	client := mc_wire.DynamicClientProvider(asyncManager)
	secretClient := v1.SecretClientProvider(client)
	virtualMeshClient := v1alpha1_2.VirtualMeshClientProvider(client)
	rootCertGenerator := certgen.NewRootCertGenerator()
	virtualMeshCertClient := cert_signer.NewVirtualMeshCertClient(secretClient, virtualMeshClient, rootCertGenerator)
	signer := certgen.NewSigner()
	virtualMeshCSRDataSourceFactory := csr_generator.NewVirtualMeshCSRDataSourceFactory()
	asyncManagerHandler, err := networking_multicluster.NewMeshNetworkingClusterHandler(asyncManager, controllerFactories, clientFactories, virtualMeshCertClient, signer, virtualMeshCSRDataSourceFactory)
	if err != nil {
		return MeshNetworkingContext{}, err
	}
	meshServiceClient := v1alpha1_3.MeshServiceClientProvider(client)
	meshWorkloadClient := v1alpha1_3.MeshWorkloadClientProvider(client)
	deploymentClientFactory := v1_2.DeploymentClientFactoryProvider()
	dynamicClientGetter := mc_wire.DynamicClientGetterProvider(asyncManagerController)
	resourceSelector := selection.NewResourceSelector(meshServiceClient, meshWorkloadClient, deploymentClientFactory, dynamicClientGetter)
	meshClient := v1alpha1_3.MeshClientProvider(client)
	trafficPolicyClient := v1alpha1_2.TrafficPolicyClientProvider(client)
	trafficPolicyMerger := preprocess.NewTrafficPolicyMerger(resourceSelector, meshClient, trafficPolicyClient)
	trafficPolicyValidator := preprocess.NewTrafficPolicyValidator(meshServiceClient, resourceSelector)
	trafficPolicyPreprocessor := preprocess.NewTrafficPolicyPreprocessor(resourceSelector, trafficPolicyMerger, trafficPolicyValidator)
	virtualServiceClientFactory := v1alpha3.VirtualServiceClientFactoryProvider()
	destinationRuleClientFactory := v1alpha3.DestinationRuleClientFactoryProvider()
	istioTranslator := istio_translator.NewIstioTrafficPolicyTranslator(dynamicClientGetter, meshClient, meshServiceClient, resourceSelector, virtualServiceClientFactory, destinationRuleClientFactory)
	serviceProfileClientFactory := v1alpha2.ServiceProfileClientFactoryProvider()
	trafficSplitClientFactory := v1alpha1_4.TrafficSplitClientFactoryProvider()
	linkerdTranslator := linkerd_translator.NewLinkerdTrafficPolicyTranslator(dynamicClientGetter, meshClient, serviceProfileClientFactory, trafficSplitClientFactory)
	v2 := TrafficPolicyMeshTranslatorsProvider(istioTranslator, linkerdTranslator)
	trafficPolicyEventWatcher := LocalTrafficPolicyEventWatcherProvider(asyncManager)
	meshServiceEventWatcher := LocalMeshServiceEventWatcherProvider(asyncManager)
	trafficPolicyTranslatorLoop := traffic_policy_translator.NewTrafficPolicyTranslatorLoop(trafficPolicyPreprocessor, v2, meshClient, meshServiceClient, trafficPolicyClient, trafficPolicyEventWatcher, meshServiceEventWatcher)
	meshWorkloadEventWatcher := LocalMeshWorkloadEventWatcherProvider(asyncManager)
	virtualMeshEventWatcher := controller_factories.NewLocalVirtualMeshEventWatcher(asyncManager)
	virtualMeshFinder := vm_validation.NewVirtualMeshFinder(meshClient)
	meshNetworkingSnapshotValidator := vm_validation.NewVirtualMeshValidator(virtualMeshFinder, virtualMeshClient)
	istioCertConfigProducer := cert_manager.NewIstioCertConfigProducer()
	virtualMeshCertificateManager := cert_manager.NewVirtualMeshCsrProcessor(dynamicClientGetter, meshClient, virtualMeshFinder, virtualMeshCertificateSigningRequestClientFactory, istioCertConfigProducer)
	vmcsrSnapshotListener := cert_manager.NewVMCSRSnapshotListener(virtualMeshCertificateManager, virtualMeshClient)
	federationStrategyChooser := strategies.NewFederationStrategyChooser()
	federationDecider := decider.NewFederationDecider(meshServiceClient, meshClient, virtualMeshClient, federationStrategyChooser)
	federationDeciderSnapshotListener := decider.NewFederationSnapshotListener(federationDecider)
	meshNetworkingSnapshotContext := MeshNetworkingSnapshotContextProvider(meshWorkloadEventWatcher, meshServiceEventWatcher, virtualMeshEventWatcher, meshNetworkingSnapshotValidator, vmcsrSnapshotListener, federationDeciderSnapshotListener)
	accessControlPolicyEventWatcher := LocalAccessControlPolicyEventWatcherProvider(asyncManager)
	accessControlPolicyClient := v1alpha1_2.AccessControlPolicyClientProvider(client)
	authorizationPolicyClientFactory := v1beta1.AuthorizationPolicyClientFactoryProvider()
	istio_translatorIstioTranslator := istio_translator2.NewIstioTranslator(meshClient, dynamicClientGetter, authorizationPolicyClientFactory)
	v3 := AccessControlPolicyMeshTranslatorsProvider(istio_translatorIstioTranslator)
	acpTranslatorLoop := acp_translator.NewAcpTranslatorLoop(accessControlPolicyEventWatcher, meshServiceEventWatcher, meshClient, accessControlPolicyClient, resourceSelector, v3)
	istioEnforcer := istio_enforcer.NewIstioEnforcer(dynamicClientGetter, authorizationPolicyClientFactory)
	v4 := GlobalAccessControlPolicyMeshEnforcersProvider(istioEnforcer)
	accessPolicyEnforcerLoop := access_policy_enforcer.NewEnforcerLoop(virtualMeshEventWatcher, virtualMeshClient, meshClient, v4)
	gatewayClientFactory := v1alpha3.GatewayClientFactoryProvider()
	envoyFilterClientFactory := v1alpha3.EnvoyFilterClientFactoryProvider()
	serviceEntryClientFactory := v1alpha3.ServiceEntryClientFactoryProvider()
	serviceClientFactory := v1.ServiceClientFactoryProvider()
	configMapClient := v1.ConfigMapClientProvider(client)
	ipAssigner := dns.NewIpAssigner(configMapClient)
	podClientFactory := v1.PodClientFactoryProvider()
	nodeClientFactory := v1.NodeClientFactoryProvider()
	externalAccessPointGetter := dns.NewExternalAccessPointGetter(dynamicClientGetter, podClientFactory, nodeClientFactory)
	istioFederationClient := istio_federation.NewIstioFederationClient(dynamicClientGetter, meshClient, gatewayClientFactory, envoyFilterClientFactory, destinationRuleClientFactory, serviceEntryClientFactory, serviceClientFactory, ipAssigner, externalAccessPointGetter)
	perMeshFederationClients := resolver.NewPerMeshFederationClients(istioFederationClient)
	federationResolver := resolver.NewFederationResolver(meshClient, meshWorkloadClient, meshServiceClient, virtualMeshClient, perMeshFederationClients, meshServiceEventWatcher)
	meshNetworkingContext := MeshNetworkingContextProvider(multiClusterDependencies, asyncManagerHandler, trafficPolicyTranslatorLoop, meshNetworkingSnapshotContext, acpTranslatorLoop, accessPolicyEnforcerLoop, federationResolver)
	return meshNetworkingContext, nil
}

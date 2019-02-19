// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sync"
	"time"

	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	istio_networking_v1alpha3 "github.com/solo-io/supergloo/pkg/api/external/istio/networking/v1alpha3"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/errutils"
)

var (
	mConfigSnapshotIn  = stats.Int64("config.supergloo.solo.io/snap_emitter/snap_in", "The number of snapshots in", "1")
	mConfigSnapshotOut = stats.Int64("config.supergloo.solo.io/snap_emitter/snap_out", "The number of snapshots out", "1")

	configsnapshotInView = &view.View{
		Name:        "config.supergloo.solo.io_snap_emitter/snap_in",
		Measure:     mConfigSnapshotIn,
		Description: "The number of snapshots updates coming in",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	configsnapshotOutView = &view.View{
		Name:        "config.supergloo.solo.io/snap_emitter/snap_out",
		Measure:     mConfigSnapshotOut,
		Description: "The number of snapshots updates going out",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
)

func init() {
	view.Register(configsnapshotInView, configsnapshotOutView)
}

type ConfigEmitter interface {
	Register() error
	Mesh() MeshClient
	MeshGroup() MeshGroupClient
	Upstream() gloo_solo_io.UpstreamClient
	RoutingRule() RoutingRuleClient
	EncryptionRule() EncryptionRuleClient
	TlsSecret() TlsSecretClient
	DestinationRule() istio_networking_v1alpha3.DestinationRuleClient
	VirtualService() istio_networking_v1alpha3.VirtualServiceClient
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *ConfigSnapshot, <-chan error, error)
}

func NewConfigEmitter(meshClient MeshClient, meshGroupClient MeshGroupClient, upstreamClient gloo_solo_io.UpstreamClient, routingRuleClient RoutingRuleClient, encryptionRuleClient EncryptionRuleClient, tlsSecretClient TlsSecretClient, destinationRuleClient istio_networking_v1alpha3.DestinationRuleClient, virtualServiceClient istio_networking_v1alpha3.VirtualServiceClient) ConfigEmitter {
	return NewConfigEmitterWithEmit(meshClient, meshGroupClient, upstreamClient, routingRuleClient, encryptionRuleClient, tlsSecretClient, destinationRuleClient, virtualServiceClient, make(chan struct{}))
}

func NewConfigEmitterWithEmit(meshClient MeshClient, meshGroupClient MeshGroupClient, upstreamClient gloo_solo_io.UpstreamClient, routingRuleClient RoutingRuleClient, encryptionRuleClient EncryptionRuleClient, tlsSecretClient TlsSecretClient, destinationRuleClient istio_networking_v1alpha3.DestinationRuleClient, virtualServiceClient istio_networking_v1alpha3.VirtualServiceClient, emit <-chan struct{}) ConfigEmitter {
	return &configEmitter{
		mesh:            meshClient,
		meshGroup:       meshGroupClient,
		upstream:        upstreamClient,
		routingRule:     routingRuleClient,
		encryptionRule:  encryptionRuleClient,
		tlsSecret:       tlsSecretClient,
		destinationRule: destinationRuleClient,
		virtualService:  virtualServiceClient,
		forceEmit:       emit,
	}
}

type configEmitter struct {
	forceEmit       <-chan struct{}
	mesh            MeshClient
	meshGroup       MeshGroupClient
	upstream        gloo_solo_io.UpstreamClient
	routingRule     RoutingRuleClient
	encryptionRule  EncryptionRuleClient
	tlsSecret       TlsSecretClient
	destinationRule istio_networking_v1alpha3.DestinationRuleClient
	virtualService  istio_networking_v1alpha3.VirtualServiceClient
}

func (c *configEmitter) Register() error {
	if err := c.mesh.Register(); err != nil {
		return err
	}
	if err := c.meshGroup.Register(); err != nil {
		return err
	}
	if err := c.upstream.Register(); err != nil {
		return err
	}
	if err := c.routingRule.Register(); err != nil {
		return err
	}
	if err := c.encryptionRule.Register(); err != nil {
		return err
	}
	if err := c.tlsSecret.Register(); err != nil {
		return err
	}
	if err := c.destinationRule.Register(); err != nil {
		return err
	}
	if err := c.virtualService.Register(); err != nil {
		return err
	}
	return nil
}

func (c *configEmitter) Mesh() MeshClient {
	return c.mesh
}

func (c *configEmitter) MeshGroup() MeshGroupClient {
	return c.meshGroup
}

func (c *configEmitter) Upstream() gloo_solo_io.UpstreamClient {
	return c.upstream
}

func (c *configEmitter) RoutingRule() RoutingRuleClient {
	return c.routingRule
}

func (c *configEmitter) EncryptionRule() EncryptionRuleClient {
	return c.encryptionRule
}

func (c *configEmitter) TlsSecret() TlsSecretClient {
	return c.tlsSecret
}

func (c *configEmitter) DestinationRule() istio_networking_v1alpha3.DestinationRuleClient {
	return c.destinationRule
}

func (c *configEmitter) VirtualService() istio_networking_v1alpha3.VirtualServiceClient {
	return c.virtualService
}

func (c *configEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *ConfigSnapshot, <-chan error, error) {
	errs := make(chan error)
	var done sync.WaitGroup
	ctx := opts.Ctx
	/* Create channel for Mesh */
	type meshListWithNamespace struct {
		list      MeshList
		namespace string
	}
	meshChan := make(chan meshListWithNamespace)
	/* Create channel for MeshGroup */
	type meshGroupListWithNamespace struct {
		list      MeshGroupList
		namespace string
	}
	meshGroupChan := make(chan meshGroupListWithNamespace)
	/* Create channel for Upstream */
	type upstreamListWithNamespace struct {
		list      gloo_solo_io.UpstreamList
		namespace string
	}
	upstreamChan := make(chan upstreamListWithNamespace)
	/* Create channel for RoutingRule */
	type routingRuleListWithNamespace struct {
		list      RoutingRuleList
		namespace string
	}
	routingRuleChan := make(chan routingRuleListWithNamespace)
	/* Create channel for EncryptionRule */
	type encryptionRuleListWithNamespace struct {
		list      EncryptionRuleList
		namespace string
	}
	encryptionRuleChan := make(chan encryptionRuleListWithNamespace)
	/* Create channel for TlsSecret */
	type tlsSecretListWithNamespace struct {
		list      TlsSecretList
		namespace string
	}
	tlsSecretChan := make(chan tlsSecretListWithNamespace)
	/* Create channel for DestinationRule */
	type destinationRuleListWithNamespace struct {
		list      istio_networking_v1alpha3.DestinationRuleList
		namespace string
	}
	destinationRuleChan := make(chan destinationRuleListWithNamespace)
	/* Create channel for VirtualService */
	type virtualServiceListWithNamespace struct {
		list      istio_networking_v1alpha3.VirtualServiceList
		namespace string
	}
	virtualServiceChan := make(chan virtualServiceListWithNamespace)

	for _, namespace := range watchNamespaces {
		/* Setup namespaced watch for Mesh */
		meshNamespacesChan, meshErrs, err := c.mesh.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Mesh watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, meshErrs, namespace+"-meshes")
		}(namespace)
		/* Setup namespaced watch for MeshGroup */
		meshGroupNamespacesChan, meshGroupErrs, err := c.meshGroup.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting MeshGroup watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, meshGroupErrs, namespace+"-meshgroups")
		}(namespace)
		/* Setup namespaced watch for Upstream */
		upstreamNamespacesChan, upstreamErrs, err := c.upstream.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Upstream watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, upstreamErrs, namespace+"-upstreams")
		}(namespace)
		/* Setup namespaced watch for RoutingRule */
		routingRuleNamespacesChan, routingRuleErrs, err := c.routingRule.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting RoutingRule watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, routingRuleErrs, namespace+"-routingrules")
		}(namespace)
		/* Setup namespaced watch for EncryptionRule */
		encryptionRuleNamespacesChan, encryptionRuleErrs, err := c.encryptionRule.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting EncryptionRule watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, encryptionRuleErrs, namespace+"-ecryptionrules")
		}(namespace)
		/* Setup namespaced watch for TlsSecret */
		tlsSecretNamespacesChan, tlsSecretErrs, err := c.tlsSecret.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting TlsSecret watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, tlsSecretErrs, namespace+"-tlssecrets")
		}(namespace)
		/* Setup namespaced watch for DestinationRule */
		destinationRuleNamespacesChan, destinationRuleErrs, err := c.destinationRule.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting DestinationRule watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, destinationRuleErrs, namespace+"-destinationrules")
		}(namespace)
		/* Setup namespaced watch for VirtualService */
		virtualServiceNamespacesChan, virtualServiceErrs, err := c.virtualService.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting VirtualService watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, virtualServiceErrs, namespace+"-virtualservices")
		}(namespace)

		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-ctx.Done():
					return
				case meshList := <-meshNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case meshChan <- meshListWithNamespace{list: meshList, namespace: namespace}:
					}
				case meshGroupList := <-meshGroupNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case meshGroupChan <- meshGroupListWithNamespace{list: meshGroupList, namespace: namespace}:
					}
				case upstreamList := <-upstreamNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case upstreamChan <- upstreamListWithNamespace{list: upstreamList, namespace: namespace}:
					}
				case routingRuleList := <-routingRuleNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case routingRuleChan <- routingRuleListWithNamespace{list: routingRuleList, namespace: namespace}:
					}
				case encryptionRuleList := <-encryptionRuleNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case encryptionRuleChan <- encryptionRuleListWithNamespace{list: encryptionRuleList, namespace: namespace}:
					}
				case tlsSecretList := <-tlsSecretNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case tlsSecretChan <- tlsSecretListWithNamespace{list: tlsSecretList, namespace: namespace}:
					}
				case destinationRuleList := <-destinationRuleNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case destinationRuleChan <- destinationRuleListWithNamespace{list: destinationRuleList, namespace: namespace}:
					}
				case virtualServiceList := <-virtualServiceNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case virtualServiceChan <- virtualServiceListWithNamespace{list: virtualServiceList, namespace: namespace}:
					}
				}
			}
		}(namespace)
	}

	snapshots := make(chan *ConfigSnapshot)
	go func() {
		originalSnapshot := ConfigSnapshot{}
		currentSnapshot := originalSnapshot.Clone()
		timer := time.NewTicker(time.Second * 1)
		sync := func() {
			if originalSnapshot.Hash() == currentSnapshot.Hash() {
				return
			}

			stats.Record(ctx, mConfigSnapshotOut.M(1))
			originalSnapshot = currentSnapshot.Clone()
			sentSnapshot := currentSnapshot.Clone()
			snapshots <- &sentSnapshot
		}

		/* TODO (yuval-k): figure out how to make this work to avoid a stale snapshot.
		   		// construct the first snapshot from all the configs that are currently there
		   		// that guarantees that the first snapshot contains all the data.
		   		for range watchNamespaces {
		      meshNamespacedList := <- meshChan
		      currentSnapshot.Meshes.Clear(meshNamespacedList.namespace)
		      meshList := meshNamespacedList.list
		   	currentSnapshot.Meshes.Add(meshList...)
		      meshGroupNamespacedList := <- meshGroupChan
		      currentSnapshot.Meshgroups.Clear(meshGroupNamespacedList.namespace)
		      meshGroupList := meshGroupNamespacedList.list
		   	currentSnapshot.Meshgroups.Add(meshGroupList...)
		      upstreamNamespacedList := <- upstreamChan
		      currentSnapshot.Upstreams.Clear(upstreamNamespacedList.namespace)
		      upstreamList := upstreamNamespacedList.list
		   	currentSnapshot.Upstreams.Add(upstreamList...)
		      routingRuleNamespacedList := <- routingRuleChan
		      currentSnapshot.Routingrules.Clear(routingRuleNamespacedList.namespace)
		      routingRuleList := routingRuleNamespacedList.list
		   	currentSnapshot.Routingrules.Add(routingRuleList...)
		      encryptionRuleNamespacedList := <- encryptionRuleChan
		      currentSnapshot.Ecryptionrules.Clear(encryptionRuleNamespacedList.namespace)
		      encryptionRuleList := encryptionRuleNamespacedList.list
		   	currentSnapshot.Ecryptionrules.Add(encryptionRuleList...)
		      tlsSecretNamespacedList := <- tlsSecretChan
		      currentSnapshot.Tlssecrets.Clear(tlsSecretNamespacedList.namespace)
		      tlsSecretList := tlsSecretNamespacedList.list
		   	currentSnapshot.Tlssecrets.Add(tlsSecretList...)
		      destinationRuleNamespacedList := <- destinationRuleChan
		      currentSnapshot.Destinationrules.Clear(destinationRuleNamespacedList.namespace)
		      destinationRuleList := destinationRuleNamespacedList.list
		   	currentSnapshot.Destinationrules.Add(destinationRuleList...)
		      virtualServiceNamespacedList := <- virtualServiceChan
		      currentSnapshot.Virtualservices.Clear(virtualServiceNamespacedList.namespace)
		      virtualServiceList := virtualServiceNamespacedList.list
		   	currentSnapshot.Virtualservices.Add(virtualServiceList...)
		   		}
		*/

		for {
			record := func() { stats.Record(ctx, mConfigSnapshotIn.M(1)) }

			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				close(snapshots)
				done.Wait()
				close(errs)
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case meshNamespacedList := <-meshChan:
				record()

				namespace := meshNamespacedList.namespace
				meshList := meshNamespacedList.list

				currentSnapshot.Meshes.Clear(namespace)
				currentSnapshot.Meshes.Add(meshList...)
			case meshGroupNamespacedList := <-meshGroupChan:
				record()

				namespace := meshGroupNamespacedList.namespace
				meshGroupList := meshGroupNamespacedList.list

				currentSnapshot.Meshgroups.Clear(namespace)
				currentSnapshot.Meshgroups.Add(meshGroupList...)
			case upstreamNamespacedList := <-upstreamChan:
				record()

				namespace := upstreamNamespacedList.namespace
				upstreamList := upstreamNamespacedList.list

				currentSnapshot.Upstreams.Clear(namespace)
				currentSnapshot.Upstreams.Add(upstreamList...)
			case routingRuleNamespacedList := <-routingRuleChan:
				record()

				namespace := routingRuleNamespacedList.namespace
				routingRuleList := routingRuleNamespacedList.list

				currentSnapshot.Routingrules.Clear(namespace)
				currentSnapshot.Routingrules.Add(routingRuleList...)
			case encryptionRuleNamespacedList := <-encryptionRuleChan:
				record()

				namespace := encryptionRuleNamespacedList.namespace
				encryptionRuleList := encryptionRuleNamespacedList.list

				currentSnapshot.Ecryptionrules.Clear(namespace)
				currentSnapshot.Ecryptionrules.Add(encryptionRuleList...)
			case tlsSecretNamespacedList := <-tlsSecretChan:
				record()

				namespace := tlsSecretNamespacedList.namespace
				tlsSecretList := tlsSecretNamespacedList.list

				currentSnapshot.Tlssecrets.Clear(namespace)
				currentSnapshot.Tlssecrets.Add(tlsSecretList...)
			case destinationRuleNamespacedList := <-destinationRuleChan:
				record()

				namespace := destinationRuleNamespacedList.namespace
				destinationRuleList := destinationRuleNamespacedList.list

				currentSnapshot.Destinationrules.Clear(namespace)
				currentSnapshot.Destinationrules.Add(destinationRuleList...)
			case virtualServiceNamespacedList := <-virtualServiceChan:
				record()

				namespace := virtualServiceNamespacedList.namespace
				virtualServiceList := virtualServiceNamespacedList.list

				currentSnapshot.Virtualservices.Clear(namespace)
				currentSnapshot.Virtualservices.Add(virtualServiceList...)
			}
		}
	}()
	return snapshots, errs, nil
}

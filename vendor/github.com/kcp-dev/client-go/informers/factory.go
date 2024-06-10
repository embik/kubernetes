//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by kcp code-generator. DO NOT EDIT.

package informers

import (
	"reflect"
	"sync"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	upstreaminformers "k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"

	admissionregistrationinformers "github.com/kcp-dev/client-go/informers/admissionregistration"
	apiserverinternalinformers "github.com/kcp-dev/client-go/informers/apiserverinternal"
	appsinformers "github.com/kcp-dev/client-go/informers/apps"
	autoscalinginformers "github.com/kcp-dev/client-go/informers/autoscaling"
	batchinformers "github.com/kcp-dev/client-go/informers/batch"
	certificatesinformers "github.com/kcp-dev/client-go/informers/certificates"
	coordinationinformers "github.com/kcp-dev/client-go/informers/coordination"
	coreinformers "github.com/kcp-dev/client-go/informers/core"
	discoveryinformers "github.com/kcp-dev/client-go/informers/discovery"
	eventsinformers "github.com/kcp-dev/client-go/informers/events"
	extensionsinformers "github.com/kcp-dev/client-go/informers/extensions"
	flowcontrolinformers "github.com/kcp-dev/client-go/informers/flowcontrol"
	"github.com/kcp-dev/client-go/informers/internalinterfaces"
	networkinginformers "github.com/kcp-dev/client-go/informers/networking"
	nodeinformers "github.com/kcp-dev/client-go/informers/node"
	policyinformers "github.com/kcp-dev/client-go/informers/policy"
	rbacinformers "github.com/kcp-dev/client-go/informers/rbac"
	resourceinformers "github.com/kcp-dev/client-go/informers/resource"
	schedulinginformers "github.com/kcp-dev/client-go/informers/scheduling"
	storageinformers "github.com/kcp-dev/client-go/informers/storage"
	storagemigrationinformers "github.com/kcp-dev/client-go/informers/storagemigration"
	clientset "github.com/kcp-dev/client-go/kubernetes"
)

// SharedInformerOption defines the functional option type for SharedInformerFactory.
type SharedInformerOption func(*SharedInformerOptions) *SharedInformerOptions

type SharedInformerOptions struct {
	customResync     map[reflect.Type]time.Duration
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	transform        cache.TransformFunc
}

type sharedInformerFactory struct {
	client           clientset.ClusterInterface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	lock             sync.Mutex
	defaultResync    time.Duration
	customResync     map[reflect.Type]time.Duration
	transform        cache.TransformFunc

	informers map[reflect.Type]kcpcache.ScopeableSharedIndexInformer
	// startedInformers is used for tracking which informers have been started.
	// This allows Start() to be called multiple times safely.
	startedInformers map[reflect.Type]bool
	// wg tracks how many goroutines were started.
	wg sync.WaitGroup
	// shuttingDown is true when Shutdown has been called. It may still be running
	// because it needs to wait for goroutines.
	shuttingDown bool
}

// WithCustomResyncConfig sets a custom resync period for the specified informer types.
func WithCustomResyncConfig(resyncConfig map[metav1.Object]time.Duration) SharedInformerOption {
	return func(opts *SharedInformerOptions) *SharedInformerOptions {
		for k, v := range resyncConfig {
			opts.customResync[reflect.TypeOf(k)] = v
		}
		return opts
	}
}

// WithTweakListOptions sets a custom filter on all listers of the configured SharedInformerFactory.
func WithTweakListOptions(tweakListOptions internalinterfaces.TweakListOptionsFunc) SharedInformerOption {
	return func(opts *SharedInformerOptions) *SharedInformerOptions {
		opts.tweakListOptions = tweakListOptions
		return opts
	}
}

// WithTransform sets a transform on all informers.
func WithTransform(transform cache.TransformFunc) SharedInformerOption {
	return func(opts *SharedInformerOptions) *SharedInformerOptions {
		opts.transform = transform
		return opts
	}
}

// NewSharedInformerFactory constructs a new instance of SharedInformerFactory for all namespaces.
func NewSharedInformerFactory(client clientset.ClusterInterface, defaultResync time.Duration) SharedInformerFactory {
	return NewSharedInformerFactoryWithOptions(client, defaultResync)
}

// NewSharedInformerFactoryWithOptions constructs a new instance of a SharedInformerFactory with additional options.
func NewSharedInformerFactoryWithOptions(client clientset.ClusterInterface, defaultResync time.Duration, options ...SharedInformerOption) SharedInformerFactory {
	factory := &sharedInformerFactory{
		client:           client,
		defaultResync:    defaultResync,
		informers:        make(map[reflect.Type]kcpcache.ScopeableSharedIndexInformer),
		startedInformers: make(map[reflect.Type]bool),
		customResync:     make(map[reflect.Type]time.Duration),
	}

	opts := &SharedInformerOptions{
		customResync: make(map[reflect.Type]time.Duration),
	}

	// Apply all options
	for _, opt := range options {
		opts = opt(opts)
	}

	// Forward options to the factory
	factory.customResync = opts.customResync
	factory.tweakListOptions = opts.tweakListOptions
	factory.transform = opts.transform

	return factory
}

// Start initializes all requested informers.
func (f *sharedInformerFactory) Start(stopCh <-chan struct{}) {
	f.lock.Lock()
	defer f.lock.Unlock()

	if f.shuttingDown {
		return
	}

	for informerType, informer := range f.informers {
		if !f.startedInformers[informerType] {
			f.wg.Add(1)
			// We need a new variable in each loop iteration,
			// otherwise the goroutine would use the loop variable
			// and that keeps changing.
			informer := informer
			go func() {
				defer f.wg.Done()
				informer.Run(stopCh)
			}()
			f.startedInformers[informerType] = true
		}
	}
}

func (f *sharedInformerFactory) Shutdown() {
	f.lock.Lock()
	f.shuttingDown = true
	f.lock.Unlock()

	// Will return immediately if there is nothing to wait for.
	f.wg.Wait()
}

// WaitForCacheSync waits for all started informers' cache were synced.
func (f *sharedInformerFactory) WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool {
	informers := func() map[reflect.Type]kcpcache.ScopeableSharedIndexInformer {
		f.lock.Lock()
		defer f.lock.Unlock()

		informers := map[reflect.Type]kcpcache.ScopeableSharedIndexInformer{}
		for informerType, informer := range f.informers {
			if f.startedInformers[informerType] {
				informers[informerType] = informer
			}
		}
		return informers
	}()

	res := map[reflect.Type]bool{}
	for informType, informer := range informers {
		res[informType] = cache.WaitForCacheSync(stopCh, informer.HasSynced)
	}
	return res
}

// InformerFor returns the SharedIndexInformer for obj.
func (f *sharedInformerFactory) InformerFor(obj runtime.Object, newFunc internalinterfaces.NewInformerFunc) kcpcache.ScopeableSharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerType := reflect.TypeOf(obj)
	informer, exists := f.informers[informerType]
	if exists {
		return informer
	}

	resyncPeriod, exists := f.customResync[informerType]
	if !exists {
		resyncPeriod = f.defaultResync
	}

	informer = newFunc(f.client, resyncPeriod)
	f.informers[informerType] = informer

	return informer
}

type ScopedDynamicSharedInformerFactory interface {
	// ForResource gives generic access to a shared informer of the matching type.
	ForResource(resource schema.GroupVersionResource) (upstreaminformers.GenericInformer, error)

	// Start initializes all requested informers. They are handled in goroutines
	// which run until the stop channel gets closed.
	Start(stopCh <-chan struct{})
}

// SharedInformerFactory provides shared informers for resources in all known
// API group versions.
//
// It is typically used like this:
//
//	ctx, cancel := context.Background()
//	defer cancel()
//	factory := NewSharedInformerFactoryWithOptions(client, resyncPeriod)
//	defer factory.Shutdown()    // Returns immediately if nothing was started.
//	genericInformer := factory.ForResource(resource)
//	typedInformer := factory.SomeAPIGroup().V1().SomeType()
//	factory.Start(ctx.Done())          // Start processing these informers.
//	synced := factory.WaitForCacheSync(ctx.Done())
//	for v, ok := range synced {
//	    if !ok {
//	        fmt.Fprintf(os.Stderr, "caches failed to sync: %v", v)
//	        return
//	    }
//	}
//
//	// Creating informers can also be created after Start, but then
//	// Start must be called again:
//	anotherGenericInformer := factory.ForResource(resource)
//	factory.Start(ctx.Done())
type SharedInformerFactory interface {
	internalinterfaces.SharedInformerFactory

	Cluster(logicalcluster.Name) ScopedDynamicSharedInformerFactory

	// Start initializes all requested informers. They are handled in goroutines
	// which run until the stop channel gets closed.
	Start(stopCh <-chan struct{})

	// Shutdown marks a factory as shutting down. At that point no new
	// informers can be started anymore and Start will return without
	// doing anything.
	//
	// In addition, Shutdown blocks until all goroutines have terminated. For that
	// to happen, the close channel(s) that they were started with must be closed,
	// either before Shutdown gets called or while it is waiting.
	//
	// Shutdown may be called multiple times, even concurrently. All such calls will
	// block until all goroutines have terminated.
	Shutdown()

	// ForResource gives generic access to a shared informer of the matching type.
	ForResource(resource schema.GroupVersionResource) (GenericClusterInformer, error)

	// WaitForCacheSync blocks until all started informers' caches were synced
	// or the stop channel gets closed.
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool

	// InformerFor returns the SharedIndexInformer for obj.
	InformerFor(obj runtime.Object, newFunc internalinterfaces.NewInformerFunc) kcpcache.ScopeableSharedIndexInformer

	Admissionregistration() admissionregistrationinformers.ClusterInterface
	Internal() apiserverinternalinformers.ClusterInterface
	Apps() appsinformers.ClusterInterface
	Autoscaling() autoscalinginformers.ClusterInterface
	Batch() batchinformers.ClusterInterface
	Certificates() certificatesinformers.ClusterInterface
	Coordination() coordinationinformers.ClusterInterface
	Core() coreinformers.ClusterInterface
	Discovery() discoveryinformers.ClusterInterface
	Events() eventsinformers.ClusterInterface
	Extensions() extensionsinformers.ClusterInterface
	Flowcontrol() flowcontrolinformers.ClusterInterface
	Networking() networkinginformers.ClusterInterface
	Node() nodeinformers.ClusterInterface
	Policy() policyinformers.ClusterInterface
	Rbac() rbacinformers.ClusterInterface
	Resource() resourceinformers.ClusterInterface
	Scheduling() schedulinginformers.ClusterInterface
	Storage() storageinformers.ClusterInterface
	Storagemigration() storagemigrationinformers.ClusterInterface
}

func (f *sharedInformerFactory) Admissionregistration() admissionregistrationinformers.ClusterInterface {
	return admissionregistrationinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Internal() apiserverinternalinformers.ClusterInterface {
	return apiserverinternalinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Apps() appsinformers.ClusterInterface {
	return appsinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Autoscaling() autoscalinginformers.ClusterInterface {
	return autoscalinginformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Batch() batchinformers.ClusterInterface {
	return batchinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Certificates() certificatesinformers.ClusterInterface {
	return certificatesinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Coordination() coordinationinformers.ClusterInterface {
	return coordinationinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Core() coreinformers.ClusterInterface {
	return coreinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Discovery() discoveryinformers.ClusterInterface {
	return discoveryinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Events() eventsinformers.ClusterInterface {
	return eventsinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Extensions() extensionsinformers.ClusterInterface {
	return extensionsinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Flowcontrol() flowcontrolinformers.ClusterInterface {
	return flowcontrolinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Networking() networkinginformers.ClusterInterface {
	return networkinginformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Node() nodeinformers.ClusterInterface {
	return nodeinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Policy() policyinformers.ClusterInterface {
	return policyinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Rbac() rbacinformers.ClusterInterface {
	return rbacinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Resource() resourceinformers.ClusterInterface {
	return resourceinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Scheduling() schedulinginformers.ClusterInterface {
	return schedulinginformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Storage() storageinformers.ClusterInterface {
	return storageinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Storagemigration() storagemigrationinformers.ClusterInterface {
	return storagemigrationinformers.New(f, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cluster(clusterName logicalcluster.Name) ScopedDynamicSharedInformerFactory {
	return &scopedDynamicSharedInformerFactory{
		sharedInformerFactory: f,
		clusterName:           clusterName,
	}
}

type scopedDynamicSharedInformerFactory struct {
	*sharedInformerFactory
	clusterName logicalcluster.Name
}

func (f *scopedDynamicSharedInformerFactory) ForResource(resource schema.GroupVersionResource) (upstreaminformers.GenericInformer, error) {
	clusterInformer, err := f.sharedInformerFactory.ForResource(resource)
	if err != nil {
		return nil, err
	}
	return clusterInformer.Cluster(f.clusterName), nil
}

func (f *scopedDynamicSharedInformerFactory) Start(stopCh <-chan struct{}) {
	f.sharedInformerFactory.Start(stopCh)
}

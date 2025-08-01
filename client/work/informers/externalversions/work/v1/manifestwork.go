// Copyright Contributors to the Open Cluster Management project
// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	context "context"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	versioned "open-cluster-management.io/api/client/work/clientset/versioned"
	internalinterfaces "open-cluster-management.io/api/client/work/informers/externalversions/internalinterfaces"
	workv1 "open-cluster-management.io/api/client/work/listers/work/v1"
	apiworkv1 "open-cluster-management.io/api/work/v1"
)

// ManifestWorkInformer provides access to a shared informer and lister for
// ManifestWorks.
type ManifestWorkInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() workv1.ManifestWorkLister
}

type manifestWorkInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewManifestWorkInformer constructs a new informer for ManifestWork type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewManifestWorkInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredManifestWorkInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredManifestWorkInformer constructs a new informer for ManifestWork type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredManifestWorkInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WorkV1().ManifestWorks(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WorkV1().ManifestWorks(namespace).Watch(context.TODO(), options)
			},
		},
		&apiworkv1.ManifestWork{},
		resyncPeriod,
		indexers,
	)
}

func (f *manifestWorkInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredManifestWorkInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *manifestWorkInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiworkv1.ManifestWork{}, f.defaultInformer)
}

func (f *manifestWorkInformer) Lister() workv1.ManifestWorkLister {
	return workv1.NewManifestWorkLister(f.Informer().GetIndexer())
}

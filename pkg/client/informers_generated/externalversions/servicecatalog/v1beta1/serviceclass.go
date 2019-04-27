package v1beta1

import (
	time "time"
	servicecatalogv1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	clientset "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/kubernetes-incubator/service-catalog/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/client/listers_generated/servicecatalog/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

type ServiceClassInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ServiceClassLister
}
type serviceClassInformer struct {
	factory			internalinterfaces.SharedInformerFactory
	tweakListOptions	internalinterfaces.TweakListOptionsFunc
	namespace		string
}

func NewServiceClassInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewFilteredServiceClassInformer(client, namespace, resyncPeriod, indexers, nil)
}
func NewFilteredServiceClassInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cache.NewSharedIndexInformer(&cache.ListWatch{ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
		if tweakListOptions != nil {
			tweakListOptions(&options)
		}
		return client.ServicecatalogV1beta1().ServiceClasses(namespace).List(options)
	}, WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
		if tweakListOptions != nil {
			tweakListOptions(&options)
		}
		return client.ServicecatalogV1beta1().ServiceClasses(namespace).Watch(options)
	}}, &servicecatalogv1beta1.ServiceClass{}, resyncPeriod, indexers)
}
func (f *serviceClassInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewFilteredServiceClassInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}
func (f *serviceClassInformer) Informer() cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.factory.InformerFor(&servicecatalogv1beta1.ServiceClass{}, f.defaultInformer)
}
func (f *serviceClassInformer) Lister() v1beta1.ServiceClassLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return v1beta1.NewServiceClassLister(f.Informer().GetIndexer())
}

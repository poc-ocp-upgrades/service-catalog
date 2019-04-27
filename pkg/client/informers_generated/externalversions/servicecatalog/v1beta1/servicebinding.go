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

type ServiceBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ServiceBindingLister
}
type serviceBindingInformer struct {
	factory			internalinterfaces.SharedInformerFactory
	tweakListOptions	internalinterfaces.TweakListOptionsFunc
	namespace		string
}

func NewServiceBindingInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewFilteredServiceBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}
func NewFilteredServiceBindingInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cache.NewSharedIndexInformer(&cache.ListWatch{ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
		if tweakListOptions != nil {
			tweakListOptions(&options)
		}
		return client.ServicecatalogV1beta1().ServiceBindings(namespace).List(options)
	}, WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
		if tweakListOptions != nil {
			tweakListOptions(&options)
		}
		return client.ServicecatalogV1beta1().ServiceBindings(namespace).Watch(options)
	}}, &servicecatalogv1beta1.ServiceBinding{}, resyncPeriod, indexers)
}
func (f *serviceBindingInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewFilteredServiceBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}
func (f *serviceBindingInformer) Informer() cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.factory.InformerFor(&servicecatalogv1beta1.ServiceBinding{}, f.defaultInformer)
}
func (f *serviceBindingInformer) Lister() v1beta1.ServiceBindingLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return v1beta1.NewServiceBindingLister(f.Informer().GetIndexer())
}

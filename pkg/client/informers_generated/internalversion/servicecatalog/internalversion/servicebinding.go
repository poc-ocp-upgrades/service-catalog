package internalversion

import (
	time "time"
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	internalclientset "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "github.com/kubernetes-incubator/service-catalog/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "github.com/kubernetes-incubator/service-catalog/pkg/client/listers_generated/servicecatalog/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

type ServiceBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.ServiceBindingLister
}
type serviceBindingInformer struct {
	factory			internalinterfaces.SharedInformerFactory
	tweakListOptions	internalinterfaces.TweakListOptionsFunc
	namespace		string
}

func NewServiceBindingInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewFilteredServiceBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}
func NewFilteredServiceBindingInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cache.NewSharedIndexInformer(&cache.ListWatch{ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
		if tweakListOptions != nil {
			tweakListOptions(&options)
		}
		return client.Servicecatalog().ServiceBindings(namespace).List(options)
	}, WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
		if tweakListOptions != nil {
			tweakListOptions(&options)
		}
		return client.Servicecatalog().ServiceBindings(namespace).Watch(options)
	}}, &servicecatalog.ServiceBinding{}, resyncPeriod, indexers)
}
func (f *serviceBindingInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewFilteredServiceBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}
func (f *serviceBindingInformer) Informer() cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.factory.InformerFor(&servicecatalog.ServiceBinding{}, f.defaultInformer)
}
func (f *serviceBindingInformer) Lister() internalversion.ServiceBindingLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return internalversion.NewServiceBindingLister(f.Informer().GetIndexer())
}

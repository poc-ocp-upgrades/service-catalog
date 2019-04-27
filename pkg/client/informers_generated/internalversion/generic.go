package internalversion

import (
	"fmt"
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	settings "github.com/kubernetes-incubator/service-catalog/pkg/apis/settings"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}
type genericInformer struct {
	informer	cache.SharedIndexInformer
	resource	schema.GroupResource
}

func (f *genericInformer) Informer() cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.informer
}
func (f *genericInformer) Lister() cache.GenericLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch resource {
	case servicecatalog.SchemeGroupVersion.WithResource("clusterservicebrokers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicecatalog().InternalVersion().ClusterServiceBrokers().Informer()}, nil
	case servicecatalog.SchemeGroupVersion.WithResource("clusterserviceclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicecatalog().InternalVersion().ClusterServiceClasses().Informer()}, nil
	case servicecatalog.SchemeGroupVersion.WithResource("clusterserviceplans"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicecatalog().InternalVersion().ClusterServicePlans().Informer()}, nil
	case servicecatalog.SchemeGroupVersion.WithResource("servicebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicecatalog().InternalVersion().ServiceBindings().Informer()}, nil
	case servicecatalog.SchemeGroupVersion.WithResource("servicebrokers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicecatalog().InternalVersion().ServiceBrokers().Informer()}, nil
	case servicecatalog.SchemeGroupVersion.WithResource("serviceclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicecatalog().InternalVersion().ServiceClasses().Informer()}, nil
	case servicecatalog.SchemeGroupVersion.WithResource("serviceinstances"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicecatalog().InternalVersion().ServiceInstances().Informer()}, nil
	case servicecatalog.SchemeGroupVersion.WithResource("serviceplans"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicecatalog().InternalVersion().ServicePlans().Informer()}, nil
	case settings.SchemeGroupVersion.WithResource("podpresets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Settings().InternalVersion().PodPresets().Informer()}, nil
	}
	return nil, fmt.Errorf("no informer found for %v", resource)
}

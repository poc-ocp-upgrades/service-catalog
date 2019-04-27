package externalversions

import (
	"fmt"
	v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1alpha1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/settings/v1alpha1"
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
	case v1beta1.SchemeGroupVersion.WithResource("clusterservicebrokers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicecatalog().V1beta1().ClusterServiceBrokers().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("clusterserviceclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicecatalog().V1beta1().ClusterServiceClasses().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("clusterserviceplans"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicecatalog().V1beta1().ClusterServicePlans().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("servicebindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicecatalog().V1beta1().ServiceBindings().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("servicebrokers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicecatalog().V1beta1().ServiceBrokers().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("serviceclasses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicecatalog().V1beta1().ServiceClasses().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("serviceinstances"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicecatalog().V1beta1().ServiceInstances().Informer()}, nil
	case v1beta1.SchemeGroupVersion.WithResource("serviceplans"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Servicecatalog().V1beta1().ServicePlans().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("podpresets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Settings().V1alpha1().PodPresets().Informer()}, nil
	}
	return nil, fmt.Errorf("no informer found for %v", resource)
}

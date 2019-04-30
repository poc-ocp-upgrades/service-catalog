package servicecatalog

import (
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type FilterOptions struct{ ClassID string }
type RegisterOptions struct {
	BasicSecret		string
	BearerSecret		string
	CAFile			string
	ClassRestrictions	[]string
	Namespace		string
	PlanRestrictions	[]string
	RelistBehavior		v1beta1.ServiceBrokerRelistBehavior
	RelistDuration		*metav1.Duration
	SkipTLS			bool
}
type ProvisionOptions struct {
	ExternalID	string
	Namespace	string
	Params		interface{}
	Secrets		map[string]string
}

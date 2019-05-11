package servicecatalog

import (
	"errors"
	"fmt"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
)

const (
	FieldExternalClassName = "spec.externalName"
)

type CreateClassFromOptions struct {
	Name		string
	Scope		Scope
	Namespace	string
	From		string
}
type Class interface {
	GetName() string
	GetNamespace() string
	GetExternalName() string
	GetDescription() string
	GetSpec() v1beta1.CommonServiceClassSpec
	GetServiceBrokerName() string
	GetStatusText() string
}

func (sdk *SDK) RetrieveClasses(opts ScopeOptions) ([]Class, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var classes []Class
	if opts.Scope.Matches(ClusterScope) {
		csc, err := sdk.ServiceCatalog().ClusterServiceClasses().List(metav1.ListOptions{})
		if err != nil {
			return nil, fmt.Errorf("unable to list cluster-scoped classes (%s)", err)
		}
		for _, c := range csc.Items {
			class := c
			classes = append(classes, &class)
		}
	}
	if opts.Scope.Matches(NamespaceScope) {
		sc, err := sdk.ServiceCatalog().ServiceClasses(opts.Namespace).List(metav1.ListOptions{})
		if err != nil {
			if apierrors.IsNotFound(err) {
				return classes, nil
			}
			return nil, fmt.Errorf("unable to list classes in %q (%s)", opts.Namespace, err)
		}
		for _, c := range sc.Items {
			class := c
			classes = append(classes, &class)
		}
	}
	return classes, nil
}
func (sdk *SDK) RetrieveClassByName(name string, opts ScopeOptions) (Class, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var searchResults []Class
	lopts := metav1.ListOptions{FieldSelector: fields.OneTermEqualSelector(FieldExternalClassName, name).String()}
	if opts.Scope.Matches(ClusterScope) {
		csc, err := sdk.ServiceCatalog().ClusterServiceClasses().List(lopts)
		if err != nil {
			return nil, fmt.Errorf("unable to search classes by name (%s)", err)
		}
		for _, c := range csc.Items {
			class := c
			searchResults = append(searchResults, &class)
		}
	}
	if opts.Scope.Matches(NamespaceScope) {
		sc, err := sdk.ServiceCatalog().ServiceClasses(opts.Namespace).List(lopts)
		if err != nil {
			if apierrors.IsNotFound(err) {
				sc = &v1beta1.ServiceClassList{}
			} else {
				return nil, fmt.Errorf("unable to search classes by name (%s)", err)
			}
		}
		for _, c := range sc.Items {
			class := c
			searchResults = append(searchResults, &class)
		}
	}
	if len(searchResults) > 1 {
		return nil, fmt.Errorf("more than one matching class found for '%s' %d", name, len(searchResults))
	}
	if len(searchResults) == 0 {
		if opts.Scope.Matches(ClusterScope) {
			return nil, fmt.Errorf("class '%s' not found in cluster scope", name)
		} else if opts.Scope.Matches(NamespaceScope) {
			if opts.Namespace == "" {
				return nil, fmt.Errorf("class '%s' not found in any namespace", name)
			}
			return nil, fmt.Errorf("class '%s' not found in namespace %s", name, opts.Namespace)
		}
		return nil, fmt.Errorf("class '%s' not found", name)
	}
	return searchResults[0], nil
}
func (sdk *SDK) RetrieveClassByID(kubeName string) (*v1beta1.ClusterServiceClass, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	class, err := sdk.ServiceCatalog().ClusterServiceClasses().Get(kubeName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("unable to get class (%s)", err)
	}
	return class, nil
}
func (sdk *SDK) RetrieveClassByPlan(plan Plan) (*v1beta1.ClusterServiceClass, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	class, err := sdk.ServiceCatalog().ClusterServiceClasses().Get(plan.GetClassID(), metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("unable to get class (%s)", err)
	}
	return class, nil
}
func (sdk *SDK) CreateClassFrom(opts CreateClassFromOptions) (Class, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if opts.Scope == AllScope {
		return nil, errors.New("invalid scope: all")
	}
	fromClass, err := sdk.RetrieveClassByName(opts.From, ScopeOptions{Scope: opts.Scope, Namespace: opts.Namespace})
	if err != nil {
		return nil, err
	}
	if opts.Scope.Matches(ClusterScope) {
		csc := fromClass.(*v1beta1.ClusterServiceClass)
		return sdk.createClusterServiceClass(csc, opts.Name)
	}
	sc := fromClass.(*v1beta1.ServiceClass)
	return sdk.createServiceClass(sc, opts.Name, opts.Namespace)
}
func (sdk *SDK) createClusterServiceClass(from *v1beta1.ClusterServiceClass, name string) (*v1beta1.ClusterServiceClass, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var class = &v1beta1.ClusterServiceClass{ObjectMeta: metav1.ObjectMeta{Name: name}, Spec: from.Spec}
	class.Spec.ExternalName = name
	created, err := sdk.ServiceCatalog().ClusterServiceClasses().Create(class)
	if err != nil {
		return nil, fmt.Errorf("unable to create cluster service class (%s)", err)
	}
	return created, nil
}
func (sdk *SDK) createServiceClass(from *v1beta1.ServiceClass, name, namespace string) (*v1beta1.ServiceClass, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var class = &v1beta1.ServiceClass{ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: namespace}, Spec: from.Spec}
	class.Spec.ExternalName = name
	created, err := sdk.ServiceCatalog().ServiceClasses(namespace).Create(class)
	if err != nil {
		return nil, fmt.Errorf("unable to create service class (%s)", err)
	}
	return created, nil
}

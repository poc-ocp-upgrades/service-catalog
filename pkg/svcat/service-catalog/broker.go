package servicecatalog

import (
	"fmt"
	"io/ioutil"
	"math"
	"time"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
)

type Broker interface {
	GetName() string
	GetNamespace() string
	GetURL() string
	GetSpec() v1beta1.CommonServiceBrokerSpec
	GetStatus() v1beta1.CommonServiceBrokerStatus
}

func (sdk *SDK) Deregister(brokerName string, scopeOpts *ScopeOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if scopeOpts.Scope.Matches(NamespaceScope) {
		err := sdk.ServiceCatalog().ServiceBrokers(scopeOpts.Namespace).Delete(brokerName, &v1.DeleteOptions{})
		if err != nil {
			return fmt.Errorf("deregister request failed (%s)", err)
		}
		return nil
	} else if scopeOpts.Scope.Matches(ClusterScope) {
		err := sdk.ServiceCatalog().ClusterServiceBrokers().Delete(brokerName, &v1.DeleteOptions{})
		if err != nil {
			return fmt.Errorf("deregister request failed (%s)", err)
		}
		return nil
	}
	return fmt.Errorf("cannot deregister broker, unrecognized scope provided (%s)", scopeOpts.Scope)
}
func (sdk *SDK) RetrieveBrokers(opts ScopeOptions) ([]Broker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var brokers []Broker
	if opts.Scope.Matches(ClusterScope) {
		csb, err := sdk.ServiceCatalog().ClusterServiceBrokers().List(v1.ListOptions{})
		if err != nil {
			return nil, fmt.Errorf("unable to list cluster-scoped brokers (%s)", err)
		}
		for _, b := range csb.Items {
			broker := b
			brokers = append(brokers, &broker)
		}
	}
	if opts.Scope.Matches(NamespaceScope) {
		sb, err := sdk.ServiceCatalog().ServiceBrokers(opts.Namespace).List(v1.ListOptions{})
		if err != nil {
			if apierrors.IsNotFound(err) {
				return brokers, nil
			}
			return nil, fmt.Errorf("unable to list brokers in %q (%s)", opts.Namespace, err)
		}
		for _, b := range sb.Items {
			broker := b
			brokers = append(brokers, &broker)
		}
	}
	return brokers, nil
}
func (sdk *SDK) RetrieveBroker(name string) (*v1beta1.ClusterServiceBroker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	broker, err := sdk.ServiceCatalog().ClusterServiceBrokers().Get(name, v1.GetOptions{})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get broker '%s'", name)
	}
	return broker, nil
}
func (sdk *SDK) RetrieveNamespacedBroker(namespace string, name string) (*v1beta1.ServiceBroker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	broker, err := sdk.ServiceCatalog().ServiceBrokers(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("unable to get broker '%s' (%s)", name, err)
	}
	return broker, nil
}
func (sdk *SDK) RetrieveBrokerByClass(class *v1beta1.ClusterServiceClass) (*v1beta1.ClusterServiceBroker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	brokerName := class.Spec.ClusterServiceBrokerName
	broker, err := sdk.ServiceCatalog().ClusterServiceBrokers().Get(brokerName, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return broker, nil
}
func (sdk *SDK) Register(brokerName string, url string, opts *RegisterOptions, scopeOpts *ScopeOptions) (Broker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	var caBytes []byte
	if opts.CAFile != "" {
		caBytes, err = ioutil.ReadFile(opts.CAFile)
		if err != nil {
			return nil, fmt.Errorf("Error opening CA file: %v", err.Error())
		}
	}
	objectMeta := v1.ObjectMeta{Name: brokerName}
	commonServiceBrokerSpec := v1beta1.CommonServiceBrokerSpec{CABundle: caBytes, InsecureSkipTLSVerify: opts.SkipTLS, RelistBehavior: opts.RelistBehavior, RelistDuration: opts.RelistDuration, URL: url, CatalogRestrictions: &v1beta1.CatalogRestrictions{ServiceClass: opts.ClassRestrictions, ServicePlan: opts.PlanRestrictions}}
	if scopeOpts.Scope.Matches(ClusterScope) {
		request := &v1beta1.ClusterServiceBroker{ObjectMeta: objectMeta, Spec: v1beta1.ClusterServiceBrokerSpec{CommonServiceBrokerSpec: commonServiceBrokerSpec}}
		if opts.BasicSecret != "" {
			request.Spec.AuthInfo = &v1beta1.ClusterServiceBrokerAuthInfo{}
			request.Spec.AuthInfo.Basic = &v1beta1.ClusterBasicAuthConfig{SecretRef: &v1beta1.ObjectReference{Name: opts.BasicSecret, Namespace: opts.Namespace}}
		} else if opts.BearerSecret != "" {
			request.Spec.AuthInfo = &v1beta1.ClusterServiceBrokerAuthInfo{}
			request.Spec.AuthInfo.Bearer = &v1beta1.ClusterBearerTokenAuthConfig{SecretRef: &v1beta1.ObjectReference{Name: opts.BearerSecret, Namespace: opts.Namespace}}
		}
		result, err := sdk.ServiceCatalog().ClusterServiceBrokers().Create(request)
		if err != nil {
			return nil, fmt.Errorf("register request failed (%s)", err)
		}
		return result, nil
	}
	request := &v1beta1.ServiceBroker{ObjectMeta: objectMeta, Spec: v1beta1.ServiceBrokerSpec{CommonServiceBrokerSpec: commonServiceBrokerSpec}}
	if opts.BasicSecret != "" {
		request.Spec.AuthInfo = &v1beta1.ServiceBrokerAuthInfo{Basic: &v1beta1.BasicAuthConfig{SecretRef: &v1beta1.LocalObjectReference{Name: opts.BasicSecret}}}
	} else if opts.BearerSecret != "" {
		request.Spec.AuthInfo = &v1beta1.ServiceBrokerAuthInfo{Bearer: &v1beta1.BearerTokenAuthConfig{SecretRef: &v1beta1.LocalObjectReference{Name: opts.BearerSecret}}}
	}
	result, err := sdk.ServiceCatalog().ServiceBrokers(scopeOpts.Namespace).Create(request)
	if err != nil {
		return nil, fmt.Errorf("register request failed (%s)", err)
	}
	return result, nil
}
func (sdk *SDK) Sync(name string, scopeOpts ScopeOptions, retries int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	success := false
	var err error
	for j := 0; j < retries && !success; j++ {
		if scopeOpts.Scope.Matches(NamespaceScope) {
			var broker *v1beta1.ServiceBroker
			namespace := scopeOpts.Namespace
			broker, err = sdk.RetrieveNamespacedBroker(namespace, name)
			if err == nil {
				broker.Spec.RelistRequests = broker.Spec.RelistRequests + 1
				_, err = sdk.ServiceCatalog().ServiceBrokers(namespace).Update(broker)
				if err == nil {
					success = true
				}
				if err != nil && !apierrors.IsConflict(err) {
					return fmt.Errorf("could not sync service broker (%s)", err)
				}
			}
		}
		if scopeOpts.Scope.Matches(ClusterScope) {
			var broker *v1beta1.ClusterServiceBroker
			broker, err = sdk.RetrieveBroker(name)
			if err == nil {
				broker.Spec.RelistRequests = broker.Spec.RelistRequests + 1
				_, err = sdk.ServiceCatalog().ClusterServiceBrokers().Update(broker)
				if err == nil {
					success = true
				}
				if err != nil && !apierrors.IsConflict(err) {
					return fmt.Errorf("could not sync service broker (%s)", err)
				}
			}
		}
		if success {
			break
		}
	}
	if !success {
		return fmt.Errorf("could not sync service broker %s (%s)", name, err)
	}
	return nil
}
func (sdk *SDK) WaitForBroker(name string, interval time.Duration, timeout *time.Duration) (broker Broker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if timeout == nil {
		notimeout := time.Duration(math.MaxInt64)
		timeout = &notimeout
	}
	err = wait.PollImmediate(interval, *timeout, func() (bool, error) {
		broker, err = sdk.RetrieveBroker(name)
		if err != nil {
			if apierrors.IsNotFound(errors.Cause(err)) {
				err = nil
			}
			return false, err
		}
		isDone := sdk.IsBrokerReady(broker) || sdk.IsBrokerFailed(broker)
		return isDone, nil
	})
	return broker, err
}
func (sdk *SDK) IsBrokerReady(broker Broker) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return sdk.BrokerHasStatus(broker, v1beta1.ServiceBrokerConditionReady)
}
func (sdk *SDK) IsBrokerFailed(broker Broker) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return sdk.BrokerHasStatus(broker, v1beta1.ServiceBrokerConditionFailed)
}
func (sdk *SDK) BrokerHasStatus(broker Broker, status v1beta1.ServiceBrokerConditionType) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, cond := range broker.GetStatus().Conditions {
		if cond.Type == status && cond.Status == v1beta1.ConditionTrue {
			return true
		}
	}
	return false
}

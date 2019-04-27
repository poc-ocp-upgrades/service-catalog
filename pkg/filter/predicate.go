package filter

import (
	"k8s.io/apimachinery/pkg/labels"
)

type Predicate interface {
	Accepts(Properties) bool
	Empty() bool
	String() string
}

func NewPredicate() Predicate {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return internalPredicate{}
}

type internalPredicate struct{ selector labels.Selector }

func (ip internalPredicate) Accepts(p Properties) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ip.Empty() {
		return true
	}
	return ip.selector.Matches(p)
}
func (ip internalPredicate) Empty() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ip.selector == nil {
		return true
	}
	return ip.selector.Empty()
}
func (ip internalPredicate) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ip.selector.String()
}

package servicecatalog

type Scope string

const (
	ClusterScope	= "cluster"
	NamespaceScope	= "namespace"
	AllScope	= "all"
)

func (s Scope) Matches(value Scope) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s == AllScope {
		return true
	}
	return s == value
}

type ScopeOptions struct {
	Namespace	string
	Scope		Scope
}

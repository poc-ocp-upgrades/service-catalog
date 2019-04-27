package command

import "github.com/spf13/pflag"

type HasNamespaceFlags interface {
	Command
	ApplyNamespaceFlags(flags *pflag.FlagSet)
}
type Namespaced struct {
	*Context
	Namespace	string
}

func NewNamespaced(cxt *Context) *Namespaced {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Namespaced{Context: cxt}
}
func (c *Namespaced) AddNamespaceFlags(flags *pflag.FlagSet, allowAll bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	flags.StringP("namespace", "n", "", "If present, the namespace scope for this request")
	if allowAll {
		flags.Bool("all-namespaces", false, "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace")
	}
}
func (c *Namespaced) ApplyNamespaceFlags(flags *pflag.FlagSet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.Namespace = c.determineNamespace(flags)
}
func (c *Namespaced) determineNamespace(flags *pflag.FlagSet) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	currentNamespace := c.Context.App.CurrentNamespace
	namespace, _ := flags.GetString("namespace")
	allNamespaces, _ := flags.GetBool("all-namespaces")
	if allNamespaces {
		return ""
	}
	if namespace != "" {
		return namespace
	}
	return currentNamespace
}

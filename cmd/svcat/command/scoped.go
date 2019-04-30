package command

import (
	"fmt"
	"errors"
	"github.com/kubernetes-incubator/service-catalog/pkg/svcat/service-catalog"
	"github.com/spf13/pflag"
)

type HasScopedFlags interface {
	ApplyScopedFlags(flags *pflag.FlagSet) error
}

var _ HasScopedFlags = NewScoped()

type Scoped struct {
	allowAll	bool
	rawScope	string
	Scope		servicecatalog.Scope
}

func NewScoped() *Scoped {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Scoped{}
}
func (c *Scoped) AddScopedFlags(flags *pflag.FlagSet, allowAll bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.allowAll = allowAll
	if allowAll {
		flags.StringVar(&c.rawScope, "scope", servicecatalog.AllScope, "Limit the command to a particular scope: cluster, namespace or all")
	} else {
		flags.StringVar(&c.rawScope, "scope", servicecatalog.NamespaceScope, "Limit the command to a particular scope: cluster or namespace")
	}
}
func (c *Scoped) ApplyScopedFlags(flags *pflag.FlagSet) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch c.rawScope {
	case servicecatalog.AllScope:
		if !c.allowAll {
			return errors.New("invalid --scope (all), allowed values are: cluster, namespace")
		}
		c.Scope = servicecatalog.Scope(c.rawScope)
		return nil
	case servicecatalog.ClusterScope, servicecatalog.NamespaceScope:
		c.Scope = servicecatalog.Scope(c.rawScope)
		return nil
	default:
		return fmt.Errorf("invalid --scope (%s), allowed values are: all, cluster, namespace", c.rawScope)
	}
}

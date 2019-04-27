package broker

import (
	"fmt"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	"github.com/kubernetes-incubator/service-catalog/pkg/svcat/service-catalog"
	"github.com/spf13/cobra"
)

type syncCmd struct {
	*command.Namespaced
	*command.Scoped
	name	string
}

func NewSyncCmd(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	syncCmd := &syncCmd{Namespaced: command.NewNamespaced(cxt), Scoped: command.NewScoped()}
	rootCmd := &cobra.Command{Use: "broker NAME", Short: "Syncs service catalog for a service broker", Example: command.NormalizeExamples(`svcat sync broker asb`), PreRunE: command.PreRunE(syncCmd), RunE: command.RunE(syncCmd)}
	syncCmd.AddScopedFlags(rootCmd.Flags(), false)
	syncCmd.AddNamespaceFlags(rootCmd.Flags(), false)
	return rootCmd
}
func (c *syncCmd) Validate(args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 1 {
		return fmt.Errorf("a broker name is required")
	}
	c.name = args[0]
	return nil
}
func (c *syncCmd) Run() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.sync()
}
func (c *syncCmd) sync() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	scopeOpts := servicecatalog.ScopeOptions{Scope: c.Scope, Namespace: c.Namespace}
	const retries = 3
	err := c.App.Sync(c.name, scopeOpts, retries)
	if err != nil {
		return err
	}
	fmt.Fprintf(c.Output, "Synchronization requested for broker: %s\n", c.name)
	return nil
}

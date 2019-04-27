package broker

import (
	"fmt"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/svcat/service-catalog"
	"github.com/spf13/cobra"
)

type DeregisterCmd struct {
	*command.Namespaced
	*command.Scoped
	*command.Waitable
	BrokerName	string
}

func NewDeregisterCmd(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	deregisterCmd := &DeregisterCmd{Namespaced: command.NewNamespaced(cxt), Scoped: command.NewScoped(), Waitable: command.NewWaitable()}
	cmd := &cobra.Command{Use: "deregister NAME", Short: "Deregisters an existing broker with service catalog", Example: command.NormalizeExamples(`
		svcat deregister mysqlbroker
		svcat deregister mysqlbroker --namespace=mysqlnamespace
		svcat deregister mysqlclusterbroker --cluster
		`), PreRunE: command.PreRunE(deregisterCmd), RunE: command.RunE(deregisterCmd)}
	deregisterCmd.AddNamespaceFlags(cmd.Flags(), false)
	deregisterCmd.AddScopedFlags(cmd.Flags(), false)
	deregisterCmd.AddWaitFlags(cmd)
	return cmd
}
func (c *DeregisterCmd) Validate(args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) == 0 {
		return fmt.Errorf("a broker name is required")
	}
	c.BrokerName = args[0]
	return nil
}
func (c *DeregisterCmd) Run() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Deregister()
}
func (c *DeregisterCmd) Deregister() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scopeOptions := &servicecatalog.ScopeOptions{Namespace: c.Namespace, Scope: c.Scope}
	err := c.Context.App.Deregister(c.BrokerName, scopeOptions)
	if err != nil {
		return err
	}
	fmt.Fprintf(c.Context.Output, "Successfully removed broker %q\n", c.BrokerName)
	return nil
}

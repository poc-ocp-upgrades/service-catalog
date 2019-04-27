package broker

import (
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/output"
	"github.com/kubernetes-incubator/service-catalog/pkg/svcat/service-catalog"
	"github.com/spf13/cobra"
)

type getCmd struct {
	*command.Namespaced
	*command.Formatted
	*command.Scoped
	name	string
}

func NewGetCmd(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	getCmd := &getCmd{Namespaced: command.NewNamespaced(cxt), Formatted: command.NewFormatted(), Scoped: command.NewScoped()}
	cmd := &cobra.Command{Use: "brokers [NAME]", Aliases: []string{"broker", "brk"}, Short: "List brokers, optionally filtered by name, scope or namespace", Example: command.NormalizeExamples(`
  svcat get brokers
  svcat get brokers --scope=cluster
  svcat get brokers --scope=all
  svcat get broker minibroker
`), PreRunE: command.PreRunE(getCmd), RunE: command.RunE(getCmd)}
	getCmd.AddOutputFlags(cmd.Flags())
	getCmd.AddScopedFlags(cmd.Flags(), true)
	getCmd.AddNamespaceFlags(cmd.Flags(), true)
	return cmd
}
func (c *getCmd) Validate(args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) > 0 {
		c.name = args[0]
	}
	return nil
}
func (c *getCmd) Run() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.name == "" {
		return c.getAll()
	}
	return c.get()
}
func (c *getCmd) getAll() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	opts := servicecatalog.ScopeOptions{Namespace: c.Namespace, Scope: c.Scope}
	brokers, err := c.App.RetrieveBrokers(opts)
	if err != nil {
		return err
	}
	output.WriteBrokerList(c.Output, c.OutputFormat, brokers...)
	return nil
}
func (c *getCmd) get() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	broker, err := c.App.RetrieveBroker(c.name)
	if err != nil {
		return err
	}
	output.WriteBroker(c.Output, c.OutputFormat, *broker)
	return nil
}

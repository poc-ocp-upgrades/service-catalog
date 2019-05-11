package binding

import (
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/output"
	"github.com/spf13/cobra"
)

type getCmd struct {
	*command.Namespaced
	*command.Formatted
	name	string
}

func NewGetCmd(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	getCmd := &getCmd{Namespaced: command.NewNamespaced(cxt), Formatted: command.NewFormatted()}
	cmd := &cobra.Command{Use: "bindings [NAME]", Aliases: []string{"binding", "bnd"}, Short: "List bindings, optionally filtered by name or namespace", Example: command.NormalizeExamples(`
  svcat get bindings
  svcat get bindings --all-namespaces
  svcat get binding wordpress-mysql-binding
  svcat get binding -n ci concourse-postgres-binding
`), PreRunE: command.PreRunE(getCmd), RunE: command.RunE(getCmd)}
	getCmd.AddNamespaceFlags(cmd.Flags(), true)
	getCmd.AddOutputFlags(cmd.Flags())
	return cmd
}
func (c *getCmd) Validate(args []string) error {
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
	if c.name == "" {
		return c.getAll()
	}
	return c.get()
}
func (c *getCmd) getAll() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bindings, err := c.App.RetrieveBindings(c.Namespace)
	if err != nil {
		return err
	}
	output.WriteBindingList(c.Output, c.OutputFormat, bindings)
	return nil
}
func (c *getCmd) get() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	binding, err := c.App.RetrieveBinding(c.Namespace, c.name)
	if err != nil {
		return err
	}
	output.WriteBinding(c.Output, c.OutputFormat, *binding)
	return nil
}

package instance

import (
	"fmt"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/output"
	"github.com/spf13/cobra"
)

type describeCmd struct {
	*command.Namespaced
	name	string
}

func NewDescribeCmd(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	describeCmd := &describeCmd{Namespaced: command.NewNamespaced(cxt)}
	cmd := &cobra.Command{Use: "instance NAME", Aliases: []string{"instances", "inst"}, Short: "Show details of a specific instance", Example: command.NormalizeExamples(`
  svcat describe instance wordpress-mysql-instance
`), PreRunE: command.PreRunE(describeCmd), RunE: command.RunE(describeCmd)}
	describeCmd.AddNamespaceFlags(cmd.Flags(), false)
	return cmd
}
func (c *describeCmd) Validate(args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) == 0 {
		return fmt.Errorf("an instance name is required")
	}
	c.name = args[0]
	return nil
}
func (c *describeCmd) Run() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.describe()
}
func (c *describeCmd) describe() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instance, err := c.App.RetrieveInstance(c.Namespace, c.name)
	if err != nil {
		return err
	}
	output.WriteInstanceDetails(c.Output, instance)
	bindings, err := c.App.RetrieveBindingsByInstance(instance)
	if err != nil {
		return err
	}
	output.WriteAssociatedBindings(c.Output, bindings)
	return nil
}

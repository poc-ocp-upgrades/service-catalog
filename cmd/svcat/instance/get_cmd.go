package instance

import (
	"fmt"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/output"
	"github.com/spf13/cobra"
)

type getCmd struct {
	*command.Namespaced
	*command.Formatted
	*command.PlanFiltered
	*command.ClassFiltered
	name	string
}

func NewGetCmd(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	getCmd := &getCmd{Namespaced: command.NewNamespaced(cxt), Formatted: command.NewFormatted(), ClassFiltered: command.NewClassFiltered(), PlanFiltered: command.NewPlanFiltered()}
	cmd := &cobra.Command{Use: "instances [NAME]", Aliases: []string{"instance", "inst"}, Short: "List instances, optionally filtered by name", Example: command.NormalizeExamples(`
  svcat get instances
  svcat get instances --class redis
  svcat get instances --plan default
  svcat get instances --all-namespaces
  svcat get instance wordpress-mysql-instance
  svcat get instance -n ci concourse-postgres-instance
`), PreRunE: command.PreRunE(getCmd), RunE: command.RunE(getCmd)}
	getCmd.AddNamespaceFlags(cmd.Flags(), true)
	getCmd.AddOutputFlags(cmd.Flags())
	getCmd.AddClassFlag(cmd)
	getCmd.AddPlanFlag(cmd)
	return cmd
}
func (c *getCmd) Validate(args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) > 0 {
		c.name = args[0]
		if c.ClassFilter != "" {
			return fmt.Errorf("class filter is not supported when specifiying instance name")
		}
		if c.PlanFilter != "" {
			return fmt.Errorf("plan filter is not supported when specifiying instance name")
		}
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
	instances, err := c.App.RetrieveInstances(c.Namespace, c.ClassFilter, c.PlanFilter)
	if err != nil {
		return err
	}
	output.WriteInstanceList(c.Output, c.OutputFormat, instances)
	return nil
}
func (c *getCmd) get() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	instance, err := c.App.RetrieveInstance(c.Namespace, c.name)
	if err != nil {
		return err
	}
	output.WriteInstance(c.Output, c.OutputFormat, *instance)
	return nil
}

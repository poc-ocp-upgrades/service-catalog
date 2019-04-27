package class

import (
	"fmt"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/output"
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/svcat/service-catalog"
	"github.com/spf13/cobra"
)

type describeCmd struct {
	*command.Context
	lookupByKubeName	bool
	kubeName		string
	name			string
}

func NewDescribeCmd(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	describeCmd := &describeCmd{Context: cxt}
	cmd := &cobra.Command{Use: "class NAME", Aliases: []string{"classes", "cl"}, Short: "Show details of a specific class", Example: command.NormalizeExamples(`
  svcat describe class mysqldb
  svcat describe class --kube-name 997b8372-8dac-40ac-ae65-758b4a5075a5
`), PreRunE: command.PreRunE(describeCmd), RunE: command.RunE(describeCmd)}
	cmd.Flags().BoolVarP(&describeCmd.lookupByKubeName, "kube-name", "k", false, "Whether or not to get the class by its Kubernetes Name (the default is by external name)")
	return cmd
}
func (c *describeCmd) Validate(args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) == 0 {
		return fmt.Errorf("a class name or Kubernetes name is required")
	}
	if c.lookupByKubeName {
		c.kubeName = args[0]
	} else {
		c.name = args[0]
	}
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
	var class servicecatalog.Class
	var err error
	if c.lookupByKubeName {
		class, err = c.App.RetrieveClassByID(c.kubeName)
	} else {
		class, err = c.App.RetrieveClassByName(c.name, servicecatalog.ScopeOptions{Scope: servicecatalog.ClusterScope})
	}
	if err != nil {
		return err
	}
	output.WriteClassDetails(c.Output, class)
	opts := servicecatalog.ScopeOptions{Scope: servicecatalog.AllScope}
	plans, err := c.App.RetrievePlans(class.GetName(), opts)
	if err != nil {
		return err
	}
	output.WriteAssociatedPlans(c.Output, plans)
	return nil
}

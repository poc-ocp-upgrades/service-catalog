package class

import (
	"fmt"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/output"
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/svcat/service-catalog"
	"github.com/spf13/cobra"
)

type CreateCmd struct {
	*command.Namespaced
	*command.Scoped
	Name	string
	From	string
}

func NewCreateCmd(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	createCmd := &CreateCmd{Namespaced: command.NewNamespaced(cxt), Scoped: command.NewScoped()}
	cmd := &cobra.Command{Use: "class [NAME] --from [EXISTING_NAME]", Short: "Copies an existing class into a new user-defined cluster-scoped class", Example: command.NormalizeExamples(`
  svcat create class newclass --from mysqldb
  svcat create class newclass --from mysqldb --scope cluster
  svcat create class newclass --from mysqldb --scope namespace --namespace newnamespace
`), PreRunE: command.PreRunE(createCmd), RunE: command.RunE(createCmd)}
	cmd.Flags().StringVarP(&createCmd.From, "from", "f", "", "Name from an existing class that will be copied (Required)")
	cmd.MarkFlagRequired("from")
	createCmd.AddNamespaceFlags(cmd.Flags(), false)
	createCmd.AddScopedFlags(cmd.Flags(), false)
	return cmd
}
func (c *CreateCmd) Validate(args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) <= 0 {
		return fmt.Errorf("new class name should be provided")
	}
	c.Name = args[0]
	return nil
}
func (c *CreateCmd) Run() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	opts := servicecatalog.CreateClassFromOptions{Scope: c.Scope, Namespace: c.Namespace, Name: c.Name, From: c.From}
	createdClass, err := c.App.CreateClassFrom(opts)
	if err != nil {
		return err
	}
	output.WriteClassList(c.Output, output.FormatTable, createdClass)
	return nil
}

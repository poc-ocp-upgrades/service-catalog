package binding

import (
	"fmt"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/output"
	"github.com/spf13/cobra"
)

type describeCmd struct {
	*command.Namespaced
	name		string
	showSecrets	bool
}

func NewDescribeCmd(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	describeCmd := &describeCmd{Namespaced: command.NewNamespaced(cxt)}
	cmd := &cobra.Command{Use: "binding NAME", Aliases: []string{"bindings", "bnd"}, Short: "Show details of a specific binding", Example: command.NormalizeExamples(`svcat describe binding wordpress-mysql-binding`), PreRunE: command.PreRunE(describeCmd), RunE: command.RunE(describeCmd)}
	describeCmd.AddNamespaceFlags(cmd.Flags(), false)
	cmd.Flags().BoolVar(&describeCmd.showSecrets, "show-secrets", false, "Output the decoded secret values. By default only the length of the secret is displayed")
	return cmd
}
func (c *describeCmd) Validate(args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) == 0 {
		return fmt.Errorf("a binding name is required")
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
	binding, err := c.App.RetrieveBinding(c.Namespace, c.name)
	if err != nil {
		return err
	}
	output.WriteBindingDetails(c.Output, binding)
	secret, err := c.App.RetrieveSecretByBinding(binding)
	output.WriteAssociatedSecret(c.Output, secret, err, c.showSecrets)
	return nil
}

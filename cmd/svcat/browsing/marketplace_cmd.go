package browsing

import (
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/output"
	"github.com/kubernetes-incubator/service-catalog/pkg/svcat/service-catalog"
	"github.com/spf13/cobra"
)

type MarketplaceCmd struct {
	*command.Namespaced
	*command.Formatted
}

func NewMarketplaceCmd(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mpCmd := &MarketplaceCmd{Namespaced: command.NewNamespaced(cxt), Formatted: command.NewFormatted()}
	cmd := &cobra.Command{Use: "marketplace", Aliases: []string{"marketplace", "mp"}, Short: "List available service offerings", Example: command.NormalizeExamples(`
  svcat marketplace
	svcat marketplace --namespace dev
`), PreRunE: command.PreRunE(mpCmd), RunE: command.RunE(mpCmd)}
	mpCmd.AddOutputFlags(cmd.Flags())
	mpCmd.AddNamespaceFlags(cmd.Flags(), true)
	return cmd
}
func (c *MarketplaceCmd) Validate(args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (c *MarketplaceCmd) Run() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	opts := servicecatalog.ScopeOptions{Namespace: c.Namespace, Scope: servicecatalog.AllScope}
	classes, err := c.App.RetrieveClasses(opts)
	if err != nil {
		return err
	}
	plans := make([][]servicecatalog.Plan, len(classes))
	classPlans, err := c.App.RetrievePlans("", opts)
	if err != nil {
		return err
	}
	for i, class := range classes {
		for _, plan := range classPlans {
			if plan.GetClassID() == class.GetName() {
				plans[i] = append(plans[i], plan)
			}
		}
	}
	output.WriteClassAndPlanDetails(c.Output, classes, plans)
	return nil
}

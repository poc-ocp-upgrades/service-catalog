package command

import (
	"github.com/spf13/cobra"
)

type HasPlanFlag interface{ ApplyPlanFlag(*cobra.Command) error }
type PlanFiltered struct{ PlanFilter string }

func NewPlanFiltered() *PlanFiltered {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &PlanFiltered{}
}
func (c *PlanFiltered) AddPlanFlag(cmd *cobra.Command) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd.Flags().StringP("plan", "p", "", "If present, specify the plan used as a filter for this request")
}
func (c *PlanFiltered) ApplyPlanFlag(cmd *cobra.Command) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	c.PlanFilter, err = cmd.Flags().GetString("plan")
	return err
}

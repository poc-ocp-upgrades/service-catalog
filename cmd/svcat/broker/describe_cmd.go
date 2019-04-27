package broker

import (
	"fmt"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/output"
	"github.com/spf13/cobra"
)

type describeCmd struct {
	*command.Context
	name	string
}

func NewDescribeCmd(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	describeCmd := &describeCmd{Context: cxt}
	cmd := &cobra.Command{Use: "broker NAME", Aliases: []string{"brokers", "brk"}, Short: "Show details of a specific broker", Example: command.NormalizeExamples(`
  svcat describe broker asb
`), PreRunE: command.PreRunE(describeCmd), RunE: command.RunE(describeCmd)}
	return cmd
}
func (c *describeCmd) Validate(args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) == 0 {
		return fmt.Errorf("a broker name is required")
	}
	c.name = args[0]
	return nil
}
func (c *describeCmd) Run() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Describe()
}
func (c *describeCmd) Describe() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	broker, err := c.App.RetrieveBroker(c.name)
	if err != nil {
		return err
	}
	output.WriteBrokerDetails(c.Output, broker)
	return nil
}

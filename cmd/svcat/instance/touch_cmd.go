package instance

import (
	"fmt"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	"github.com/spf13/cobra"
)

type touchInstanceCmd struct {
	*command.Namespaced
	name	string
}

func NewTouchCommand(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	touchInstanceCmd := &touchInstanceCmd{Namespaced: command.NewNamespaced(cxt)}
	cmd := &cobra.Command{Use: "instance", Short: "Touch an instance to make service-catalog try to process the spec again", Long: `Touch instance will increment the updateRequests field on the instance. 
Then, service catalog will process the instance's spec again. It might do an update, a delete, or 
nothing.`, Example: command.NormalizeExamples(`svcat touch instance wordpress-mysql-instance --namespace mynamespace`), PreRunE: command.PreRunE(touchInstanceCmd), RunE: command.RunE(touchInstanceCmd)}
	touchInstanceCmd.AddNamespaceFlags(cmd.Flags(), false)
	return cmd
}
func (c *touchInstanceCmd) Validate(args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) == 0 {
		return fmt.Errorf("an instance name is required")
	}
	c.name = args[0]
	return nil
}
func (c *touchInstanceCmd) Run() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	const retries = 3
	return c.App.TouchInstance(c.Namespace, c.name, retries)
}

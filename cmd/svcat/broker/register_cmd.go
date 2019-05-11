package broker

import (
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/output"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/svcat/service-catalog"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RegisterCmd struct {
	*command.Namespaced
	*command.Scoped
	*command.Waitable
	BasicSecret			string
	BearerSecret		string
	BrokerName			string
	CAFile				string
	ClassRestrictions	[]string
	PlanRestrictions	[]string
	SkipTLS				bool
	RelistBehavior		string
	RelistDuration		time.Duration
	URL					string
}

func NewRegisterCmd(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	registerCmd := &RegisterCmd{Namespaced: command.NewNamespaced(cxt), Scoped: command.NewScoped(), Waitable: command.NewWaitable()}
	cmd := &cobra.Command{Use: "register NAME --url URL", Short: "Registers a new broker with service catalog", Example: command.NormalizeExamples(`
		svcat register mysqlbroker --url http://mysqlbroker.com
		`), PreRunE: command.PreRunE(registerCmd), RunE: command.RunE(registerCmd)}
	cmd.Flags().StringVar(&registerCmd.URL, "url", "", "The broker URL (Required)")
	cmd.MarkFlagRequired("url")
	cmd.Flags().StringVar(&registerCmd.BasicSecret, "basic-secret", "", "A secret containing basic auth (username/password) information to connect to the broker")
	cmd.Flags().StringVar(&registerCmd.BearerSecret, "bearer-secret", "", "A secret containing a bearer token to connect to the broker")
	cmd.Flags().StringVar(&registerCmd.CAFile, "ca", "", "A file containing the CA certificate to connect to the broker")
	cmd.Flags().StringSliceVar(&registerCmd.ClassRestrictions, "class-restrictions", []string{}, "A list of restrictions to apply to the classes allowed from the broker")
	cmd.Flags().StringSliceVar(&registerCmd.PlanRestrictions, "plan-restrictions", []string{}, "A list of restrictions to apply to the plans allowed from the broker")
	cmd.Flags().StringVar(&registerCmd.RelistBehavior, "relist-behavior", "", "Behavior for relisting the broker's catalog. Valid options are manual or duration. Defaults to duration with an interval of 15m.")
	cmd.Flags().DurationVar(&registerCmd.RelistDuration, "relist-duration", 0*time.Second, "Interval to refetch broker catalog when relist-behavior is set to duration, specified in human readable format: 30s, 1m, 1h")
	cmd.Flags().BoolVar(&registerCmd.SkipTLS, "skip-tls", false, "Disables TLS certificate verification when communicating with this broker. This is strongly discouraged. You should use --ca instead.")
	registerCmd.AddNamespaceFlags(cmd.Flags(), false)
	registerCmd.AddScopedFlags(cmd.Flags(), false)
	registerCmd.AddWaitFlags(cmd)
	return cmd
}
func (c *RegisterCmd) Validate(args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) < 1 {
		return fmt.Errorf("a broker name is required")
	}
	c.BrokerName = args[0]
	if c.BasicSecret != "" && c.BearerSecret != "" {
		return fmt.Errorf("cannot use both basic auth and bearer auth")
	}
	if c.CAFile != "" {
		_, err := os.Stat(c.CAFile)
		if err != nil {
			return fmt.Errorf("error finding CA file: %v", err.Error())
		}
	}
	if c.RelistBehavior != "" {
		c.RelistBehavior = strings.ToLower(c.RelistBehavior)
		if c.RelistBehavior != "duration" && c.RelistBehavior != "manual" {
			return fmt.Errorf("invalid --relist-duration value, allowed values are: duration, manual")
		}
	}
	return nil
}
func (c *RegisterCmd) Run() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	opts := &servicecatalog.RegisterOptions{BasicSecret: c.BasicSecret, BearerSecret: c.BearerSecret, CAFile: c.CAFile, ClassRestrictions: c.ClassRestrictions, Namespace: c.Namespace, PlanRestrictions: c.PlanRestrictions, SkipTLS: c.SkipTLS}
	scopeOpts := &servicecatalog.ScopeOptions{Namespace: c.Namespace, Scope: c.Scope}
	if c.RelistBehavior == "duration" {
		opts.RelistBehavior = v1beta1.ServiceBrokerRelistBehaviorDuration
		opts.RelistDuration = &metav1.Duration{Duration: c.RelistDuration}
	} else if c.RelistBehavior == "manual" {
		opts.RelistBehavior = v1beta1.ServiceBrokerRelistBehaviorManual
	}
	broker, err := c.Context.App.Register(c.BrokerName, c.URL, opts, scopeOpts)
	if err != nil {
		return err
	}
	if c.Wait {
		fmt.Fprintln(c.Output, "Waiting for the broker to be registered...")
		finalBroker, err := c.Context.App.WaitForBroker(c.BrokerName, c.Interval, c.Timeout)
		if err == nil {
			broker = finalBroker.(*v1beta1.ClusterServiceBroker)
		}
		output.WriteBrokerDetails(c.Output, broker)
		return err
	}
	output.WriteBrokerDetails(c.Context.Output, broker)
	return nil
}

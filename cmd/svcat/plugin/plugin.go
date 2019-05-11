package plugin

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	Name						= "svcat"
	EnvPluginCaller				= "KUBECTL_PLUGINS_CALLER"
	EnvPluginLocalFlagPrefix	= "KUBECTL_PLUGINS_LOCAL_FLAG"
	EnvPluginNamespace			= "KUBECTL_PLUGINS_CURRENT_NAMESPACE"
	EnvPluginGlobalFlagPrefix	= "KUBECTL_PLUGINS_GLOBAL_FLAG"
	EnvPluginVerbose			= EnvPluginGlobalFlagPrefix + "_V"
	EnvPluginPath				= "KUBECTL_PLUGINS_PATH"
)

func IsPlugin() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := os.LookupEnv(EnvPluginCaller)
	return ok
}
func BindEnvironmentVariables(vip *viper.Viper, cmd *cobra.Command) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	vip.BindEnv("namespace", EnvPluginNamespace)
	vip.BindEnv("v", EnvPluginVerbose)
	vip.SetEnvPrefix(EnvPluginLocalFlagPrefix)
	vip.BindPFlags(cmd.Flags())
	vip.AutomaticEnv()
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if !f.Changed && vip.IsSet(f.Name) {
			cmd.Flags().Set(f.Name, vip.GetString(f.Name))
		}
	})
}

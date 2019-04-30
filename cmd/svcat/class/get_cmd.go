package class

import (
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/output"
	"github.com/kubernetes-incubator/service-catalog/pkg/svcat/service-catalog"
	"github.com/spf13/cobra"
)

type getCmd struct {
	*command.Namespaced
	*command.Scoped
	*command.Formatted
	lookupByKubeName	bool
	kubeName		string
	name			string
}

func NewGetCmd(cxt *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	getCmd := &getCmd{Namespaced: command.NewNamespaced(cxt), Scoped: command.NewScoped(), Formatted: command.NewFormatted()}
	cmd := &cobra.Command{Use: "classes [NAME]", Aliases: []string{"class", "cl"}, Short: "List classes, optionally filtered by name, scope or namespace", Example: command.NormalizeExamples(`
  svcat get classes
  svcat get classes --scope cluster
  svcat get classes --scope namespace --namespace dev
  svcat get class mysqldb
  svcat get class --kube-name 997b8372-8dac-40ac-ae65-758b4a5075a5
`), PreRunE: command.PreRunE(getCmd), RunE: command.RunE(getCmd)}
	cmd.Flags().BoolVarP(&getCmd.lookupByKubeName, "kube-name", "k", false, "Whether or not to get the class by its Kubernetes name (the default is by external name)")
	getCmd.AddOutputFlags(cmd.Flags())
	getCmd.AddNamespaceFlags(cmd.Flags(), true)
	getCmd.AddScopedFlags(cmd.Flags(), true)
	return cmd
}
func (c *getCmd) Validate(args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) > 0 {
		if c.lookupByKubeName {
			c.kubeName = args[0]
		} else {
			c.name = args[0]
		}
	}
	return nil
}
func (c *getCmd) Run() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.kubeName == "" && c.name == "" {
		return c.getAll()
	}
	return c.get()
}
func (c *getCmd) getAll() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	opts := servicecatalog.ScopeOptions{Namespace: c.Namespace, Scope: c.Scope}
	classes, err := c.App.RetrieveClasses(opts)
	if err != nil {
		return err
	}
	output.WriteClassList(c.Output, c.OutputFormat, classes...)
	return nil
}
func (c *getCmd) get() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var class servicecatalog.Class
	var err error
	if c.lookupByKubeName {
		class, err = c.App.RetrieveClassByID(c.kubeName)
	} else if c.name != "" {
		class, err = c.App.RetrieveClassByName(c.name, servicecatalog.ScopeOptions{Scope: c.Scope, Namespace: c.Namespace})
	}
	if err != nil {
		return err
	}
	output.WriteClass(c.Output, c.OutputFormat, class)
	return nil
}

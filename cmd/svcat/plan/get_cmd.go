package plan

import (
	"fmt"
	"strings"
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
	classFilter		string
	classKubeName		string
	className		string
}

func NewGetCmd(ctx *command.Context) *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	getCmd := &getCmd{Namespaced: command.NewNamespaced(ctx), Scoped: command.NewScoped(), Formatted: command.NewFormatted()}
	cmd := &cobra.Command{Use: "plans [NAME]", Aliases: []string{"plan", "pl"}, Short: "List plans, optionally filtered by name, class, scope or namespace", Example: command.NormalizeExamples(`
  svcat get plans
  svcat get plans --scope cluster
  svcat get plans --scope namespace --namespace dev
  svcat get plan PLAN_NAME
  svcat get plan CLASS_NAME/PLAN_NAME
  svcat get plan --kube-name PLAN_KUBE_NAME
  svcat get plans --class CLASS_NAME
  svcat get plan --class CLASS_NAME PLAN_NAME
  svcat get plans --kube-name --class CLASS_KUBE_NAME
  svcat get plan --kube-name --class CLASS_KUBE_NAME PLAN_KUBE_NAME
`), PreRunE: command.PreRunE(getCmd), RunE: command.RunE(getCmd)}
	cmd.Flags().BoolVarP(&getCmd.lookupByKubeName, "kube-name", "k", false, "Whether or not to get the plan by its Kubernetes name (the default is by external name)")
	cmd.Flags().StringVarP(&getCmd.classFilter, "class", "c", "", "Filter plans based on class. When --kube-name is specified, the class name is interpreted as a kubernetes name.")
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
		} else if strings.Contains(args[0], "/") {
			names := strings.Split(args[0], "/")
			if len(names) != 2 {
				return fmt.Errorf("failed to parse class/plan name combination '%s'", c.name)
			}
			c.className = names[0]
			c.name = names[1]
		} else {
			c.name = args[0]
		}
	}
	if c.classFilter != "" {
		if c.lookupByKubeName {
			c.classKubeName = c.classFilter
		} else {
			c.className = c.classFilter
		}
	}
	return nil
}
func (c *getCmd) Run() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Println("KUBENAME: ", c.kubeName)
	fmt.Println("EXTERNAL NAME: ", c.name)
	fmt.Println("CLASS EXTERNAL NAME: ", c.className)
	fmt.Println("CLASS NAME: ", c.classKubeName)
	if c.kubeName == "" && c.name == "" {
		return c.getAll()
	}
	return c.get()
}
func (c *getCmd) getAll() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	classOpts := servicecatalog.ScopeOptions{Namespace: c.Namespace, Scope: c.Scope}
	classes, err := c.App.RetrieveClasses(classOpts)
	if err != nil {
		return fmt.Errorf("unable to list classes (%s)", err)
	}
	var classID string
	opts := servicecatalog.ScopeOptions{Namespace: c.Namespace, Scope: c.Scope}
	if c.classFilter != "" {
		if !c.lookupByKubeName {
			for _, class := range classes {
				if c.className == class.GetExternalName() {
					c.classKubeName = class.GetName()
					break
				}
			}
		}
		classID = c.classKubeName
	}
	plans, err := c.App.RetrievePlans(classID, opts)
	fmt.Println("PLANS: ", plans)
	if err != nil {
		return fmt.Errorf("unable to list plans (%s)", err)
	}
	output.WritePlanList(c.Output, c.OutputFormat, plans, classes)
	return nil
}
func (c *getCmd) get() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var plan servicecatalog.Plan
	var err error
	opts := servicecatalog.ScopeOptions{Namespace: c.Namespace, Scope: c.Scope}
	switch {
	case c.lookupByKubeName:
		plan, err = c.App.RetrievePlanByID(c.kubeName, opts)
	case c.className != "":
		plan, err = c.App.RetrievePlanByClassAndName(c.className, c.name, opts)
	default:
		plan, err = c.App.RetrievePlanByName(c.name, opts)
	}
	if err != nil {
		return err
	}
	class, err := c.App.RetrieveClassByID(plan.GetClassID())
	if err != nil {
		return err
	}
	output.WritePlan(c.Output, c.OutputFormat, plan, *class)
	return nil
}

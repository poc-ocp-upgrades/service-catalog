package command

import (
	"fmt"
	"strings"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/output"
	"github.com/spf13/pflag"
)

type HasFormatFlags interface {
	ApplyFormatFlags(lags *pflag.FlagSet) error
}
type Formatted struct{ OutputFormat string }

func NewFormatted() *Formatted {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Formatted{OutputFormat: output.FormatTable}
}
func (c *Formatted) AddOutputFlags(flags *pflag.FlagSet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	flags.StringVarP(&c.OutputFormat, "output", "o", output.FormatTable, "The output format to use. Valid options are table, json or yaml. If not present, defaults to table")
}
func (c *Formatted) ApplyFormatFlags(flags *pflag.FlagSet) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.OutputFormat = strings.ToLower(c.OutputFormat)
	switch c.OutputFormat {
	case output.FormatTable, output.FormatJSON, output.FormatYAML:
		return nil
	default:
		return fmt.Errorf("invalid --output format %q, allowed values are: table, json and yaml", c.OutputFormat)
	}
}

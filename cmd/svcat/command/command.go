package command

import (
	"strings"
	"unicode"
	"fmt"
	"github.com/spf13/cobra"
)

type Command interface {
	Validate(args []string) error
	Run() error
}

func PreRunE(cmd Command) func(*cobra.Command, []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(c *cobra.Command, args []string) error {
		if nsCmd, ok := cmd.(HasNamespaceFlags); ok {
			nsCmd.ApplyNamespaceFlags(c.Flags())
		}
		if scopedCmd, ok := cmd.(HasScopedFlags); ok {
			err := scopedCmd.ApplyScopedFlags(c.Flags())
			if err != nil {
				return err
			}
		}
		if fmtCmd, ok := cmd.(HasFormatFlags); ok {
			err := fmtCmd.ApplyFormatFlags(c.Flags())
			if err != nil {
				return err
			}
		}
		if classFilteredCmd, ok := cmd.(HasClassFlag); ok {
			err := classFilteredCmd.ApplyClassFlag(c)
			if err != nil {
				return err
			}
		}
		if planFilteredCmd, ok := cmd.(HasPlanFlag); ok {
			err := planFilteredCmd.ApplyPlanFlag(c)
			if err != nil {
				return err
			}
		}
		if waitCmd, ok := cmd.(HasWaitFlags); ok {
			err := waitCmd.ApplyWaitFlags()
			if err != nil {
				return err
			}
		}
		err := cmd.Validate(args)
		if err != nil {
			fmt.Fprintln(c.OutOrStderr(), err)
			fmt.Fprintln(c.OutOrStdout(), c.UsageString())
		}
		return err
	}
}
func RunE(cmd Command) func(*cobra.Command, []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(_ *cobra.Command, args []string) error {
		return cmd.Run()
	}
}
func NormalizeExamples(examples string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	indentedLines := []string{}
	var baseIndentation *string
	for _, line := range strings.Split(examples, "\n") {
		if baseIndentation == nil {
			if len(strings.TrimSpace(line)) == 0 {
				continue
			}
			whitespaceAtFront := line[:strings.Index(line, strings.TrimSpace(line))]
			baseIndentation = &whitespaceAtFront
		}
		trimmed := strings.TrimPrefix(line, *baseIndentation)
		indented := "  " + trimmed
		indentedLines = append(indentedLines, indented)
	}
	indentedString := strings.Join(indentedLines, "\n")
	return strings.TrimRightFunc(indentedString, unicode.IsSpace)
}

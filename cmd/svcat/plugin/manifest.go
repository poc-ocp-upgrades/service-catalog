package plugin

import (
	"strings"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var reservedFlags = map[string]struct{}{"alsologtostderr": {}, "as": {}, "as-group": {}, "cache-dir": {}, "certificate-authority": {}, "client-certificate": {}, "client-key": {}, "cluster": {}, "context": {}, "help": {}, "insecure-skip-tls-verify": {}, "kubeconfig": {}, "log-backtrace-at": {}, "log-dir": {}, "log-flush-frequency": {}, "logtostderr": {}, "match-server-version": {}, "n": {}, "namespace": {}, "password": {}, "request-timeout": {}, "s": {}, "server": {}, "stderrthreshold": {}, "token": {}, "user": {}, "username": {}, "v": {}, "vmodule": {}}
var commandsToSkip = map[string]struct{}{"svcat install": {}}

type Manifest struct {
	Plugin `json:",inline"`
}
type Plugin struct {
	Name		string		`json:"name"`
	Use		string		`json:"use"`
	ShortDesc	string		`json:"shortDesc"`
	LongDesc	string		`json:"longDesc,omitempty"`
	Example		string		`json:"example,omitempty"`
	Command		string		`json:"command"`
	Flags		[]Flag		`json:"flags,omitempty"`
	Tree		[]Plugin	`json:"tree,omitempty"`
}
type Flag struct {
	Name		string	`json:"name"`
	Shorthand	string	`json:"shorthand,omitempty"`
	Desc		string	`json:"desc"`
	DefValue	string	`json:"defValue,omitempty"`
}

func (m *Manifest) Load(rootCmd *cobra.Command) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.Plugin = m.convertToPlugin(rootCmd)
}
func (m *Manifest) convertToPlugin(cmd *cobra.Command) Plugin {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p := Plugin{}
	p.Name = strings.Split(cmd.Use, " ")[0]
	p.Use = cmd.Use
	p.ShortDesc = cmd.Short
	if p.ShortDesc == "" {
		p.ShortDesc = " "
	}
	p.LongDesc = cmd.Long
	p.Command = "./" + cmd.CommandPath()
	p.Example = cmd.Example
	p.Flags = []Flag{}
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		result := m.convertToFlag(flag)
		if result != nil {
			p.Flags = append(p.Flags, *result)
		}
	})
	p.Tree = []Plugin{}
	for _, subCmd := range cmd.Commands() {
		if _, skip := commandsToSkip[subCmd.CommandPath()]; !skip {
			p.Tree = append(p.Tree, m.convertToPlugin(subCmd))
		}
	}
	return p
}
func (m *Manifest) convertToFlag(src *pflag.Flag) *Flag {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, reserved := reservedFlags[src.Name]; reserved {
		return nil
	}
	dest := &Flag{Name: src.Name, Desc: src.Usage}
	if _, reserved := reservedFlags[src.Shorthand]; !reserved {
		dest.Shorthand = src.Shorthand
	}
	return dest
}

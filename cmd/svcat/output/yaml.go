package output

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

func writeYAML(w io.Writer, obj interface{}, n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	yBytes, err := yaml.Marshal(obj)
	if err != nil {
		fmt.Fprintf(w, "err marshaling yaml: %v\n", err)
		return
	}
	y := string(yBytes)
	if n > 0 {
		indent := strings.Repeat(" ", n)
		y = indent + strings.Replace(y, "\n", "\n"+indent, -1)
		y = strings.TrimRight(y, " ")
	}
	fmt.Fprint(w, y)
}
func writeParameters(w io.Writer, parameters *runtime.RawExtension) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Fprintln(w, "\nParameters:")
	if parameters == nil || string(parameters.Raw) == "" || string(parameters.Raw) == "{}" {
		fmt.Fprintln(w, "  No parameters defined")
		return
	}
	var params map[string]interface{}
	err := json.Unmarshal(parameters.Raw, &params)
	if err != nil {
		fmt.Fprintln(w, string(parameters.Raw))
	} else {
		writeYAML(w, params, 2)
	}
}
func writeParametersFrom(w io.Writer, parametersFrom []v1beta1.ParametersFromSource) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(parametersFrom) == 0 {
		return
	}
	headerPrinted := false
	for _, p := range parametersFrom {
		if p.SecretKeyRef != nil {
			if !headerPrinted {
				fmt.Fprintln(w, "\nParameters From:")
				headerPrinted = true
			}
			fmt.Fprintf(w, "  Secret: %s.%s\n", p.SecretKeyRef.Name, p.SecretKeyRef.Key)
		}
	}
}

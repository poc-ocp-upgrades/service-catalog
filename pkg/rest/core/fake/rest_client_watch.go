package fake

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/testapi"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	pkgwatch "k8s.io/apimachinery/pkg/watch"
)

func doWatch(ch <-chan pkgwatch.Event, w http.ResponseWriter) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for evt := range ch {
		codec, err := testapi.GetCodecForObject(evt.Object)
		if err != nil {
			errStr := fmt.Sprintf("error getting codec (%s)", err)
			log.Fatal(errStr)
			http.Error(w, errStr, http.StatusInternalServerError)
			return
		}
		objBytes, err := runtime.Encode(codec, evt.Object)
		if err != nil {
			errStr := fmt.Sprintf("error encoding item (%s)", err)
			log.Fatal(errStr)
			http.Error(w, errStr, http.StatusInternalServerError)
			return
		}
		evt := metav1.WatchEvent{Type: fmt.Sprintf("%s", evt.Type), Object: runtime.RawExtension{Object: evt.Object, Raw: objBytes}}
		b, err := json.Marshal(&evt)
		if err != nil {
			errStr := fmt.Sprintf("error encoding JSON (%s)", err)
			log.Fatal(errStr)
			http.Error(w, errStr, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json, */*")
		w.Write(b)
	}
}

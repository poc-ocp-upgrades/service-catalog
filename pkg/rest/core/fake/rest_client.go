package fake

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
	scmeta "github.com/kubernetes-incubator/service-catalog/pkg/api/meta"
	sc "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/testapi"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	fakerestclient "k8s.io/client-go/rest/fake"
)

var (
	accessor = meta.NewAccessor()
)

type ObjStorage map[string]runtime.Object
type TypedStorage map[string]ObjStorage
type NamespacedStorage map[string]TypedStorage

func NewTypedStorage() TypedStorage {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return map[string]ObjStorage{}
}
func (s NamespacedStorage) Set(ns, tipe, name string, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, ok := s[ns]; !ok {
		s[ns] = make(TypedStorage)
	}
	if _, ok := s[ns][tipe]; !ok {
		s[ns][tipe] = make(ObjStorage)
	}
	s[ns][tipe][name] = obj
}
func (s NamespacedStorage) GetList(ns, tipe string) []runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	itemMap, ok := s[ns][tipe]
	if !ok {
		return []runtime.Object{}
	}
	items := make([]runtime.Object, 0, len(itemMap))
	for _, item := range itemMap {
		items = append(items, item)
	}
	return items
}
func (s NamespacedStorage) Get(ns, tipe, name string) runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	item, ok := s[ns][tipe][name]
	if !ok {
		return nil
	}
	return item
}
func (s NamespacedStorage) Delete(ns, tipe, name string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	delete(s[ns][tipe], name)
}

type RESTClient struct {
	Storage		NamespacedStorage
	Watcher		*Watcher
	accessor	meta.MetadataAccessor
	*fakerestclient.RESTClient
}

func NewRESTClient(newEmptyObj func() runtime.Object) *RESTClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	storage := make(NamespacedStorage)
	watcher := NewWatcher()
	coreCl := &fakerestclient.RESTClient{Client: fakerestclient.CreateHTTPClient(func(request *http.Request) (*http.Response, error) {
		r := getRouter(storage, watcher, newEmptyObj)
		rw := newResponseWriter()
		r.ServeHTTP(rw, request)
		return rw.getResponse(), nil
	}), NegotiatedSerializer: serializer.DirectCodecFactory{CodecFactory: Codecs}}
	return &RESTClient{Storage: storage, Watcher: watcher, accessor: meta.NewAccessor(), RESTClient: coreCl}
}

type responseWriter struct {
	header		http.Header
	headerSet	bool
	body		[]byte
}

func newResponseWriter() *responseWriter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &responseWriter{header: make(http.Header)}
}
func (rw *responseWriter) Header() http.Header {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rw.header
}
func (rw *responseWriter) Write(bytes []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !rw.headerSet {
		rw.WriteHeader(http.StatusOK)
	}
	rw.body = append(rw.body, bytes...)
	return len(bytes), nil
}
func (rw *responseWriter) WriteHeader(status int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rw.headerSet = true
	rw.header.Set("status", strconv.Itoa(status))
}
func (rw *responseWriter) getResponse() *http.Response {
	_logClusterCodePath()
	defer _logClusterCodePath()
	status, err := strconv.ParseInt(rw.header.Get("status"), 10, 16)
	if err != nil {
		panic(err)
	}
	return &http.Response{StatusCode: int(status), Header: rw.header, Body: ioutil.NopCloser(bytes.NewBuffer(rw.body))}
}
func getRouter(storage NamespacedStorage, watcher *Watcher, newEmptyObj func() runtime.Object) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := mux.NewRouter()
	r.StrictSlash(true)
	r.HandleFunc("/apis/servicecatalog.k8s.io/v1beta1/namespaces/{namespace}/{type}", getItems(storage)).Methods("GET")
	r.HandleFunc("/apis/servicecatalog.k8s.io/v1beta1/namespaces/{namespace}/{type}", createItem(storage, newEmptyObj)).Methods("POST")
	r.HandleFunc("/apis/servicecatalog.k8s.io/v1beta1/namespaces/{namespace}/{type}/{name}", getItem(storage)).Methods("GET")
	r.HandleFunc("/apis/servicecatalog.k8s.io/v1beta1/namespaces/{namespace}/{type}/{name}", updateItem(storage, newEmptyObj)).Methods("PUT")
	r.HandleFunc("/apis/servicecatalog.k8s.io/v1beta1/namespaces/{namespace}/{type}/{name}", deleteItem(storage)).Methods("DELETE")
	r.HandleFunc("/apis/servicecatalog.k8s.io/v1beta1/watch/namespaces/{namespace}/{type}/{name}", watchItem(watcher)).Methods("GET")
	r.HandleFunc("/apis/servicecatalog.k8s.io/v1beta1/watch/namespaces/{namespace}/{type}", watchList(watcher)).Methods("GET")
	r.HandleFunc("/api/v1/namespaces", listNamespaces(storage)).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	return r
}
func watchItem(watcher *Watcher) func(http.ResponseWriter, *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(w http.ResponseWriter, r *http.Request) {
		ch := watcher.ReceiveChan()
		doWatch(ch, w)
	}
}
func watchList(watcher *Watcher) func(http.ResponseWriter, *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	const timeout = 1 * time.Second
	return func(w http.ResponseWriter, r *http.Request) {
		ch := watcher.ReceiveChan()
		doWatch(ch, w)
	}
}
func getItems(storage NamespacedStorage) func(http.ResponseWriter, *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(rw http.ResponseWriter, r *http.Request) {
		ns := mux.Vars(r)["namespace"]
		tipe := mux.Vars(r)["type"]
		objs := storage.GetList(ns, tipe)
		items := make([]runtime.Object, 0, len(objs))
		for _, obj := range objs {
			item := obj.DeepCopyObject()
			items = append(items, item)
		}
		var list runtime.Object
		var codec runtime.Codec
		var err error
		switch tipe {
		case "clusterservicebrokers":
			list = &sc.ClusterServiceBrokerList{TypeMeta: newTypeMeta("cluster-broker-list")}
			if err := meta.SetList(list, items); err != nil {
				errStr := fmt.Sprintf("Error setting list items (%s)", err)
				http.Error(rw, errStr, http.StatusInternalServerError)
				return
			}
			codec, err = testapi.GetCodecForObject(&sc.ClusterServiceBrokerList{})
		case "clusterserviceclasses":
			list = &sc.ClusterServiceClassList{TypeMeta: newTypeMeta("service-class-list")}
			if err := meta.SetList(list, items); err != nil {
				errStr := fmt.Sprintf("Error setting list items (%s)", err)
				http.Error(rw, errStr, http.StatusInternalServerError)
				return
			}
			codec, err = testapi.GetCodecForObject(&sc.ClusterServiceClassList{})
		case "serviceinstances":
			list = &sc.ServiceInstanceList{TypeMeta: newTypeMeta("instance-list")}
			if err := meta.SetList(list, items); err != nil {
				errStr := fmt.Sprintf("Error setting list items (%s)", err)
				http.Error(rw, errStr, http.StatusInternalServerError)
				return
			}
			codec, err = testapi.GetCodecForObject(&sc.ServiceInstanceList{})
		case "servicebindings":
			list = &sc.ServiceBindingList{TypeMeta: newTypeMeta("binding-list")}
			if err := meta.SetList(list, items); err != nil {
				errStr := fmt.Sprintf("Error setting list items (%s)", err)
				http.Error(rw, errStr, http.StatusInternalServerError)
				return
			}
			codec, err = testapi.GetCodecForObject(&sc.ServiceBindingList{})
		default:
			errStr := fmt.Sprintf("unrecognized resource type: %s", tipe)
			http.Error(rw, errStr, http.StatusInternalServerError)
			return
		}
		if err != nil {
			errStr := fmt.Sprintf("error getting codec: %s", err)
			http.Error(rw, errStr, http.StatusInternalServerError)
			return
		}
		listBytes, err := runtime.Encode(codec, list)
		if err != nil {
			errStr := fmt.Sprintf("error encoding list: %s", err)
			http.Error(rw, errStr, http.StatusInternalServerError)
			return
		}
		rw.Write(listBytes)
	}
}
func createItem(storage NamespacedStorage, newEmptyObj func() runtime.Object) func(rw http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(rw http.ResponseWriter, r *http.Request) {
		ns := mux.Vars(r)["namespace"]
		tipe := mux.Vars(r)["type"]
		codec, err := testapi.GetCodecForObject(newEmptyObj())
		if err != nil {
			errStr := fmt.Sprintf("error getting a codec for %#v (%s)", newEmptyObj(), err)
			http.Error(rw, errStr, http.StatusInternalServerError)
			return
		}
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errStr := fmt.Sprintf("error getting body bytes: %s", err)
			http.Error(rw, errStr, http.StatusInternalServerError)
			return
		}
		item, err := runtime.Decode(codec, bodyBytes)
		if err != nil {
			errStr := fmt.Sprintf("error decoding body bytes: %s", err)
			http.Error(rw, errStr, http.StatusInternalServerError)
			return
		}
		name, err := accessor.Name(item)
		if err != nil {
			errStr := fmt.Sprintf("couldn't get object name: %s", err)
			http.Error(rw, errStr, http.StatusInternalServerError)
			return
		}
		if storage.Get(ns, tipe, name) != nil {
			rw.WriteHeader(http.StatusConflict)
			return
		}
		accessor.SetResourceVersion(item, "1")
		storage.Set(ns, tipe, name, item)
		setContentType(rw)
		rw.WriteHeader(http.StatusCreated)
		bytes, err := runtime.Encode(codec, item)
		if err != nil {
			errStr := fmt.Sprintf("error encoding item (%s)", err)
			http.Error(rw, errStr, http.StatusInternalServerError)
			return
		}
		rw.Write(bytes)
	}
}
func getItem(storage NamespacedStorage) func(http.ResponseWriter, *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(rw http.ResponseWriter, r *http.Request) {
		ns := mux.Vars(r)["namespace"]
		tipe := mux.Vars(r)["type"]
		name := mux.Vars(r)["name"]
		item := storage.Get(ns, tipe, name)
		if item == nil {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		codec, err := testapi.GetCodecForObject(item)
		if err != nil {
			errStr := fmt.Sprintf("error getting codec (%s)", err)
			http.Error(rw, errStr, http.StatusInternalServerError)
			return
		}
		bytes, err := runtime.Encode(codec, item)
		if err != nil {
			errStr := fmt.Sprintf("error encoding item (%s)", err)
			http.Error(rw, errStr, http.StatusInternalServerError)
			return
		}
		setContentType(rw)
		rw.Write(bytes)
	}
}
func updateItem(storage NamespacedStorage, newEmptyObj func() runtime.Object) func(http.ResponseWriter, *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(rw http.ResponseWriter, r *http.Request) {
		ns := mux.Vars(r)["namespace"]
		tipe := mux.Vars(r)["type"]
		name := mux.Vars(r)["name"]
		origItem := storage.Get(ns, tipe, name)
		if origItem == nil {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		codec, err := testapi.GetCodecForObject(newEmptyObj())
		if err != nil {
			errStr := fmt.Sprintf("error getting codec: %s", err)
			http.Error(rw, errStr, http.StatusInternalServerError)
			return
		}
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			errStr := fmt.Sprintf("error getting body bytes: %s", err)
			http.Error(rw, errStr, http.StatusInternalServerError)
			return
		}
		item, err := runtime.Decode(codec, bodyBytes)
		if err != nil {
			errStr := fmt.Sprintf("error decoding body bytes: %s", err)
			http.Error(rw, errStr, http.StatusInternalServerError)
			return
		}
		origResourceVersionStr, err := accessor.ResourceVersion(item)
		if err != nil {
			errStr := fmt.Sprintf("error getting resource version")
			http.Error(rw, errStr, http.StatusInternalServerError)
			return
		}
		resourceVersionStr, err := accessor.ResourceVersion(item)
		if err != nil {
			errStr := fmt.Sprintf("error getting resource version")
			http.Error(rw, errStr, http.StatusInternalServerError)
			return
		}
		if resourceVersionStr != "0" && resourceVersionStr != origResourceVersionStr {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		resourceVersion, err := strconv.Atoi(origResourceVersionStr)
		resourceVersion++
		accessor.SetResourceVersion(item, strconv.Itoa(resourceVersion))
		finalizers, err := scmeta.GetFinalizers(item)
		if err != nil {
			errStr := fmt.Sprintf("error getting finalizers (%s)", err)
			http.Error(rw, errStr, http.StatusInternalServerError)
			return
		}
		oldDT, err := scmeta.GetDeletionTimestamp(origItem)
		if err != nil && err != scmeta.ErrNoDeletionTimestamp {
			errStr := fmt.Sprintf("error getting deletion timestamp on existing obj (%s)", err)
			http.Error(rw, errStr, http.StatusInternalServerError)
			return
		}
		newDT, err := scmeta.GetDeletionTimestamp(item)
		if err != nil && err != scmeta.ErrNoDeletionTimestamp {
			errStr := fmt.Sprintf("error getting deletion timestamp on new obj (%s)", err)
			http.Error(rw, errStr, http.StatusInternalServerError)
			return
		}
		if newDT != nil && oldDT != nil {
			if newDT.Unix() != oldDT.Unix() {
				errStr := fmt.Sprintf("you cannot update the deletion timestamp (old: %#v, new: %#v)", oldDT.String(), newDT.String())
				http.Error(rw, errStr, http.StatusBadRequest)
				return
			}
		}
		if len(finalizers) == 0 && newDT != nil {
			storage.Delete(ns, tipe, name)
		} else {
			storage.Set(ns, tipe, name, item)
		}
		bytes, err := runtime.Encode(codec, item)
		if err != nil {
			errStr := fmt.Sprintf("error encoding item: %s", err)
			http.Error(rw, errStr, http.StatusInternalServerError)
			return
		}
		setContentType(rw)
		rw.Write(bytes)
	}
}
func deleteItem(storage NamespacedStorage) func(http.ResponseWriter, *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(rw http.ResponseWriter, r *http.Request) {
		ns := mux.Vars(r)["namespace"]
		tipe := mux.Vars(r)["type"]
		name := mux.Vars(r)["name"]
		item := storage.Get(ns, tipe, name)
		if item == nil {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		finalizers, err := scmeta.GetFinalizers(item)
		if err != nil {
			http.Error(rw, fmt.Sprintf("error getting finalizers (%s)", err), http.StatusInternalServerError)
			return
		}
		if len(finalizers) == 0 {
			storage.Delete(ns, tipe, name)
		} else {
			if err := scmeta.SetDeletionTimestamp(item, time.Now()); err != nil {
				http.Error(rw, fmt.Sprintf("error setting deletion timestamp (%s)", err), http.StatusInternalServerError)
				return
			}
			storage.Set(ns, tipe, name, item)
		}
		rw.WriteHeader(http.StatusOK)
	}
}
func listNamespaces(storage NamespacedStorage) func(http.ResponseWriter, *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(rw http.ResponseWriter, r *http.Request) {
		nsList := corev1.NamespaceList{}
		for ns := range storage {
			nsList.Items = append(nsList.Items, corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns}})
		}
		setContentType(rw)
		if err := json.NewEncoder(rw).Encode(&nsList); err != nil {
			log.Printf("Error encoding namespace list (%s)", err)
			rw.WriteHeader(http.StatusInternalServerError)
		}
	}
}
func notFoundHandler(rw http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rw.WriteHeader(http.StatusNotFound)
}
func newTypeMeta(kind string) metav1.TypeMeta {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return metav1.TypeMeta{Kind: kind, APIVersion: sc.GroupName + "/v1beta1'"}
}

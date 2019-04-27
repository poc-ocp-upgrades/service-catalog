package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/kubernetes-incubator/service-catalog/contrib/pkg/broker/controller"
	"github.com/kubernetes-incubator/service-catalog/contrib/pkg/brokerapi"
)

var _ controller.Controller = &Controller{}

func TestCatalogReturnsHTTPErrorOnError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	handler := createHandler(&Controller{t: t, catalog: func() (*brokerapi.Catalog, error) {
		return nil, errors.New("Catalog retrieval error")
	}})
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, httptest.NewRequest("GET", "/v2/catalog", nil))
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected HTTP status http.StatusBadRequest (%d), got %d", http.StatusBadRequest, rr.Code)
	}
	if contentType := rr.Header().Get("content-type"); contentType != "application/json" {
		t.Errorf("Expected response content-type 'application/json', got '%s'", contentType)
	}
	if body := rr.Body.String(); body != `{"Error":"Catalog retrieval error"}` {
		t.Errorf("Expected structured error response; got '%s'", body)
	}
}
func TestCatalogReturnsCompliantJSON(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	handler := createHandler(&Controller{t: t, catalog: func() (*brokerapi.Catalog, error) {
		return &brokerapi.Catalog{Services: []*brokerapi.Service{{Name: "foo"}}}, nil
	}})
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, httptest.NewRequest("GET", "/v2/catalog", nil))
	if rr.Code != http.StatusOK {
		t.Errorf("Expected HTTP status http.StatusOK (%d), got %d", http.StatusOK, rr.Code)
	}
	if contentType := rr.Header().Get("content-type"); contentType != "application/json" {
		t.Errorf("Expected response content-type 'application/json', got '%s'", contentType)
	}
	catalog, err := readJSON(rr)
	if err != nil {
		t.Errorf("Failed to parse JSON response with error %v", err)
	}
	if len(catalog) != 1 {
		t.Errorf("Expected catalog to have 1 element, got %d", len(catalog))
	}
	if _, ok := catalog["services"]; !ok {
		t.Errorf("Expected catalog %v to contain key 'services'", catalog)
	}
	services := catalog["services"].([]interface{})
	if services == nil {
		t.Error("Expected 'services' property of the returned catalog to be not nil, got nil")
	}
	var service map[string]interface{}
	service = services[0].(map[string]interface{})
	if name, ok := service["name"]; !ok {
		t.Error("Returned service doesn't have a 'name' property.")
	} else if name != "foo" {
		t.Errorf("Expected returned service name to be 'foo', got '%s'", name)
	}
}
func readJSON(rr *httptest.ResponseRecorder) (map[string]interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var result map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &result)
	return result, err
}

type Controller struct {
	t				*testing.T
	catalog				func() (*brokerapi.Catalog, error)
	getServiceInstanceLastOperation	func(id string) (*brokerapi.LastOperationResponse, error)
	createServiceInstance		func(id string, req *brokerapi.CreateServiceInstanceRequest) (*brokerapi.CreateServiceInstanceResponse, error)
	updateServiceInstance		func(id string, req *brokerapi.UpdateServiceInstanceRequest) (*brokerapi.UpdateServiceInstanceResponse, error)
	removeServiceInstance		func(id string) (*brokerapi.DeleteServiceInstanceResponse, error)
	bind				func(instanceID string, bindingID string, req *brokerapi.BindingRequest) (*brokerapi.CreateServiceBindingResponse, error)
	unBind				func(instanceID string, bindingID string) error
}

func (controller *Controller) Catalog() (*brokerapi.Catalog, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if controller.catalog == nil {
		controller.t.Error("Test failed to provide 'catalog' handler")
	}
	return controller.catalog()
}
func (controller *Controller) GetServiceInstanceLastOperation(instanceID, serviceID, planID, operation string) (*brokerapi.LastOperationResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if controller.getServiceInstanceLastOperation == nil {
		controller.t.Error("Test failed to provide 'getServiceInstanceLastOperation' handler")
	}
	return controller.getServiceInstanceLastOperation(instanceID)
}
func (controller *Controller) CreateServiceInstance(id string, req *brokerapi.CreateServiceInstanceRequest) (*brokerapi.CreateServiceInstanceResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if controller.createServiceInstance == nil {
		controller.t.Error("Test failed to provide 'createServiceInstance' handler")
	}
	return controller.createServiceInstance(id, req)
}
func (controller *Controller) UpdateServiceInstance(id string, req *brokerapi.UpdateServiceInstanceRequest) (*brokerapi.UpdateServiceInstanceResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if controller.updateServiceInstance == nil {
		controller.t.Error("Test failed to provide 'updateServiceInstance' handler")
	}
	return controller.updateServiceInstance(id, req)
}
func (controller *Controller) RemoveServiceInstance(instanceID, serviceID, planID string, acceptsIncomplete bool) (*brokerapi.DeleteServiceInstanceResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if controller.removeServiceInstance == nil {
		controller.t.Error("Test failed to provide 'removeServiceInstance' handler")
	}
	return controller.removeServiceInstance(instanceID)
}
func (controller *Controller) Bind(instanceID string, bindingID string, req *brokerapi.BindingRequest) (*brokerapi.CreateServiceBindingResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if controller.bind == nil {
		controller.t.Error("Test failed to provide 'bind' handler")
	}
	return controller.bind(instanceID, bindingID, req)
}
func (controller *Controller) UnBind(instanceID, bindingID, serviceID, planID string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if controller.unBind == nil {
		controller.t.Error("Test failed to provide 'unBind' handler")
	}
	return controller.unBind(instanceID, bindingID)
}

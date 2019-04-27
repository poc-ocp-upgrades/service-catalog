package controller_test

import (
	"github.com/kubernetes-incubator/service-catalog/pkg/controller"
	osb "github.com/pmorie/go-open-service-broker-client/v2"
	"testing"
)

func TestBrokerClientManager_CreateBrokerClient(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	osbCl1, _ := osb.NewClient(testOsbConfig("osb-1"))
	osbCl2, _ := osb.NewClient(testOsbConfig("osb-2"))
	brokerClientFunc := clientFunc(osbCl1, osbCl2)
	manager := controller.NewBrokerClientManager(brokerClientFunc)
	createdClient1, _ := manager.UpdateBrokerClient(controller.NewClusterServiceBrokerKey("broker1"), testOsbConfig("osb-1"))
	createdClient2, _ := manager.UpdateBrokerClient(controller.NewServiceBrokerKey("prod", "broker1"), testOsbConfig("osb-2"))
	gotClient1, exists1 := manager.BrokerClient(controller.NewClusterServiceBrokerKey("broker1"))
	gotClient2, exists2 := manager.BrokerClient(controller.NewServiceBrokerKey("prod", "broker1"))
	_, exists3 := manager.BrokerClient(controller.NewServiceBrokerKey("stage", "broker1"))
	if !exists1 {
		t.Fatal("Broker client osb-1 does not exist")
	}
	if !exists2 {
		t.Fatal("Broker client osb-2 does not exist")
	}
	if exists3 {
		t.Fatal("Broker client for namespace 'stage' must not exist")
	}
	if osbCl1 != createdClient1 {
		t.Fatalf("Wrong client from broker1")
	}
	if osbCl2 != createdClient2 {
		t.Fatalf("Wrong client from broker2")
	}
	if osbCl1 != gotClient1 {
		t.Fatalf("Wrong client from broker1")
	}
	if osbCl2 != gotClient2 {
		t.Fatalf("Wrong client from broker2")
	}
}
func TestBrokerClientManager_RemoveBrokerClient(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	osbCl1, _ := osb.NewClient(testOsbConfig("osb-1"))
	osbCl2, _ := osb.NewClient(testOsbConfig("osb-2"))
	brokerClientFunc := clientFunc(osbCl1, osbCl2)
	manager := controller.NewBrokerClientManager(brokerClientFunc)
	manager.UpdateBrokerClient(controller.NewClusterServiceBrokerKey("broker1"), testOsbConfig("osb-1"))
	manager.UpdateBrokerClient(controller.NewServiceBrokerKey("prod", "broker1"), testOsbConfig("osb-2"))
	manager.RemoveBrokerClient(controller.NewClusterServiceBrokerKey("broker1"))
	_, exists1 := manager.BrokerClient(controller.NewClusterServiceBrokerKey("broker1"))
	_, exists2 := manager.BrokerClient(controller.NewServiceBrokerKey("prod", "broker1"))
	if exists1 {
		t.Fatal("Broker client for 'broker1' must not exist")
	}
	if !exists2 {
		t.Fatal("Broker client osb-2 does not exist")
	}
}
func TestBrokerClientManager_UpdateBrokerClient(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	osbCl1, _ := osb.NewClient(testOsbConfig("osb-1"))
	osbCl2, _ := osb.NewClient(testOsbConfig("osb-2"))
	osbCl3, _ := osb.NewClient(testOsbConfig("osb-3"))
	brokerClientFunc := clientFunc(osbCl1, osbCl2, osbCl3)
	manager := controller.NewBrokerClientManager(brokerClientFunc)
	osbCfg := testOsbConfig("osb-1")
	osbCfg.AuthConfig = &osb.AuthConfig{BasicAuthConfig: &osb.BasicAuthConfig{Username: "user-1", Password: "password-1"}}
	osbCfgWithPasswordChange := testOsbConfig("osb-1")
	osbCfgWithPasswordChange.AuthConfig = &osb.AuthConfig{BasicAuthConfig: &osb.BasicAuthConfig{Username: "user-1", Password: "password-changed"}}
	manager.UpdateBrokerClient(controller.NewClusterServiceBrokerKey("broker1"), osbCfg)
	manager.UpdateBrokerClient(controller.NewServiceBrokerKey("prod", "broker1"), testOsbConfig("osb-2"))
	manager.UpdateBrokerClient(controller.NewClusterServiceBrokerKey("broker1"), osbCfgWithPasswordChange)
	gotClient, exists := manager.BrokerClient(controller.NewClusterServiceBrokerKey("broker1"))
	if !exists {
		t.Fatal("Broker client osb-2 does not exist")
	}
	if gotClient != osbCl3 {
		t.Fatalf("Broker client must have updated auth config")
	}
}
func clientFunc(clients ...osb.Client) osb.CreateFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i = 0
	return func(_ *osb.ClientConfiguration) (osb.Client, error) {
		client := clients[i]
		i++
		return client, nil
	}
}
func testOsbConfig(name string) *osb.ClientConfiguration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &osb.ClientConfiguration{Name: name}
}

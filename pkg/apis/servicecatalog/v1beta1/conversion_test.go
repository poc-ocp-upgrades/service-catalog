package v1beta1

import (
	"strings"
	"testing"
)

type conversionFunc func(string, string) (string, string, error)
type testcase struct {
	name		string
	inLabel		string
	inValue		string
	outLabel	string
	outValue	string
	success		bool
	expectedError	string
}

func TestClusterServicePlanFieldLabelConversionFunc(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []testcase{{name: "spec.externalName works", inLabel: "spec.externalName", inValue: "somenamehere", outLabel: "spec.externalName", outValue: "somenamehere", success: true}, {name: "spec.clusterServiceClassRef.name works", inLabel: "spec.clusterServiceClassRef.name", inValue: "someref", outLabel: "spec.clusterServiceClassRef.name", outValue: "someref", success: true}, {name: "spec.clusterServiceBrokerName works", inLabel: "spec.clusterServiceBrokerName", inValue: "somebroker", outLabel: "spec.clusterServiceBrokerName", outValue: "somebroker", success: true}, {name: "spec.externalID works", inLabel: "spec.externalID", inValue: "externalid", outLabel: "spec.externalID", outValue: "externalid", success: true}, {name: "random fails", inLabel: "spec.random", inValue: "randomvalue", outLabel: "", outValue: "", success: false, expectedError: "field label not supported: spec.random"}}
	runTestCases(t, cases, "ClusterServicePlanFieldLabelConversionFunc", ClusterServicePlanFieldLabelConversionFunc)
}
func TestServicePlanFieldLabelConversionFunc(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []testcase{{name: "spec.externalName works", inLabel: "spec.externalName", inValue: "somenamehere", outLabel: "spec.externalName", outValue: "somenamehere", success: true}, {name: "spec.serviceClassRef.name works", inLabel: "spec.serviceClassRef.name", inValue: "someref", outLabel: "spec.serviceClassRef.name", outValue: "someref", success: true}, {name: "spec.serviceBrokerName works", inLabel: "spec.serviceBrokerName", inValue: "somebroker", outLabel: "spec.serviceBrokerName", outValue: "somebroker", success: true}, {name: "spec.externalID works", inLabel: "spec.externalID", inValue: "externalid", outLabel: "spec.externalID", outValue: "externalid", success: true}, {name: "random fails", inLabel: "spec.random", inValue: "randomvalue", outLabel: "", outValue: "", success: false, expectedError: "field label not supported: spec.random"}}
	runTestCases(t, cases, "ServicePlanFieldLabelConversionFunc", ServicePlanFieldLabelConversionFunc)
}
func TestClusterServiceClassFieldLabelConversionFunc(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []testcase{{name: "spec.externalName works", inLabel: "spec.externalName", inValue: "somenamehere", outLabel: "spec.externalName", outValue: "somenamehere", success: true}, {name: "spec.clusterServiceClassRef.name fails", inLabel: "spec.clusterServiceClassRef.name", inValue: "someref", outLabel: "", outValue: "", success: false, expectedError: "field label not supported: spec.clusterServiceClassRef.name"}, {name: "spec.clusterServiceBrokerName works", inLabel: "spec.clusterServiceBrokerName", inValue: "somebroker", outLabel: "spec.clusterServiceBrokerName", outValue: "somebroker", success: true}, {name: "spec.externalID works", inLabel: "spec.externalID", inValue: "externalid", outLabel: "spec.externalID", outValue: "externalid", success: true}, {name: "random fails", inLabel: "spec.random", inValue: "randomvalue", outLabel: "", outValue: "", success: false, expectedError: "field label not supported: spec.random"}}
	runTestCases(t, cases, "ClusterServiceClassFieldLabelConversionFunc", ClusterServiceClassFieldLabelConversionFunc)
}
func TestServiceInstanceFieldLabelConversionFunc(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []testcase{{name: "spec.clusterServiceClassRef.name works", inLabel: "spec.clusterServiceClassRef.name", inValue: "someref", outLabel: "spec.clusterServiceClassRef.name", outValue: "someref", success: true}, {name: "spec.clusterServicePlanRef.name works", inLabel: "spec.clusterServicePlanRef.name", inValue: "someref", outLabel: "spec.clusterServicePlanRef.name", outValue: "someref", success: true}, {name: "spec.serviceClassRef.name works", inLabel: "spec.serviceClassRef.name", inValue: "someref", outLabel: "spec.serviceClassRef.name", outValue: "someref", success: true}, {name: "spec.servicePlanRef.name works", inLabel: "spec.servicePlanRef.name", inValue: "someref", outLabel: "spec.servicePlanRef.name", outValue: "someref", success: true}, {name: "spec.externalID works", inLabel: "spec.externalID", inValue: "externalid", outLabel: "spec.externalID", outValue: "externalid", success: true}, {name: "random fails", inLabel: "spec.random", inValue: "randomvalue", outLabel: "", outValue: "", success: false, expectedError: "field label not supported: spec.random"}}
	runTestCases(t, cases, "ServiceInstanceFieldLabelConversionFunc", ServiceInstanceFieldLabelConversionFunc)
}
func TestServiceBindingFieldLabelConversionFunc(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []testcase{{name: "spec.externalID works", inLabel: "spec.externalID", inValue: "externalid", outLabel: "spec.externalID", outValue: "externalid", success: true}, {name: "random fails", inLabel: "spec.random", inValue: "randomvalue", outLabel: "", outValue: "", success: false, expectedError: "field label not supported: spec.random"}}
	runTestCases(t, cases, "ServiceBindingFieldLabelConversionFunc", ServiceBindingFieldLabelConversionFunc)
}
func runTestCases(t *testing.T, cases []testcase, testFuncName string, testFunc conversionFunc) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, tc := range cases {
		outLabel, outValue, err := testFunc(tc.inLabel, tc.inValue)
		if tc.success {
			if err != nil {
				t.Errorf("%s:%s -- unexpected failure : %q", testFuncName, tc.name, err.Error())
			} else {
				if a, e := outLabel, tc.outLabel; a != e {
					t.Errorf("%s:%s -- label mismatch, expected %q got %q", testFuncName, tc.name, e, a)
				}
				if a, e := outValue, tc.outValue; a != e {
					t.Errorf("%s:%s -- value mismatch, expected %q got %q", testFuncName, tc.name, e, a)
				}
			}
		} else {
			if err == nil {
				t.Errorf("%s:%s -- unexpected success, expected: %q", testFuncName, tc.name, tc.expectedError)
			} else {
				if !strings.Contains(err.Error(), tc.expectedError) {
					t.Errorf("%s:%s -- did not find expected error %q got %q", testFuncName, tc.name, tc.expectedError, err)
				}
			}
		}
	}
}

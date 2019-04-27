package v1beta1

import (
	"fmt"
	"testing"
)

func TestPlanReference_Format(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testcases := []struct {
		name	string
		format	string
		want	string
		planRef	PlanReference
	}{{"all: external-name", "%v", `{ClusterServiceClassExternalName:"foo", ClusterServicePlanExternalName:"bar"}`, PlanReference{ClusterServiceClassExternalName: "foo", ClusterServicePlanExternalName: "bar"}}, {"all: external-id", "%v", `{ClusterServiceClassExternalID:"foo-abc123", ClusterServicePlanExternalID:"bar-def456"}`, PlanReference{ClusterServiceClassExternalID: "foo-abc123", ClusterServicePlanExternalID: "bar-def456"}}, {"all: cluster-name", "%v", `{ClusterServiceClassName:"k8s-foo1232", ClusterServicePlanName:"k8s-bar456"}`, PlanReference{ClusterServiceClassName: "k8s-foo1232", ClusterServicePlanName: "k8s-bar456"}}, {"short: external-name", "%s", `foo/bar`, PlanReference{ClusterServiceClassExternalName: "foo", ClusterServicePlanExternalName: "bar"}}, {"short: external-id", "%s", `foo-abc123/bar-def456`, PlanReference{ClusterServiceClassExternalID: "foo-abc123", ClusterServicePlanExternalID: "bar-def456"}}, {"short: cluster-name", "%s", `k8s-foo1232/k8s-bar456`, PlanReference{ClusterServiceClassName: "k8s-foo1232", ClusterServicePlanName: "k8s-bar456"}}, {"class: external-name", "%c", `{ClusterServiceClassExternalName:"foo"}`, PlanReference{ClusterServiceClassExternalName: "foo", ClusterServicePlanExternalName: "bar"}}, {"class: external-id", "%c", `{ClusterServiceClassExternalID:"foo-abc123"}`, PlanReference{ClusterServiceClassExternalID: "foo-abc123", ClusterServicePlanExternalID: "bar-def456"}}, {"class: cluster-name", "%c", `{ClusterServiceClassName:"k8s-foo1232"}`, PlanReference{ClusterServiceClassName: "k8s-foo1232", ClusterServicePlanName: "k8s-bar456"}}, {"plan: external-name", "%b", `{ClusterServicePlanExternalName:"bar"}`, PlanReference{ClusterServiceClassExternalName: "foo", ClusterServicePlanExternalName: "bar"}}, {"plan: external-id", "%b", `{ClusterServicePlanExternalID:"bar-def456"}`, PlanReference{ClusterServiceClassExternalID: "foo-abc123", ClusterServicePlanExternalID: "bar-def456"}}, {"plan: cluster-name", "%b", `{ClusterServicePlanName:"k8s-bar456"}`, PlanReference{ClusterServiceClassName: "k8s-foo1232", ClusterServicePlanName: "k8s-bar456"}}}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := fmt.Sprintf(tc.format, tc.planRef)
			if tc.want != got {
				t.Fatalf("\nwant:\t%#v\ngot:\t%#v", tc.want, got)
			}
		})
	}
}

#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

gather_artificats() {
    set +e
    oc describe pods -n kube-service-catalog >  /tmp/artifacts/describe-catalog-pods.txt
    oc get events -n kube-service-catalog >  /tmp/artifacts/catalog-events.txt
    oc get all -n kube-service-catalog >  /tmp/artifacts/all-objects-in-catalog-ns.txt
    oc get operatorgroups --all-namespaces >  /tmp/artifacts/all-operator-groups.txt
    oc get csv  svcat.v0.1.34 -n kube-service-catalog -o yaml > /tmp/artifacts/svc-cat-csv.yaml
    oc describe csv  svcat.v0.1.34 -n kube-service-catalog > /tmp/artifacts/describe-svc-cat-csv.txt
    oc get clusterrole > /tmp/artifacts/cluster-roles.txt
    oc get clusterrole system:service-catalog:aggregate-to-admin -o yaml > /tmp/artifacts/svc-cat-aggregated-cluster-roles.yaml
    oc get clusterrole system:service-catalog:aggregate-to-edit >> /tmp/artifacts/svc-cat-aggregated-cluster-roles.yaml
    oc get clusterrole system:service-catalog:aggregate-to-view >> /tmp/artifacts/svc-cat-aggregated-cluster-roles.yaml
}


delete_resources() {
    oc delete subscription svcat -n kube-service-catalog
    oc delete clusterserviceversion svcat.v0.1.34 -n kube-service-catalog
    oc delete namespace kube-service-catalog
}

for sig in INT TERM EXIT; do
    trap "set +e;gather_artificats; [[ $sig == EXIT ]] || kill -$sig $BASHPID" $sig
done


echo "`date +%T` Waiting for up to 10 minutes for Service Catalog APIs to be available..."

TARGET="$(date -d '5 minutes' +%s)"
NOW="$(date +%s)"
while [[ "${NOW}" -lt "${TARGET}" ]]; do
  REMAINING="$((TARGET - NOW))"
  if oc --request-timeout="${REMAINING}s" get --raw /apis/servicecatalog.k8s.io/v1beta1 ; then
    break
  fi
  sleep 20
  NOW="$(date +%s)"
done

if [ "${NOW}" -ge "${TARGET}" ];then
    echo "`date +%T`: timeout waiting for service-catalog apis to be available"
    # could fail out here with an exit 1, leave it to fail e2e for now.
fi

echo "Add missing rbac"
set +e
cat <<'EOF' | oc create -f -
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: add-servicebindingfinalizers
rules:
- apiGroups:
  - servicecatalog.k8s.io
  resources:
  - servicebindings/finalizers
  verbs:
  - update
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: add-servicebindingfinalizers
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: add-servicebindingfinalizers
subjects:
- kind: ServiceAccount
  name: service-catalog-controller
  namespace: kube-service-catalog
EOF
sleep 5
set -e

oc get pods -l app=controller-manager -n kube-service-catalog  -o name |  xargs -I{} oc logs {} -n kube-service-catalog  | grep -o "Service Catalog version.*" > /tmp/artifacts/service-catalog-version.txt


echo "`date +%T`: Service Catalog APIs available, executing Service Catalog E2E"

SERVICECATALOGCONFIG=$KUBECONFIG bin/e2e.test -v 10 -alsologtostderr -broker-image quay.io/kubernetes-service-catalog/user-broker:latest


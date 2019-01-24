#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

SVCAT_NAMESPACE=kube-service-catalog

gather_artificats() {
    set +e
    oc describe pods -n $SVCAT_NAMESPACE >  /tmp/artifacts/svcat-describe-catalog-pods.txt
    oc get events -n $SVCAT_NAMESPACE >  /tmp/artifacts/svcat-catalog-events.txt
    oc get all -n $SVCAT_NAMESPACE >  /tmp/artifacts/svcat-all-objects-in-catalog-ns.txt
    oc get operatorgroups --all-namespaces >  /tmp/artifacts/svcat-operator-groups.txt
    oc get subscription --all-namespaces > /tmp/artifacts/svcat-subscriptions.txt
    oc get csv --all-namespaces > /tmp/artifacts/svcat-csvs.txt
    oc get catalogsourceconfigs --all-namespaces > /tmp/artifacts/svcat-catalogsourceconfigs.txt
    oc get catalogsources --all-namespaces > /tmp/artifacts/svcat-catalogsources.txt
    oc describe csv  svcat.v0.1.34 -n $SVCAT_NAMESPACE > /tmp/artifacts/svcat-describe-svc-cat-csv.txt
    oc get clusterrole > /tmp/artifacts/svcat-cluster-roles.txt
    oc get clusterrole system:service-catalog:aggregate-to-admin -o yaml > /tmp/artifacts/svcat-aggregated-cluster-roles.yaml
    oc get clusterrole system:service-catalog:aggregate-to-edit -o yaml >> /tmp/artifacts/svcat-aggregated-cluster-roles.yaml
    oc get clusterrole system:service-catalog:aggregate-to-view -o yaml >> /tmp/artifacts/svcat-aggregated-cluster-roles.yaml
}


delete_resources() {
    oc delete subscription svcat -n $SVCAT_NAMESPACE
    oc delete clusterserviceversion svcat.v0.1.34 -n $SVCAT_NAMESPACE
}

for sig in INT TERM EXIT; do
    trap "set +e;gather_artificats; [[ $sig == EXIT ]] || kill -$sig $BASHPID" $sig
done


echo "`date +%T` Waiting for up to 10 minutes for Service Catalog APIs to be available..."

TARGET="$(date -d '10 minutes' +%s)"
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


oc get pods -l app=controller-manager -n $SVCAT_NAMESPACE  -o name |  xargs -I{} oc logs {} -n $SVCAT_NAMESPACE  | grep -o "Service Catalog version.*" > /tmp/artifacts/service-catalog-version.txt


echo "`date +%T`: Service Catalog APIs available, executing Service Catalog E2E"

SERVICECATALOGCONFIG=$KUBECONFIG bin/e2e.test -v 10 -alsologtostderr -broker-image quay.io/kubernetes-service-catalog/user-broker:latest


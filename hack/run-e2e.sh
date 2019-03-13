#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

gather_artificats() {
    set +e
    oc describe deployment openshift-service-catalog-apiserver-operator -n openshift-service-catalog-apiserver-operator >  /tmp/artifacts/svcat-describe-deployment-openshift-svcat-apiserver-operator.txt
    oc describe deployment openshift-service-catalog-controller-manager-operator -n openshift-service-catalog-controller-manager-operator >  /tmp/artifacts/svcat-describe-deployment-openshift-svcat-controller-manager-operator.txt
    oc describe pods -n openshift-service-catalog-apiserver >  /tmp/artifacts/svcat-describe-openshift-service-catalog-apiserver-pods.txt
    oc describe pods -n openshift-service-catalog-controller-manager >  /tmp/artifacts/svcat-describe-openshift-service-catalog-controller-manager-pods.txt
    oc get events --sort-by='.lastTimestamp' -n openshift-service-catalog-apiserver >  /tmp/artifacts/svcat-catalog-events.txt
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



echo "`date +%T` Waiting for up to 5 minutes for Service Catalog Controller Manager to be available..."

TARGET="$(date -d '5 minutes' +%s)"
NOW="$(date +%s)"
while [[ "${NOW}" -lt "${TARGET}" ]]; do
  REMAINING="$((TARGET - NOW))"
  # todo - use jq vs awk
  if oc get clusteroperators service-catalog-controller-manager  |  awk '/service-catalog-controller-manager/ {print $3}' | grep True ; then
    break
  fi
  sleep 10
  NOW="$(date +%s)"
done

if [ "${NOW}" -ge "${TARGET}" ];then
    echo "`date +%T`: timeout waiting for clusteroperators service-catalog-controller-manager to be available"
    # could fail out here with an exit 1, leave it to fail e2e for now.
fi

oc get pods -l app=apiserver -n openshift-service-catalog-apiserver  -o name |  xargs -I{} oc logs {} -n openshift-service-catalog-apiserver  | grep -o "Service Catalog version.*" > /tmp/artifacts/service-catalog-version.txt
cat /tmp/artifacts/service-catalog-version.txt


echo "`date +%T`: Service Catalog APIs available, executing Service Catalog E2E"

SERVICECATALOGCONFIG=$KUBECONFIG bin/e2e.test -v 10 -alsologtostderr -broker-image quay.io/kubernetes-service-catalog/user-broker:latest


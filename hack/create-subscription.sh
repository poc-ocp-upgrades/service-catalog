#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

echo Preparing Service Catalog E2E at $(TZ=EST date) / $(TZ=UTC date +"%k:%M:%S %Z")


oc create namespace kube-service-catalog || true

oc apply -f hack/svcat-catalogsourceconfig.yaml

# wait for the catalog source pod
#oc get pods -lolm.catalogSource=service-catalog -n kube-service-catalog


echo "`date +%T` Waiting for for catalogsource to be ready..."
TARGET="$(date -d '5 minutes' +%s)"
NOW="$(date +%s)"
while [[ "${NOW}" -lt "${TARGET}" ]]; do
  REMAINING="$((TARGET - NOW))"
  oc get pods -lolm.catalogSource=service-catalog -n kube-service-catalog || true
  JSONPATH="{range .items[*]}{@.metadata.name}:{range @.status.conditions[*]}{@.type}={@.status};{end}{end}"
  STATUS=`oc get pods -lolm.catalogSource=service-catalog -n kube-service-catalog -o jsonpath="$JSONPATH"`
  if [[ $STATUS == *"Ready=True"* ]];then
    break
  fi
  sleep 20
  NOW="$(date +%s)"
done

if [ "${NOW}" -ge "${TARGET}" ];then
    echo "`date +%T`: warning: catalogsource not ready"
else
    echo "`date +%T`: catalogsource is running & ready, proceeding"
fi

oc apply -f hack/svcat-operatorgroup.yaml
oc apply -f hack/svcat-subscription.yaml

#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

echo Preparing Service Catalog E2E at $(TZ=EST date) / $(TZ=UTC date +"%k:%M:%S %Z")
oc create namespace kube-service-catalog || true
oc label ns kube-service-catalog ns=kube-service-catalog || true

cat <<'EOF' | oc create -f -
apiVersion: operators.coreos.com/v1alpha1
kind: Subscription
metadata:
  name: svcat
  namespace: kube-service-catalog
spec:
  channel: alpha
  name: svcat
  source: rh-operators
  installPlanApproval: Automatic
  catalogSourceNamespace: kube-service-catalog
---
apiVersion: operators.coreos.com/v1alpha2
kind: OperatorGroup
metadata:
  name: service-catalog
  namespace: kube-service-catalog
spec:
  selector:
    matchLabels:
      ns: kube-service-catalog
EOF

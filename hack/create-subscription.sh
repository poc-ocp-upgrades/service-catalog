#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

#
# todo:  rename this file and in CI
#
# This script toggles the managementState of the ServiceCatalogAPIServers and
# ServiceCatalogControllerManagers custom resources to "Managed" causing the
# service catalog cluster operators to install Service Catalog.
#
#
echo Preparing Service Catalog E2E at $(TZ=EST date) / $(TZ=UTC date +"%k:%M:%S %Z")

cat <<'EOF' | oc apply -f -
apiVersion: operator.openshift.io/v1
kind: ServiceCatalogAPIServer
metadata:
  name: cluster
spec:
  logLevel: "Normal"
  managementState: Managed
EOF

cat <<'EOF' | oc apply -f -
apiVersion: operator.openshift.io/v1
kind: ServiceCatalogControllerManager
metadata:
  name: cluster
spec:
  logLevel: "Normal"
  managementState: Managed
EOF


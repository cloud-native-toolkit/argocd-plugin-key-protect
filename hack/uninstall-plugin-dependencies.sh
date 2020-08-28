#!/usr/bin/env bash

SCRIPT_DIR=$(cd $(dirname "$0"); pwd -P)

NAMESPACE="$1"
if [[ -z "${NAMESPACE}" ]]; then
  echo "Namespace is required as the first argument"
  exit 1
fi

RESOURCE="deployments/argocd-repo-server"

kubectl patch "${RESOURCE}" -n "${NAMESPACE}" --type json -p "$(cat "${SCRIPT_DIR}/uninstall-plugin-dependencies-patch.json")"

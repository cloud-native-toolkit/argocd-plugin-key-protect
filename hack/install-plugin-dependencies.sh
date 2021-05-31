#!/usr/bin/env bash

SCRIPT_DIR=$(cd $(dirname "$0"); pwd -P)

NAMESPACE="$1"
if [[ -z "${NAMESPACE}" ]]; then
  echo "Namespace is required as the first argument"
  exit 1
fi

NAME="$2"
if [[ -z "${NAME}" ]]; then
  NAME="argocd"
fi

DEPLOYMENT="deployments/${NAME}-repo-server"

echo "Patching ${DEPLOYMENT} in ${NAMESPACE}"
kubectl patch "${DEPLOYMENT}" -n "${NAMESPACE}" --type json -p "$(cat "${SCRIPT_DIR}/install-plugin-dependencies-patch.json")"

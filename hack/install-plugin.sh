#!/usr/bin/env bash

SCRIPT_DIR=$(cd $(dirname "$0"); pwd -P)

NAMESPACE="$1"
if [[ -z "${NAMESPACE}" ]]; then
  echo "Namespace is required as the first argument"
  exit 1
fi

ARGOCD="$2"
if [[ -z "${ARGOCD}" ]]; then
  ARGOCD="argocd"
fi

PATCH_VALUE="$(cat "${SCRIPT_DIR}/install-plugin-patch.yaml")"

if [[ -n $(kubectl get argocd "${ARGOCD}" -n "${NAMESPACE}" -o jsonpath='{.spec.configManagementPlugins}') ]]; then
  VALUE="$(kubectl get argocd "${ARGOCD}" -n "${NAMESPACE}" -o jsonpath='{.spec.configManagementPlugins}')"
  if echo "${VALUE}" | grep -q "key-protect-secret"; then
    echo "key-protect-secret plugin already installed"
    exit 0
  fi

  VALUE="$(echo "$PATCH_VALUE"; echo "$VALUE" | sed -E "s/^(.*)/    \1/g")"

  echo "$VALUE"
  kubectl patch argocd "${ARGOCD}" -n "${NAMESPACE}" --type merge --patch "${VALUE}"
else
  kubectl patch argocd "${ARGOCD}" -n "${NAMESPACE}" --type merge --patch "${PATCH_VALUE}"
fi

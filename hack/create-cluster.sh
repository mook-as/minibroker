#!/usr/bin/env bash

set -euo pipefail

: "${VM_DRIVER:=virtualbox}"
: "${VM_MEMORY:=$(( 1024 * 4 ))}"

if [[ "$(minikube status)" != *"Running"* ]]; then
    set -x
    minikube start \
      --vm-driver="${VM_DRIVER}" \
      --memory="${VM_MEMORY}" \
      --kubernetes-version=v1.11.3 \
      --bootstrapper=kubeadm
else
    echo "Using current running instance of Minikube..."
    set -x
fi

minikube addons enable heapster

kubectl apply -f https://raw.githubusercontent.com/Azure/helm-charts/master/docs/prerequisities/helm-rbac-config.yaml
helm init --service-account tiller --wait

helm repo add svc-cat https://svc-catalog-charts.storage.googleapis.com
helm upgrade --install catalog --namespace svc-cat svc-cat/catalog --wait

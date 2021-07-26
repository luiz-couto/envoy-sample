#!/bin/bash
set -e

DIR=$1
NAMESPACE=envoy-sample

if [ -z "$DIR" ]; then
    echo "Usage: $(basename $0) dir" >&2
    exit 1
fi

if [ -f "${DIR}/build.sh" ]; then
    ( cd "${DIR}" && sh build.sh )
fi

hasNamespace=$(kubectl get namespaces |  grep ${NAMESPACE})

if [ ${#hasNamespace} == 0 ]; then
    echo "creating ${NAMESPACE} namespace..."
    kubectl create namespace ${NAMESPACE}
fi

kubectl apply -f "${DIR}/deployment.yaml" -n ${NAMESPACE}
kubectl apply -f "${DIR}/service.yaml" -n ${NAMESPACE}
#!/bin/bash
set -e

DIR=$1
NAMESPACE=envoy-sample

if [ -z "$DIR" ]; then
    echo "Usage: $(basename $0) dir" >&2
    exit 1
fi

kubectl delete -f "${DIR}/deployment.yaml" -n ${NAMESPACE}
kubectl delete -f "${DIR}/service.yaml" -n ${NAMESPACE}
#!/usr/bin/env bash
if [[ -z ${IMAGE_TAG} ]]; then
  echo "IMAGE_TAG is empty"
  exit 1
fi
envsubst < config.$ENVIRONMENT.yml > k8s-config.yml && kubectl apply -f k8s-config.yml -n $NAMESPACE
envsubst < migrate.yml > k8s-migrate.yml
kubectl delete job ${SERVICE_NAME} --ignore-not-found=true --namespace=${NAMESPACE}
kubectl apply -f k8s-migrate.yml --namespace=${NAMESPACE}
kubectl wait --for=condition=complete --timeout=180s job.batch/${SERVICE_NAME} --namespace=${NAMESPACE}
if [[ $? != 0 ]]; then
    echo "Migration job failed!"
    kubectl logs job/${SERVICE_NAME} --namespace=${NAMESPACE}
    exit 1
fi
kubectl logs job/${SERVICE_NAME} --namespace=${NAMESPACE}

#!/usr/bin/env bash
if [[ -z $IMAGE_TAG ]]; then
  echo "IMAGE_TAG is empty"
  exit 1
fi

if [[ "$CI_COMMIT_TAG" != "" ]]; then
      export IMAGE_TAG=$CI_COMMIT_TAG
fi

envsubst < config.$ENVIRONMENT.yml > k8s-config.yml && kubectl apply -f k8s-config.yml -n $NAMESPACE

envsubst < main.$ENVIRONMENT.yml > k8s-main.yml && kubectl apply -f k8s-main.yml -n $NAMESPACE

if [[ $? != 0 ]]; then exit 1; fi

kubectl rollout status deployments/$SERVICE_NAME -n $NAMESPACE

if [[ $? != 0 ]]; then
    kubectl logs $(kubectl get pods -n $NAMESPACE --sort-by=.metadata.creationTimestamp | grep "$SERVICE_NAME" | awk '{print $1}' | tac | head -1 ) -n $NAMESPACE --tail=20 && exit 1;
fi

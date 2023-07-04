#!/bin/bash

# shellcheck disable=SC2046
eval $(minikube -p minikube docker-env)

sed -e s/service_name/$1/g Tdockerfile > Dockerfile


docker build -t $1 .

kubectl delete -f $1.yaml
kubectl apply -f $1.svc.yaml
kubectl apply -f $1.yaml
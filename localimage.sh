#!/bin/bash

eval $(minikube -p minikube docker-env)

kubectl run $1 --image=$1 --image-pull-policy=Never  --image-pull-policy=Never
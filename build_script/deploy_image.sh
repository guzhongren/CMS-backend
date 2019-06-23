#!/bin/bash

NAMESPACE=$1
DEPLOYMENT=$2
CONTAINER=$3

IMAGE=`cat image_name`
echo "镜像名称：${IMAGE}"
echo "要更新的容器:${CONTAINER}"
echo "要更新的 Deployment:${DEPLOYMENT}"
echo "namespace: ${NAMESPACE}"
echo "开始更新镜像"
kubectl --kubeconfig /root/.kube/config -n ${NAMESPACE} set image deployment/${DEPLOYMENT} "${CONTAINER}"=${IMAGE}
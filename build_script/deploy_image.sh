#!/bin/bash

DEPLOYMENT=$1
# APP=$2
IMAGE=`cat image_name`
echo "镜像名称：${IMAGE}"
# echo "要更新的容器:${APP}"
echo "要更新的 Deployment:${DEPLOYMENT}"
echo "测试kubectl是否可用"
kubectl --kubeconfig .kube/config  get nodes
echo "测试结束"
kubectl --kubeconfig .kube/config set image delplyments/${DEPLOYMENT} "${DEPLOYMENT}-container"=${IMAGE}
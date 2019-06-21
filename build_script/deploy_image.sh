#!/bin/bash

DEPLOYMENT=$1
APP=$2
IMAGE=`cat image_name`
echo "镜像名称：${IMAGE}"
echo "要更新的容器:${APP}"
echo "要更新的 Deployment:${DEPLOYMENT}"
# kubectl set image delplyments/${DEPLOYMENT} ${APP}=${IMAGE}
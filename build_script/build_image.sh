#!/bin/bash

HUB_DOMAIN=$1
DOCKERUSER=$2
DOCKERPASSWORD=$3
PROJECT=$4
APP=$5

echo "私有仓库域: ${HUB_DOMAIN}"
echo "docker 用户名: ${DOCKERUSER}"
echo "docker 密码: ******"
echo "项目名称: ${PROJECT}"
echo "应用名称: ${APP}"
TIME=`date "+%Y%m%d%H%M"`
GIT_REVISION=`git log -1 --pretty=format:"%H"`
# IMAGE_NAME=hub.k8s.com:8080/cms/${MODULE}:${TIME}_${GIT_REVISION}
IMAGE_NAME=${HUB_DOMAIN}/${PROJECT}/${APP}:${TIME}_${GIT_REVISION}

# cd ${MODULE}

docker login -u ${DOCKERUSER} -p ${DOCKERPASSWORD} ${HUB_DOMAIN}

docker build -t ${IMAGE_NAME} .

# cd -
docker push ${IMAGE_NAME}
echo ${IMAGE_NAME} > image_name
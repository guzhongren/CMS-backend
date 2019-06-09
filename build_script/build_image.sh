#!/bin/bash

DOCKERUSER=$1
DOCKERPASSWORD=$2
MODULE=$3
TIME=`date "+%Y%m%d%H%M"`
GIT_REVISION=`git log -1 --pretty=format:"%H"`
# IMAGE_NAME=hub.k8s.com:8080/cms/${MODULE}:${TIME}_${GIT_REVISION}
IMAGE_NAME=hub.k8s.com/cms/${MODULE}:${TIME}_${GIT_REVISION}

# cd ${MODULE}

docker login -u ${DOCKERUSER} -p ${DOCKERPASSWORD} hub.k8s.com

docker build -t ${IMAGE_NAME} .

# cd -
docker push ${IMAGE_NAME}
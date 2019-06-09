#!/bin/bash

MODULE=$1
TIME=`date "+%Y%m%d%H%M"`
GIT_REVISION=`git log -1 --pretty=format:"%H"`
# IMAGE_NAME=hub.k8s.com:8080/cms/${MODULE}:${TIME}_${GIT_REVISION}
IMAGE_NAME=hub.k8s.com/cms/${MODULE}:${TIME}_${GIT_REVISION}

# cd ${MODULE}

docker build -t ${IMAGE_NAME} .

# cd -
docker push ${IMAGE_NAME}
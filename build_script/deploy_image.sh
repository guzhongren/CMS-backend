#!/bin/bash

DEPLOYMENT=$1
IMAGE_NAME=`cat image_name`
echo ${IMAGE_NAME}
echo ${DEPLOYMENT}
# Kubectl set image delplyments/${DEPLOYMENT} user-edge-service = ${IMAGE_NAME}
#!/bin/bash

source .env

TAG=$(git rev-parse --short HEAD)
IMAGE_NAME=benchmark:$TAG
GCR_IMAGE=gcr.io/$GOOGLE_CLOUD_PROJECT/$IMAGE_NAME

docker build -t $GCR_IMAGE .
docker push $GCR_IMAGE
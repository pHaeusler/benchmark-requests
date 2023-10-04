#!/bin/bash

source .env

TAG=$(git rev-parse --short HEAD)
IMAGE_NAME=benchmark:$TAG
GCR_IMAGE=gcr.io/$GOOGLE_CLOUD_PROJECT/$IMAGE_NAME
echo $GCR_IMAGE
exit

gcloud compute instances create-with-container benchmark \
  --machine-type=n1-standard-1 \
  --zone=us-west1-a \
  --tags=http-server \
  --container-image $GCR_IMAGE

gcloud compute firewall-rules create allow-http --allow tcp:80 --target-tags http-server

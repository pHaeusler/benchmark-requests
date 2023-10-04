#!/bin/bash

source .env

TAG=$(git rev-parse --short HEAD)
IMAGE_NAME=benchmark:$TAG
GCR_IMAGE=gcr.io/$GOOGLE_CLOUD_PROJECT/$IMAGE_NAME
echo $GCR_IMAGE
exit

gcloud compute instances delete benchmark --zone=us-west1-a

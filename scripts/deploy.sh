#!/usr/bin/env bash

set -ex

parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
cd "${parent_path}"
# Navigate to root folder
cd ..

PROJECT="licenta-diana"
# Build docker image and push to GCR
docker build --tag=gcr.io/${PROJECT}/file-service .
docker push gcr.io/${PROJECT}/file-service

# Optional: Remove other old images if automatically if not needed
# prev_ver=`gcloud container images list-tags gcr.io/${PROJECT}/watermark | grep -v latest | awk '{if(NR>1)print $1}'`
# if [ ! -z "$prev_ver" ]; then
#     yes | gcloud container images delete gcr.io/${PROJECT}/watermark@sha256:${prev_ver} --force-delete-tags
# fi

# If cluster is not already created use this to create a new cluster (this shouldn't be here)
# gcloud beta container --project "licenta-diana" clusters create "cluster" --zone "us-east1-b" --username "admin" --cluster-version "1.11.7-gke.4" --machine-type "g1-small" --image-type "COS" --disk-type "pd-standard" --disk-size "30" --scopes "https://www.googleapis.com/auth/devstorage.read_only","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" --num-nodes "1" --enable-cloud-logging --enable-cloud-monitoring --no-enable-ip-alias --network "projects/licenta-diana/global/networks/default" --subnetwork "projects/licenta-diana/regions/us-east1/subnetworks/default" --addons HorizontalPodAutoscaling,HttpLoadBalancing --enable-autoupgrade --enable-autorepair

# Also, don't forget to expose the app to internet
# kubectl expose deployment watermark --type=LoadBalancer --port 80 --target-port 8000

# Deploy a new version
kubectl apply -f pod.yaml
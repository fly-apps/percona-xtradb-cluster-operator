#!/bin/sh

IMAGE=flyio/percona-xtradb-cluster-operator:pxc-5
docker build --platform linux/amd64 -t $IMAGE .
docker push $IMAGE

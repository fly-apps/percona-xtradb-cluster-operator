#!/bin/sh

IMAGE=flyio/percona-xtradb-cluster-operator:pxc-4
docker build --platform linux/amd64 -t $IMAGE .
docker push $IMAGE

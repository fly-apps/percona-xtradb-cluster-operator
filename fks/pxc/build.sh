#!/bin/sh

IMAGE=flyio/percona-server-mysql-operator:pxc-2
docker build --platform linux/amd64 -t $IMAGE .
docker push $IMAGE

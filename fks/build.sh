#!/bin/sh

IMAGE=flyio/percona-xtradb-cluster-operator:fks-5

docker tag perconalab/percona-xtradb-cluster-operator:js-fks $IMAGE
docker push $IMAGE

#!/bin/bash

set -u

echo $PROJECTNAME

telepresence \
--swap-deployment $PROJECTNAME \
--method container \
--docker-run --rm -it \
-v $GOBASE:/app \
$IMAGE:dev

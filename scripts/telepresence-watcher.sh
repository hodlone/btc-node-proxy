#!/bin/bash

set -u

telepresence \
--swap-deployment $PROJECTNAME \
--method container \
--expose $PORT \
--docker-run --rm -it \
-v $GOBASE:/app \
$IMAGE:dev
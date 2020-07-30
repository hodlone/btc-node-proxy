#!/bin/bash

set -u

docker run -it --rm \
-v $GOBASE:/app \
--net=host \
--env GO111MODULE=auto \
--env PORT=$PORT \
--env BTC_NODE_ZMQ_ADDR=$BTC_NODE_ZMQ_ADDR \
--env NATS_ADDR=$NATS_ADDR \
--env NATS_NAME=$NATS_NAME \
$IMAGE:$TARGET

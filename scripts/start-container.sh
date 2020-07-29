#!/bin/bash

IMAGE=$1
TARGET=$2
PROJECT_PATH=$(pwd)

echo $PROJECT_PATH
docker build --target=$TARGET -t $IMAGE:$TARGET .

docker run -it --rm \
-v $PROJECT_PATH:/app \
--net=host \
--env GO111MODULE=auto \
--env PORT=4000 \
--env BTC_NODE_ZMQ_ADDR=tcp://0.0.0.0:29000 \
--env NATS_ADDR=nats://localhost:4222 \
--env NATS_NAME=test-test-test \
$IMAGE:$TARGET

# docker run \
# --net=host \
# --env PORT=4000 \
# --env BTC_NODE_ZMQ_ADDR=tcp://0.0.0.0:29000 \
# --env NATS_ADDR=nats://localhost:4222 \
# --env NATS_NAME=test-test-test \
# ricardo/btc-node-proxy-hot-reload:latest
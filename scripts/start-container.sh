#!/bin/bash

NAME=btc-node-proxy

docker build -t $NAME .

docker run -it --rm \
--net=host \
--env PORT=4000 \
--env BTC_NODE_ZMQ_ADDR=tcp://0.0.0.0:29000 \
--env NATS_ADDR=nats://localhost:4222 \
--env NATS_NAME=test-test-test \
$NAME:latest
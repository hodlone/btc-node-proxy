#!/bin/bash

PORT=4000 \
BTC_NODE_ZMQ_ADDR=tcp://0.0.0.0:29000 \
NATS_ADDR=nats://localhost:4222 \
NATS_NAME=test-test-test \
go run \
./main.go
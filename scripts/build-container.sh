#!/bin/bash

set -u

if [[ $TARGET == "dev" || $TARGET == "prod" ]]; then
    echo "Name $IMAGE:$TARGET"
    docker build --target=$TARGET -t $IMAGE:$TARGET .
else
    echo "Wrong build target, posible targets dev|prod"
fi

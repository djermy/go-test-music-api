#!/bin/sh

PROJECT_DIR="/go/src/api"

cd "$(dirname "$0")/" || exit 1

docker run -it --rm \
    --name music-api \
    -w "${PROJECT_DIR}" \
    -e "API_DIR=${PROJECT_DIR}" \
    -e "air_wd=${PROJECT_DIR}" \
    -e GOFLAGS="-buildvcs=false" \
    -v "$(pwd)"/:"${PROJECT_DIR}" \
    -v ./.go-mod:/go/pkg/mod \
    --network host \
    cosmtrek/air:v1.51.0

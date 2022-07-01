#!/usr/bin/env bash
docker run --rm \
    --user "$(id -u):$(id -g)" \
    -v "$(pwd)"/chatws:/api \
    -v "$(pwd)"/chatws:/goclient \
    -v "$(pwd)"/html/lib:/jsclient \
    leehom/grpc-web-generators \
    protoc -I /api \
        --go_out=plugins=grpc,paths=source_relative:/goclient \
        --js_out=import_style=commonjs:/jsclient \
        --grpc-web_out=import_style=commonjs,mode=grpcwebtext:/jsclient \
        /api/chatws.proto

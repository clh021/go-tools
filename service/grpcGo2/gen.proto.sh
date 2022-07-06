#!/usr/bin/env bash
# leehom Chen clh021@gmail.com
docker run --rm \
    --user "$(id -u):$(id -g)" \
    -v "$(pwd)"/chat:/api \
    -v "$(pwd)"/chat:/goclient \
    -v "$(pwd)"/html/lib:/jsclient \
    leehom/grpc-web-generators \
    protoc -I /api \
        --go_out=plugins=grpc,paths=source_relative:/goclient \
        --js_out=import_style=commonjs:/jsclient \
        --grpc-web_out=import_style=commonjs,mode=grpcwebtext:/jsclient \
        /api/chat.proto

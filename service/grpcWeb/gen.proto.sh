#!/usr/bin/env bash
docker run \
    -v `pwd`/api:/api \
    -v `pwd`/time/goclient:/goclient \
    -v `pwd`/frontend/src/jsclient:/jsclient \
    leehom/grpc-web-generators \
    protoc -I /api \
        --go_out=plugins=grpc,paths=source_relative:/goclient \
        --js_out=import_style=commonjs:/jsclient \
        --grpc-web_out=import_style=commonjs,mode=grpcwebtext:/jsclient \
        /api/time/v1/time_service.proto

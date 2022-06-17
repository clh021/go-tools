#!/usr/bin/env bash
webPath=$(pwd)
servicePath="$(dirname "$webPath")"
docker run --rm \
    --user "$(id -u):$(id -g)" \
    -v "${servicePath}"/echovue:/api \
    -v "${servicePath}"/echovue:/goclient \
    -v "${servicePath}"/vue/src/lib:/jsclient \
    leehom/grpc-web-generators \
    protoc -I /api \
        --plugin=protoc-gen-ts=/usr/local/bin/protoc-gen-ts \
        --js_out=import_style=commonjs,binary:/jsclient \
        --ts_out=service=grpc-web:/jsclient \
        /api/echo.proto

# protoc \
#   --plugin=protoc-gen-ts=/usr/local/bin/protoc-gen-ts \
#   --plugin=protoc-gen-go=/golang/go/bin/protoc-gen-go \
#   -I ./api \
#   --js_out=import_style=commonjs,binary:/jsclient \
#   --go_out=plugins=grpc:/goclient \
#   --ts_out=service=grpc-web:/jsclient \
#   /api/echo.proto

# protoc -I /api \
#     --go_out=plugins=grpc,paths=source_relative:/goclient \
#     --js_out=import_style=commonjs:/jsclient \
#     --grpc-web_out=import_style=commonjs,mode=grpcwebtext:/jsclient \
#     /api/echo.proto
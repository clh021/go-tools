#!/usr/bin/env bash
# leehom Chen clh021@gmail.com
# 1. https://github.com/protocolbuffers/protobuf/releases
# 2. go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protos/helloworld.proto

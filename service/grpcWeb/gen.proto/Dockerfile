# docker build -t leehom/grpc-web-generators .

FROM ubuntu:latest

ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update \
    && apt-get install -y curl \
    && curl -sSL https://gitee.com/clh21/sh/raw/master/mirror.sh | sh
RUN apt-get update && apt-get install -y \
    automake \
    build-essential \
    git \
    libtool \
    make \
    npm \
    wget \
    unzip \
    libprotoc-dev \
    python3-pip \
    golang

## Install protoc

ENV PROTOBUF_VERSION 3.20.1

RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOBUF_VERSION/protoc-$PROTOBUF_VERSION-linux-x86_64.zip && \
    unzip protoc-$PROTOBUF_VERSION-linux-x86_64.zip -d /usr/local/ && \
    rm -rf protoc-$PROTOBUF_VERSION-linux-x86_64.zip

## Install protoc-gen-go

ENV PROTOC_GEN_GO_VERSION v1.5.2
ENV GOPROXY https://goproxy.cn
ENV GOROOT /usr/lib/go
ENV GOPATH /golang/go
ENV PATH $GOPATH/bin:$PATH
RUN git clone https://github.com/golang/protobuf /golang/go/src/github.com/golang/protobuf && \
    cd /golang/go/src/github.com/golang/protobuf && \
    git fetch --all --tags --prune && \
    git checkout tags/$PROTOC_GEN_GO_VERSION && \
    go install ./protoc-gen-go && \
    ln -s /golang/go/bin/protoc-gen-go /usr/local/bin/protoc-gen-go && \
    rm -rf /golang/go/src

## Install protoc-gen-grpc-web

ENV PROTOC_GEN_GRPC_WEB_VERSION 1.3.1

RUN wget -c https://github.com/grpc/grpc-web/releases/download/$PROTOC_GEN_GRPC_WEB_VERSION/protoc-gen-grpc-web-$PROTOC_GEN_GRPC_WEB_VERSION-linux-x86_64 \
    -O /usr/local/bin/protoc-gen-grpc-web \
    && chmod +x /usr/local/bin/protoc-gen-grpc-web

## Install protoc-gen-ts

# ENV PROTOC_GEN_TS_VERSION 0.15.0

RUN npm install ts-protoc-gen google-protobuf && \
    ln -s /node_modules/.bin/protoc-gen-ts /usr/local/bin/protoc-gen-ts

## Install protoc-gen-python

RUN pip3 install grpcio-tools

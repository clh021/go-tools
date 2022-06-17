#!/usr/bin/env bash
# leehom Chen clh021@gmail.com
# 启动 grpcVue 服务
# 本来想只通过 docker-compose.yaml 方式， docker-compose up --build 启动开发环境
# 由于此处是作为子父母开发，所以这里就必须通过脚本编译结果拷贝到临时目录来启动服务
set -ex

# 路径准备
OldPath=$(pwd)
SCRIPT_PATH=$(realpath "$0")
ProjectPath="$(dirname "$(dirname "$(dirname "$(dirname "$SCRIPT_PATH")")")")"

prepareServer() {
    pushd "$ProjectPath" > /dev/null
    pwd
#    ./build.sh
    cp ./bin/app ./service/grpcVue/vue/server/server
    popd > /dev/null
}

runService() {
  docker-compose up --build
}

main() {
  prepareServer
  runService
}

main

cd "$OldPath"
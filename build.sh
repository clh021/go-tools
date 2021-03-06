#!/usr/bin/env bash
# leehom Chen clh021@gmail.com
echo "正在准备编译标记版本"
gitTime=$(git log -1 --format=%at | xargs -I{} date -d @{} +%Y%m%d_%H%M%S)
# gitTime=$(date +%Y%m%d%H%M%S)
gitCID=$(git rev-parse HEAD)
gitTag=$(git tag --list --sort=version:refname 'v*' | tail -1)
gitCount=$(git log --pretty=format:'' | wc -l)/$(git rev-list --all --count)
buildStr="${gitTime}.${gitCID:0:7}.${gitTag}.${gitCount}"
echo "正在生成静态文件缓存"
go mod tidy
go generate



export CGO_ENABLED=0
go build -ldflags "-X main.build=${buildStr}" -o "bin/app"



# echo "正在检查编译的容器再利用"
# exist_builder=`docker ps -a | grep "gift_build" | wc -l`
# if [[ exist_builder -gt 0 ]]
# then

# echo "正在使用存在的容器编译……"
# docker start gift_build

# else

#     echo "正在创建新容器编译……"
#     docker run -it \
#         --name gift_build \
#         -v `pwd`:/host \
#         -e GOPROXY="https://goproxy.cn,direct" \
#         -w /host golang:alpine \
#         go build -ldflags "-X main.build=${gitTime}.${gitCID}" -o "bin.${gitTime}"
#     # docker run --rm golang:alpine go version
# fi
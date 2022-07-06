#!/usr/bin/env bash
# leehom Chen clh021@gmail.com
docker rm gift_test || true
# docker run -it --rm -v /home/working/html:/host golang:alpine /host/web service start

# docker run -d \
#   --name gift_test \
#   -v `pwd`/config:/config \
#   -v `pwd`:/host \
#   alpine \
#   /host/web service install



  # -v `pwd`/config:/config \
docker run -it \
  --name gift_test \
  -v `pwd`/nging_linux_amd64:/host \
  alpine \
  ash
# docker build -t tmp-server .
FROM scratch
COPY server /server
ENV APPINTO=grpcWeb
# 在 alpine 容器中运行 go 二进制文件时遇到了
# exec /server: no such file or directory
# 使用 export CGO_ENABLED=0 即可解决
ENTRYPOINT ["/server"]
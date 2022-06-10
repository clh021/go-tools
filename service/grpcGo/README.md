
Install Go plugins
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

grpc-go
```bash
git clone -b v1.47.0 --depth 1 https://github.com/grpc/grpc-go
```

Use
```bash
APPINTO=grpcGoServer ./bin/app
APPINTO=grpcGoClient ./bin/app
APPINTO=grpcGoClient ./bin/app -name=Leehom
APPINTO=grpcGoClient ./bin/app -name Leehom
```
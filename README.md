# be-api

proto for grpc





buf.lock用于控制buf的版本

使用以下命令升级更新buf.lock

```
buf mod update
```

使用前请先进行前置依赖安装将相关依赖下载到gopath的bin目录下面

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/go-errors/errors/cmd/protoc-gen-go-errors@latest
```

kratos的errors相关依赖按理来讲也应该改成本地,但是暂时还没搞明白该怎么做

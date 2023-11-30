# grpc
> 学习来自bilibili的码神之路，仅供参考。


# protobuf生成对应文件命令
```protobuf
//注意proto文件的位置；点表示在当前位置生成
protoc --go_out=. hello.proto //生成go语言文件，对应的是hello.pb.go
protoc --go-grpc_out=. hello.proto  //生成grpc文件，对应hello_grpc.pb.go
```



# source
bilibili狂神说的grpc教程，仅作为学习参考。

# install
1. 官网[下载](https://github.com/protocolbuffers/protobuf/releases?page=2)grpc编译器，然后把安装路径一直到bin文件夹下配置到环境变量中。
2. 下载go的依赖：
   ```
   go get google.golang.org/grpc
   
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```
3. 这儿有个小小的坑，`github.com/go1ang/protobuf/protoc-gen-go`和 `goog1e.go1ang.org/protobuf/cmd/protoc-gen-go`是不同的。区别在于*前者是旧版本*，**后者**是google接管后的**新版本**，他们之间的API是不同的，也就是说用于生成的命令，以及生成的文件都是不一样的。因为目前的grpc-go源码中的example用的是后者的生成方式，为了与时俱
   进，我们也采取最新的方式。
4. 在编写代码时，如果没有相关高亮显示，需要安装插件，直接搜索protoc

# 生成go、rpc代码
```
//注意proto文件的位置
protoc --go_out=. hello.proto
protoc --go-grpc_out=. hello.proto
```

# 代码编写
## 服务端编写
1. 创建gRPC Server对象，你可以理解为它是Server端的抽象对象
2. 将server(其包含需要被调闲的服务端接口)注册到gRPC Server的内部注册中心。 这样可以在接受到请求时，通过内部的服务发现，发现该服务端接口并转接进行罗银处理
3. 创建Listen,监听TCP端口
4. gRPC Server开始Iis.Accept,直到Stop
## 客户端编写
1. 创建与给定目标（服务端）的连接交互
2. 创建server的客户端对象
3. 发送RPC请求，等待同步响应，得到回调后返回响应结果
4. 输出晌应结果



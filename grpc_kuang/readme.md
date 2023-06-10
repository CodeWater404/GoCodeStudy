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

# 代码编写基本流程
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

# 认证--安全传输

1. `key`:服务器上的私钥文件，用于对发送给客户端数据的加密，以及对从客户端接收到数据的解密
2. `csr`:证书签名请求文件，用于提交给证书颁发机构(CA)对证书签名。
3. `crt`:由证书颁发机构(CA)签名后的证书，或者是开发者自签名的证书，包含证书持有人的信息，持有人的公钥，以及签署者的签名等信息。
4. `pem`:是基于Base64编码的证书格式，扩展名包括PEM、CRT和CER。

# SSL/TLS认证方式
![img.png](img.png)
![img_1.png](img_1.png)
![img_2.png](img_2.png)
简单来说就是安装openssl的安装包，然后添加到环境变量中去，最后用命令行来生成指定的密钥文件
```
//生成私钥
openssl genrsa -out server.key 2048

//生成证书（可以全部回车、不填的；如果需要填写可以看一些提示的什么）
openssl req -new -x509 -key server.key -out server.crt -days 36500

//生成csr文件
openssl req -new -key server.key -out server.csr

//更改openssl.cfg(Linux是openssl.cnf)
//1)复制一份你安装的openssl的bin目录里面的openssl.cnf文件到你项目所在的目录
//2)找到[CA_default],打开copy_extensions = copy(就是把前面的#去掉)
//3)找到[req].打开req_extensions = v3_req  #The extensions to add to a certificate request
//4)找到[v3_req],添加subjectAltName = @alt-names
//5〉添加新的标兹[alt_names],和标签字段
DNS.1 = *.codewater.com

//生成证书私胡test.key
openssl genpkey -algorithm RSA -out test.key

//通过私钥test.key生成证书请求文件test.csr(注意cfg和cnf)
openssl req -new -nodes -key test.key -out test.csr -days 3650 -subj "/C=cn/OU=myorg/O=mycomp/CN=myname" -config
./openssl.cnf -extensions v3_req
//test.csr是上面生成的证书请求文件，ca.crt/server.key是CA证书文件和key,用来对test.csr进行签名认证，这两个义件在第一部分生成。

//生成SAN证书 pem
openssl x509 -req -days 365 -in test.csr -out test.pem -CA server.crt -CAkey server.key -CAcreateserial -extfile
./openssl.cnf -extensions v3_req
```


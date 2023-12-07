# grpc
> 学习来自bilibili的码神之路，仅供参考。


# protobuf生成对应文件命令
```protobuf
//注意proto文件的位置；点表示在当前位置生成
protoc --go_out=. hello.proto //生成go语言文件，对应的是hello.pb.go
protoc --go-grpc_out=. hello.proto  //生成grpc文件，对应hello_grpc.pb.go
```

# grpc各种相关文件解释
在 [key目录](2_tls%2Fkey)下 ，这些文件是与 `TLS`（Transport Layer Security）有关的文件，用于在 gRPC 服务器和客户端之间建立安全的通信。这些文件的一般用途：

1. **server.crt**: 这是服务器的证书文件，包含服务器的公钥。它由证书颁发机构（CA）签名，用于验证服务器身份。

2. **server.csr**: 证书签名请求 (CSR)，是由服务器生成并发送给 CA 的文件，请求签名以获得证书。CSR 中包含了服务器的公钥信息。

3. **server.key**: 这是服务器的私钥文件，用于对通信进行加密。私钥应该被妥善保管，只有服务器才能访问它。

4. **server.srl**: 这是序列号文件，包含证书的序列号。每个由 CA 签名的证书都有一个唯一的序列号。

5. **client.csr**: 这是客户端的证书签名请求文件，类似于服务器的 CSR。客户端将其发送给 CA 以获取签名后的证书。

6. **client.key**: 这是客户端的私钥文件，用于对通信进行加密。客户端应妥善保管私钥。

7. **client.pem**: 这是客户端的证书文件，包含客户端的公钥。由 CA 签名，用于验证客户端身份。

在 gRPC 中，TLS 主要用于两个目的：

- **身份验证**: 通过证书验证，确保客户端和服务器都是它们声称的实体。

- **加密通信**: 使用公钥/私钥对加密通信，确保数据在传输过程中是安全的。

这些文件的生成和使用通常需要一定的步骤，包括创建自己的 CA 或使用现有的 CA 进行签名等。生成这些文件的具体步骤可能因为你使用的工具和 CA 不同而有所差异。


# openssl生成key 报错解决
1. `openssl req -new -x509 -key server.key -out server.crt -days 36500`,Can't open "C:\Program Files\Common Files\SSL/openssl.cnf" for reading, No such file or directory
   54700000:error:80000003:system library:BIO_new_file:No such process:crypto\bio\bss_file.c:67:calling fopen(C:\Program Files\Common Files\SSL/openssl.cnf, r)      
   54700000:error:10000080:BIO routines:BIO_new_file:no such file:crypto\bio\bss_file.c:75. 修改路径
   :
    
    ```
   //linux
   set OPENSSL_CONF=F:\Code\GoCode\grpc\code_god\2_tls\key\openssl.cfg
openssl req -new -x509 -key server.key -out server.crt -days 36500
    // windows
   $env:OPENSSL_CONF = "F:\Code\GoCode\grpc\code_god\2_tls\key\openssl.cfg"
    openssl req -new -x509 -key server.key -out server.crt -days 36500
   ```


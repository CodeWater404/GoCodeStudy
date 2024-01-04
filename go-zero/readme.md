# go-zero学习


```markdown
.
├── code              # 参考code码神之路(暂停在微服务生成rpc文件)
│   └── helloword     # 入门案列
├── go.mod
├── go.sum
└── readme.md         # 本文档

```

# 遇到的问题

## 1. 生成rpc代码失败
执行下列命令失败，有两种情况
```shell
goctl rpc protoc ./code/microhelloworld/mall/user/rpc/user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
```

1. 错误是`go.mod already exists `, 在`go.work`那一级目录下执行命令`go work use user`(user是对应的模块服务名，看你自己)
2. 错误是`invalid go module`：
   3. 一种情况是：你没有在模块项目下建立`go.mod`文件，执行命令`go mod init user`(user是对应的模块服务名，看你自己)
   4. 另一种情况是：你不是在`go.mod`那一级目录下执行的命令，这种直接在mod文件那一级目录下执行命令就行了


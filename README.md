# GoCodeStudy

> 中文 | [English](./readme-en.md)



记录学习`Golang`的历程，包括：基本语法、算法与数据结构、框架（gin、gorm...）、并发、面试题、设计模式....

详细内容参考目录介绍部分。

# 学习路线

1. 基本语法
2. web框架: gin(gin , learn more when i have time:beego , iris) gorm grpc zero(distributed)
3. 中间件: redis mq(rabbitmq) 



# 项目推荐

> 这部分目前关注比较少，后续看到有好的学习项目会继续更新。



## web项目

1. gin-vue-admin





# 目录介绍

```text
./
├── Gin_demo				gin框架联系的一些例子
├── LICENSE
├── NetTest					一个关于网络的笔试题
├── algorithm				算法，主要是acwing和力扣
├── bin
├── concurrent-program		并发编程
├── designPattern			设计模式
├── exercise				一些练习题，包括：并发、excel工具库、编译、笔试题
├── gin_demo_qimi			七米的gin学习代码	
├── gmp						gmp模型相关代码
├── go-redis				redis的操作使用
├── go-zero					go-zero的学习
├── google_go_study			一个学习golang的基本语法练习
├── grpc					grpc的学习
├── log-project				一个日志库的项目学习（未完成）
├── pkg
├── readme-en.md
├── readme.md
└── web_exercise_qimi       七米的付费视频代码
```



# 一些常用的命令

1. `go env`: 列出go的环境变量。

2. `go env -w GO111MODULE=on`; the parametere w means setting environment in go environment.This command "set go111module=on" is void.

	> 在 Go 1.11 版本之后，Go 引入了模块化开发的特性，默认启用了 Go Modules。GO111MODULE 是一个控制模块支持的环境变量，它有以下几个可选值：
	> auto：根据当前工作目录自动启用或禁用模块支持（默认值）。
	> on：强制启用模块支持，无视当前工作目录。
	> off：禁用模块支持，将使用旧版的 GOPATH 模式进行开发。
	> 通过执行 go env -w GO111MODULE=off 命令，您将模块支持设置为禁用状态，即使用旧版的 GOPATH 模式进行开发。

	> 请注意，这是一个全局设置，将会影响您所有的 Go 项目。如果您希望在某个特定的项目中启用模块支持，可以在该项目的根目录下创建一个名为 go.mod 的文件，或者执行 go mod init 命令来初始化一     个新的模块。(有时候命令报错，可以关闭终端重新尝试下)

3. `go get -u XXXX`: 获取XXXX包。
	-u表示更新已安装的包或模块到最新版本。如果已经安装了包，-u 标志将检查远程仓库是否有更新的版本，并将其下载并安装到本地。
	-d 只下载不安装
	-f 只有在你包含了 -u 参数的时候才有效，不让 -u 去验证 import 中的每一个都已经获取了，这对于本地 fork 的包特别有用
	-fix 在获取源码之后先运行 fix，然后再去做其他的事情
	-t 同时也下载需要为运行测试所需要的包
	-u 强制使用网络去更新包和它的依赖包
	-v 显示执行的命令信息

4. 


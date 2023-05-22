# GoCodeStudy
The way of learning to go language.

# branch

1. master: exercise to learn the code
2. upstream: github's exercise projects , and i always update original project's info.
    * Maybe we could name the branch with the project name , to represent a project?
3. exercise: i practice the code on my own. After finish an mini test demo, i will merge into the master branch.

# learning directing
1. basic grammer
2. web application: gin(gin , learn more when i have time:beego , iris) gorm grpc zero(distributed)
3. middleware: redis mq(rabbitmq) 

# project
1. gin-vue-admin

# catalogue introduction
|-- ─ Gin_demo: learn gin(kuangshen):https://www.bilibili.com/video/BV1fA411F7aM/?spm_id_from=333.337.search-card.all.click&vd_source=428d32ac5556a6a38659408b8c8fb403
|-- ─ LICENSE
|-- ─ NetTest: an interview question
|-- ─ README.md
|-- ─ githubExercise
`-- ─ google_go_study:learn the basic grammer of go

# some common commands
1. `go env`: list the environment variables for go
2. `go env -w GO111MODULE=on`; the parametere w means setting environment in go environment.This command "set go111module=on" is void.
    > 在 Go 1.11 版本之后，Go 引入了模块化开发的特性，默认启用了 Go Modules。GO111MODULE 是一个控制模块支持的环境变量，它有以下几个可选值：
    >    auto：根据当前工作目录自动启用或禁用模块支持（默认值）。
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



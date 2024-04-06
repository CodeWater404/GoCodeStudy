# GoCodeStudy

> [Chinese](./readme.md) | English
>
> This document is machine translated, if there is any error please correct.



Documenting the journey of learning `Golang`, including: basic syntax, algorithms and data structures, frameworks (gin, gorm...) , concurrency, interview questions, design patterns ....

Refer to the introduction section of the catalog for details.

# Learning Route

1. basic syntax
2. web framework: gin(gin , learn more when i have time:beego , iris) gorm grpc zero(distributed)
3. middleware: redis mq(rabbitmq) 



# Project Recommendations

> This part of the current attention is relatively small, the next see a good learning program will continue to update.



## web projects

1. gin-vue-admin





# Catalog introduction

```text
. /
├── Gin_demo 			Some examples of the gin framework connections.
├── LICENSE
├── NetTest 			A written question about networking
├── algorithm 			algorithms, mainly acwing and force buckling
├── bin
├── concurrent-program 	Concurrent Programming
├─ designPattern 		Design Patterns
├── exercise 			some practice questions, including: concurrency, excel tool library, compilation, written test questions
├── gin_demo_qimi		gin learning code of qimi	
├── gmp 				gmp model related code
├── go-redis 			redis operation use
├── go-zero 			go-zero learning
├── google_go_study 	A basic syntax exercise to learn golang
├── grpc 				grpc study
├── log-project 		A logging library project to learn (unfinished)
├── pkg
├── readme-en.md
├── readme.md
└── web_exercise_qimi 	qimi of paid video code
```



# Some common commands

1. `go env`: Lists the go environment variables.

2. `go env -w GO111MODULE=on`; the parametere w means setting environment in go environment.This command "set go111module=on" is void.

	> After Go version 1.11, Go introduced modular development with Go Modules enabled by default.GO111MODULE is an environment variable that controls module support and has the following optional values:
	> auto: Automatically enable or disable module support based on the current working directory (default).
	> on: force module support to be enabled, ignoring the current working directory.
	> off: disables module support and will use the old GOPATH mode for development.
	> By executing the go env -w GO111MODULE=off command, you set module support to the disabled state, i.e., development will be done using the old GOPATH mode.

	> Note that this is a global setting that will affect all your Go projects. If you want to enable module support for a particular project, you can create a file called go.mod in the root directory of that project, or run go mod init to initialize a new module. (Sometimes the command throws an error, so close the terminal and try again.)

3. `go get -u XXXX`: Get XXXX package.
	The -u command updates an installed package or module to the latest version. If the package is already installed, the -u flag will check if the remote repository has the

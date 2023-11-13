# 生成虚拟目标，对应下面的 all build run gotool clean help；如果不写，
#make会判断当前环境下是否有这个命令，如果有就不执行了
.PHONY: all build run gotool clean help

# 项目编译之后的文件名
BINARY="bluebell"

#
all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	# @表示不显示命令本身，只显示命令的输出结果；如果不加@，则会显示命令本身和输出结果
	@go run ./main.go conf/config.yaml

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"
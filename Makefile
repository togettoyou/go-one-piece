# 定义伪目标。不创建目标文件，而是去执行这个目标下面的命令。
.PHONY: all build-linux run gotool clean help

# 生成的二进制文件名
BINARY_NAME="go-one-server"

# 编译添加版本信息
versionDir = "${BINARY_NAME}/util/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)
ldflags="-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"

# 执行make命令时所执行的所有命令
all: gotool clean
	go build -v -ldflags ${ldflags} -o ${BINARY_NAME} .

# 交叉编译linux amd64版本
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags ${ldflags} -o ${BINARY_NAME} .

# 运行项目
run:
	@go run ./

# gotool工具
gotool:
    # 整理代码格式
	gofmt -w .
    # 代码静态检查
	go vet . | grep -v vendor;true

# 清理二进制文件
clean:
	@if [ -f ${BINARY_NAME} ] ; then rm ${BINARY_NAME} ; fi

# 帮助
help:
	@echo "make - 运行 gotool 工具, 并编译生成当前平台可运行的二进制文件"
	@echo "make build-linux - 编译 Go 代码, 生成linux amd64可运行的二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"
	@echo "make clean - 清理编译生成的二进制文件"

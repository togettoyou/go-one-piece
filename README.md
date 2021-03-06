<p align="center"><img src="https://user-images.githubusercontent.com/55381228/97401757-56c5ef80-192c-11eb-8822-67b458609093.png" width="256px"/></p>

# go-one-server

go-one-server 是基于 Gin 进行快速构建 RESTFUL API 服务的项目基础模板

# 脚手架安装

```
go get -u github.com/togettoyou/go-one-server/cmd/gos
```

# 使用

```
# 创建项目模板
gos new helloworld

cd helloworld
# 运行程序
gos run

# 生成 swag 文档
gos swag

# 使用镜像代理
export GOS_LAYOUT_REPO=https://github.com.cnpmjs.org/togettoyou/go-one-server.git
gos new helloworld

# 更多帮助
gos -h
```

## 集成

- [Request、Response处理](https://github.com/togettoyou/go-one-server/blob/main/handler/handler.go)
- [参数校验](https://github.com/togettoyou/go-one-server/blob/main/util/validator/validate.go)
- [全局配置文件](https://github.com/togettoyou/go-one-server/blob/main/util/conf/conf.go)
- [配置文件运行时热更新](https://github.com/togettoyou/go-one-server/blob/main/util/util.go)
- [日志记录](https://github.com/togettoyou/go-one-server/blob/main/util/logger/logger.go)
- [错误码](https://github.com/togettoyou/go-one-server/blob/main/util/errno/code.go)
- [gormV2配置](https://github.com/togettoyou/go-one-server/blob/main/model/model.go)
- [gin中间件](https://github.com/togettoyou/go-one-server/blob/main/router/middleware/README.md)
- [常用tools工具](https://github.com/togettoyou/gtools)
- [版本信息](https://github.com/togettoyou/go-one-server/blob/main/util/version/version.go)
- [swagger文档](https://github.com/togettoyou/go-one-server/blob/main/docs)

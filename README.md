# go-framework

基于`gin`的框架, 增加了很多的功能

## 如何运行

### 加载依赖

1. 通过 module 模式加载
    - `Goland` 配置 `Go->Go Modules`
        - 勾选 `Enable Go Moudles`
        - `Proxy: direct`
        - `Vendoring mode`
2. 通过 GOPATH 加载 (`待补充...`)

### 配置 env

```shell script
cp .env.example .env
# 主要配置以下项
# APP_HOST 服务监听的地址端口
# DB_* 连接到的数据库
# REDIS_* 连接到的 Redis
# DEBUG 是否开启 DEBUG 模式, 主要影响日志的输出
# APP_ENV 配置当前环境, 本地使用 local, 部署使用 production
```

### 编译

- `go build`
- `./go-framework`

### 开发

```shell script
go run main.go
```

## 功能

- [x] 连接`MySQL`数据库
- [x] 测试中加载框架
- [x] cmd中加载框架
- [x] 配置模块
- [x] 储存模块(储存接口)
- [x] 日志模块
- [x] 多通道日志
- [x] 连接到`Redis`
- [x] 命令行工具
- [x] 数据库迁移
- [x] 路由结构
- [x] ORM
- [x] Swagger
- [x] 中间件
- [x] 注册自定义表单验证规则
- [x] faker
- [x] cron 定时任务
- [x] faker
- [] redis 多连接配置
- [] mysql 多连接配置

## 相关文档

[gin](https://github.com/gin-gonic/gin)
[faker](https://github.com/bxcodec/faker)
[gorm](https://gorm.io/zh_CN/docs/)
[log](https://github.com/sirupsen/logrus)
[gin 模型验证tag文档](https://godoc.org/gopkg.in/go-playground/validator.v9)
[redis](https://github.com/go-redis/redis)
[swag](https://github.com/swaggo/swag)

# go-framework

gin 项目脚手架，包含完整的应用实例，清晰的依赖关系。

## Getting Help

Email: grianchan@gmail.com

## Feature

- 使用 `GORM` v2
- viper toml 配置文件加载

## 如何运行

前提：项目使用 module 模式运行。

### 配置 config/config.toml

```shell script
cp configs/default.tmol configs/local.toml
```

### 运行

```shell script
go run cmd/lgo/main.go
```

### 编译

- `go build -o lgo cmd/lgo/main.go`
- `./lgo --conf=configs/local.toml` // 固定读取 config 目录下的配置文件

### 测试

配置 env LGO_TEST_FILE={path}/configs/local.toml，测试中会自动加载框架

默认开启端口 `http://127.0.0.1:8080` 访问服务

## Roadmap

- [x] 数据库
- [x] http 接口测试
- [x] 加载配置
- [x] 日志
- [x] 缓存
- [x] 数据库
- [x] 路由结构
- [x] GORM
- [x] Swagger
- [x] 中间件
- [x] 注册自定义表单验证规则
- [x] faker 结构体数据填充
- [x] faker 数据填充
- [] 中间件
    - [x] JWT 中间件加入
    - [x] cors 跨域中间件 (github.com/gin-contrib/cors)
    - [] 请求速率 rate limiter 中间件
- [x] pprof 性能监控 (使用 `go tool pprof http://localhost:8090/debug/pprof/heap` 访问)
- [x] 配置模块改用 github.com/BurntSushi/toml
- [x] 使用 context 传递上下文
- [] grpc 服务端
- [] grpc 客户端
- [] kafka 消息队列
- [] 升级到 go1.18
- [] 使用 cobra 创建命令行工具

## 相关文档

[gin 框架](https://github.com/gin-gonic/gin)

[faker 结构体数据填充](https://github.com/bxcodec/faker)

[gorm ORM](https://gorm.io/zh_CN/docs/)

[log 日志](https://github.com/sirupsen/logrus)

[validate 参数校验](https://godoc.org/gopkg.in/go-playground/validator.v9)

[Redis](https://github.com/go-redis/redis)

[Swag](https://github.com/swaggo/swag)

[gin 官方中间件](https://github.com/gin-contrib)

[gin-pprof](https://github.com/gin-contrib/pprof)

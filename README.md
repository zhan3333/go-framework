# go-framework

基于 `gin` 框架 API 项目脚手架, 示例 API 采用 DDD领域驱动模型 应用设计, 项目整体拥有良好的代码规范与示例作用, 适合中中小型项目使用.

## Getting Help

Email: grianchan@gmail.com

## Feature

- `领域驱动(DDD)模型`: 示例 API 使用 `领域驱动(DDD)模型` 展示, 有较好的示例与规范作用
- `迁移版本控制(Migrator Version)`: 实现与使用上类似 Laravel Migrate 功能, 能够应对项目迭代上线时的数据库变更需求
- `多通道日志`
- `多数据库连接管理`
- 支持命令行模式 + HTTP 模式
- 协程下的定时任务功能
- 使用 `GORM` v2
- toml 配置文件加载
- 待完善的储存系统

## 如何运行

### 加载依赖

项目推荐使用 [module 模式](#配置使用module模式) 开发

### 配置 config/config.toml

```shell script
cp config/default.tmol config/local.toml
```

### 编译

- `go build`
- `./go-framework --config=local.tmol` // 固定读取 config 目录下的配置文件

### 开发

```shell script
go run main.go
```

### 测试

配置 env LGO_TEST_FILE={path}/config/local.toml，测试中会自动加载框架

默认开启端口 `http://127.0.0.1:8080` 访问服务

## Roadmap

- [x] 连接`MySQL`数据库
- [x] 测试中加载框架
- [x] console 模式中加载框架
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
- [x] faker 结构体数据填充
- [x] cron 定时任务
- [x] faker 数据填充
- [x] Redis 多连接配置
- [x] MySQL 多连接配置
- [x] console 模块优化 (使用 cobra)
- [] storage 模块优化
    - [] 目录映射
    - [] 使用文档
    - [] 动态资源访问 URL
    - [] 带 token 的资源访问
    - [] 带 失效时间的 资源访问
- [] 中间件
    - [x] JWT 中间件加入
    - [x] cors 跨域中间件 (github.com/gin-contrib/cors)
    - [] 请求速率 throttle 中间件
- [x] pprof 性能监控 (使用 `go tool pprof http://localhost:8090/debug/pprof/heap` 访问)
- [x] 配置模块改用 github.com/BurntSushi/toml
- [] 使用 context 传递上下文

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

## 测试

通过环境变量 `LGO_TEST_FILE=~/go-framework/config.local` 来指定测试时使用的配置文件

## 其它

### 配置使用module模式

mac:

```shell script
export GO111MODULE=on
```

windows:

```shell script
set GO111MODULE=on
```

输入

```shell script
go env
```

确认 `GO111MODULE` 选项

运行

```shell script
go mod vendor
```

下载依赖到 `vendor` 目录
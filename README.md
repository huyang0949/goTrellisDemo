# goTrellisDemo

这是一个 Go 微服务示例项目，用于演示基于 Gin、go-micro 和 RPC 调用的 HTTP 微服务基础结构。

## 项目结构

```text
cmd/
  api-gateway/      Gin HTTP 网关服务
  greeter-service/  go-micro RPC 服务
  local-dev/        本地开发启动入口，单进程同时启动网关和服务
api/
  proto/            gRPC/protobuf 接口契约
internal/
  app/              业务逻辑
  config/           配置读取和配置结构定义
  infra/            基础设施初始化
    mongo/          MongoDB 客户端
    redis/          Redis 客户端
    resources/      基础设施资源聚合初始化
  transport/
    grpc/           go-micro RPC 处理器
    http/           Gin 路由和 HTTP 处理器
configs/            环境配置文件
deployments/        Docker 和 Kubernetes 部署配置
scripts/            开发和运维脚本
```

## 本地运行

启动本地一体化开发入口：

```powershell
go run ./cmd/local-dev
```

调用 HTTP API：

```powershell
curl "http://127.0.0.1:8081/hello?name=Tom"
```

也可以分别启动服务：

```powershell
go run ./cmd/greeter-service
go run ./cmd/api-gateway
```

分别启动服务时，需要根据目标环境配置共享的 go-micro 服务注册中心。

## 配置

默认读取配置文件：

```text
configs/local.json
```

当前配置包含 MongoDB 和 Redis：

```json
{
  "mongodbUrl": "mongodb://10.0.0.10:27017/cook-king-server-dev",
  "redis": {
    "mode": "redis",
    "cluster": false,
    "port": 6579,
    "host": "10.0.0.10",
    "defaultCfg": {
      "redis": {
        "ttl": 300
      }
    }
  }
}
```

也可以通过环境变量指定配置文件：

```powershell
$env:APP_CONFIG="configs/local.json"
go run ./cmd/local-dev
```

启动时会初始化 MongoDB 和 Redis 客户端，但不会强制 ping 远端服务，避免开发网络不可达时阻塞本地服务启动。

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
  infra/            基础设施初始化
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

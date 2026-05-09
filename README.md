# goTrellisDemo

Go microservice demo project.

## Project Layout

```text
cmd/
  api-gateway/      Gin HTTP gateway
  greeter-service/  go-micro RPC service
  local-dev/        Local single-process launcher for gateway + service
api/
  proto/            gRPC/protobuf contracts
internal/
  app/              Business logic
  infra/            Infrastructure setup
  transport/
    grpc/           go-micro RPC handlers
    http/           Gin routes and handlers
configs/            Environment configuration files
deployments/        Docker and Kubernetes deployment files
scripts/            Development and operations scripts
```

## Run Locally

Start the local all-in-one launcher:

```powershell
go run ./cmd/local-dev
```

Call the HTTP API:

```powershell
curl "http://127.0.0.1:8081/hello?name=Tom"
```

Run the services separately:

```powershell
go run ./cmd/greeter-service
go run ./cmd/api-gateway
```

When running services separately, configure a shared go-micro registry for the target environment.

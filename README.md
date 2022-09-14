# king-gin-api

## Build & Run

```shell
go mod tidy

go build

go run main.go
## Docker & CI

暂无 有时间更新

Makefile 功能：

```shell
# 构建容器化所需的 Linux 可执行文件，并使用 upx 压缩
make

# 构建并运行 web（仅可运行于Linux）
make run-web

# 构建并打包为镜像，并在本地 Docker 中运行 web
make docker-web

# 构建并打包为镜像，并推送到远程仓库
make devops
```


## Air 热加载

```shell
go install github.com/cosmtrek/air@latest
 
air init

air
```

## Swagger 接口文档

安装 swagger 到 GOPATH:

```shell
go install github.com/swaggo/swag/cmd/swag
```

已集成 gin-swagger， 在项目目录下执行:

```shell
swag init -g routes/swagger.go -o "docs/web"
```

自动扫描注释生成 swagger.json，并添加 /swagger 路由

可修改 Air 参数实现热加载自动生成 Swagger：

```shell
[build]
  args_bin = []
  cmd = "swag init -g routes/swagger.go -o docs/web && go build -o ./tmp/main ."
  exclude_dir = ["assets", "tmp", "vendor", "testdata", "docs"]
```

访问 Swagger： http://localhost:8080/swagger/index.html

## pprof

```shell
go tool pprof http://localhost:8080/debug/pprof/heap
go tool pprof http://localhost:8080/debug/pprof/profile
go tool pprof http://localhost:8080/debug/pprof/block

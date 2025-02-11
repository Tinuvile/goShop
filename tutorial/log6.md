# 2025/2/5

## 前端界面

[![Bootstrap](https://img.shields.io/badge/Bootstrap-v5.1.3-purple?style=flat-square)](https://getbootstrap.com)
[![Font Awesome](https://img.shields.io/badge/Font_Awesome-v6.0.0-blue?style=flat-square)](https://fontawesome.com)
[![Go Template](https://img.shields.io/badge/Go_Template-v1.16.3-orange?style=flat-square)](https://golang.org/pkg/html/template/)

### 重构项目

goShop/app 放置项目代码

goShop/idl 放置接口文档

使用hz提供的注解文件[api.proto](https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/usage-protobuf/)

```protobuf
// frontend/home.proto
syntax = "proto3";

package frontend.home;

option go_package = "frontend/home";

import "idl/api.proto";

message Empty {}

service HomeService {
  rpc Home(Empty) returns(Empty) {
    option (api.get) = "/";
  }
}
```

### 使用cwgo生成代码

```powershell
PS F:\goShop\goShop\app\frontend> cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend -module github.com/Tinuvile/goShop/app/frontend -I ../../idl
```

这里我又踩了一个坑，因为一开始对go管理依赖不太了解，前面生成的代码中一直是 <strong>go.mod</strong> ,
在运行这个命令的时候一直报错，一开始以为是go版本的问题，后来发现需要用 <strong>go.work</strong>来管理多模块工作区。

| 特性                     | go.mod                           | go.work                         |  
|------------------------|----------------------------------|----------------------------------|  
| 依赖优先级                | 从远程仓库或本地缓存加载依赖         | 优先使用 use 中指定的本地模块代码  |  
| 版本控制是否提交           | ✅ 需提交到 Git                  | ❌ 不提交（开发环境临时配置）       |  
| 影响范围                 | 模块内生效                        | 全局生效（作用于当前工作区）        |  
| 典型操作命令              | go mod init / go get            | go work init / go work use      |

生成项目整体结构

```text
F:\goShop\goShop\app\frontend
├── biz/             # 业务核心代码（遵循 DDD 分层设计）
├── hertz_gen/       # IDL 生成的代码（Hertz 框架自动生成）
├── conf/            # 配置文件管理
├── script/          # 辅助脚本
│
├── go.mod           # Go 模块依赖定义
├── main.go          # 服务入口文件
└── docker-compose.yaml  # Docker 本地开发环境配置
```

- 入口层 <strong>main.go</strong> ：服务启动入口，集成配置加载、路由注册、中间件初始化

```go
func main() {
// init dal
// dal.Init()
address := conf.GetConf().Hertz.Address
h := server.New(server.WithHostPorts(address))

registerMiddleware(h)

// add a ping route to test
h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
})

router.GeneratedRegister(h)

h.Spin()
}
```

- 业务逻辑层 <strong>biz/</strong>

```text
biz/
├── handler/         # HTTP 请求处理器（对接路由）
├── router/          # 路由定义
├── service/         # 业务逻辑实现（核心）
├── dal/             # 数据访问层（Database Access Layer）
└── utils/           # 通用工具函数
```

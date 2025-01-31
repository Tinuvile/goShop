# 2025/1/31

## 服务注册

### 安装[Kitex的Consul拓展](https://www.cloudwego.io/zh/docs/kitex/tutorials/service-governance/service_discovery/consul/)

```powershell
PS F:\goShop\goShop> cd demo/auth
PS F:\goShop\goShop\demo\auth> go get github.com/kitex-contrib/registry-consul
go: added github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da
go: added github.com/fatih/color v1.13.0
go: added github.com/hashicorp/consul/api v1.20.0
go: added github.com/hashicorp/go-cleanhttp v0.5.1
go: added github.com/hashicorp/go-hclog v1.6.3
go: added github.com/hashicorp/go-immutable-radix v1.0.0
go: added github.com/hashicorp/go-rootcerts v1.0.2
go: added github.com/hashicorp/golang-lru v0.5.4
go: added github.com/hashicorp/serf v0.10.1
go: added github.com/kitex-contrib/registry-consul v0.1.0
go: added github.com/mattn/go-colorable v0.1.12
go: added github.com/mattn/go-isatty v0.0.14
go: added github.com/mitchellh/go-homedir v1.1.0
go: added github.com/mitchellh/mapstructure v1.4.1
```

### 编写服务端代码

```go
// demo/auth/main.go
func kitexInit() (opts []server.Option) {
	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
    if err != nil {
        log.Fatal(err)
    }
    opts = append(opts, server.WithRegistry(r))
}
```

### 添加依赖

```powershell
PS F:\goShop\goShop> go get github.com/kitex-contrib/registry-consul
go: downloading google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f
go: downloading google.golang.org/genproto/googleapis/rpc v0.0.0-20250127172529-29210b9bc287
go: downloading google.golang.org/genproto v0.0.0-20250127172529-29210b9bc287
go: added github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da
go: added github.com/fatih/color v1.13.0
go: added github.com/hashicorp/consul/api v1.20.0
go: added github.com/hashicorp/go-cleanhttp v0.5.1
go: added github.com/hashicorp/go-hclog v1.6.3
go: added github.com/hashicorp/go-immutable-radix v1.0.0
go: added github.com/hashicorp/go-rootcerts v1.0.2
go: added github.com/hashicorp/golang-lru v0.5.4
go: added github.com/hashicorp/serf v0.10.1
go: added github.com/kitex-contrib/registry-consul v0.1.0
go: added github.com/mattn/go-colorable v0.1.12
go: added github.com/mattn/go-isatty v0.0.14
go: added github.com/mitchellh/go-homedir v1.1.0
go: added github.com/mitchellh/mapstructure v1.4.1
go: upgraded google.golang.org/genproto/googleapis/rpc v0.0.0-20231106174013-bbf56f31fb17 => v0.0.0-20250127172529-29210b9bc287
```

### 启动[Consul](https://developer.hashicorp.com/consul/docs/intro)容器

```powershell
PS F:\goShop\goShop> docker compose up -d
time="2025-01-31T20:47:25+08:00" level=warning msg="F:\\goShop\\goShop\\docker-compose.yaml: the attribute `version` is obsolete, it will be ignored, please remove it to avoid potential confusion"
[+] Running 7/7
 ✔ consul Pulled                                                                                                                                                                                   63.7s 
   ✔ 27d6d74a7c1d Download complete                                                                                                                                                                 3.9s 
   ✔ a23041e1d950 Download complete                                                                                                                                                                 4.5s 
   ✔ d9a4cda1fc71 Download complete                                                                                                                                                                 1.9s 
   ✔ c0e228c45cba Download complete                                                                                                                                                                14.1s 
   ✔ d078792c4f91 Download complete                                                                                                                                                                 9.2s 
   ✔ fbdc56b403c6 Download complete                                                                                                                                                                 2.6s 
[+] Running 2/2
 ✔ Network goshop_default     Created                                                                                                                                                               0.1s 
 ✔ Container goshop-consul-1  Started
```

### 修改配置

```yaml
registry:
  registry_address:
    - 127.0.0.1:8500
  username: ""
  password: ""
```

### 启动服务

```powershell
PS F:\goShop\goShop\demo\auth> go run .
&{Env:test Kitex:{Service:auth Address::8888 LogLevel:info LogFileName:log/kitex.log LogMaxSize:10 LogMaxBackups:50 LogMaxAge:3} MySQL:{DSN:gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local} Redis:{Address:127.0.0.1:6379 Username: Password: DB:0} Registry:{RegistryAddress:[127.0.0.1:8500] Username: Password:}}
```

## 服务发现

### 编写客户端代码

```go
// demo/auth/cmd/client/main.go
package main

import (
	"context"
	"fmt"
	"github.com/Tinuvile/goShop/demo/auth/kitex_gen/auth"
	"log"

	"github.com/Tinuvile/goShop/demo/auth/kitex_gen/auth/authservice" // 客户端代码
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	// 1. 创建Consul解析器
	resolver, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal("创建解析器失败:", err)
	}

	// 2. 创建客户端（服务名称需要与服务端保持一致）
	authClient, err := authservice.NewClient(
		"auth", // 服务名称
		client.WithResolver(resolver),
	)
	if err != nil {
		log.Fatal("创建客户端失败:", err)
	}

	fmt.Printf("%v", authClient)

	// 3. 测试Token颁发功能
	deliverResp, err := authClient.DeliverTokenByRPC(context.Background(), &auth.DeliverTokenReq{
		UserId: 1001,
	})
	if err != nil {
		log.Fatal("颁发Token失败:", err)
	}
	log.Printf("颁发的Token: %s", deliverResp.Token)

	// 4. 测试Token验证功能
	verifyResp, err := authClient.VerifyTokenByRPC(context.Background(), &auth.VerifyTokenReq{
		Token: deliverResp.Token,
	})
	if err != nil {
		log.Fatal("验证Token失败:", err)
	}
	log.Printf("验证结果: %t", verifyResp.Res)
}
```

这里踩了个大坑，前一天使用Kitex生成的代码使用了错误的导入路径，authservice文件夹以及go.mod全部需要手动修改。

### 检测效果

```powershell
PS F:\goShop\goShop\demo\auth> go run cmd/client/main.go
&{0xc001b5f4d0}2025/01/31 23:33:10 颁发Token失败:service discovery error: no service found
exit status 1
```

可以用浏览器访问[Consul Web界面](http://localhost:8500/)

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


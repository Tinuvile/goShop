# 2025/2/13

## 用户服务

### Frontend Page的编写

<strong>sign-in.tmpl</strong>

```html
{{define "sign-in"}}
{{template "header" .}}
    <div class="row justify-content-center">
        <div class="col-md-4">
            <form method="POST" action="/auth/login">
                <div class="mb-3">
                    <label for="email" class="form-label">Email {{template "required"}}</label>
                    <input type="email" class="form-control" id="email" name="email">
                </div>
                <div class="mb-3">
                    <label for="password" class="form-label">Password {{template "required"}}</label>
                    <input type="password" class="form-control" id="password" name="password">
                </div>
                <div class="mb-3">
                    Don't have an account? Click <a href="/sign-up">here</a> to sign up.
                </div>
                <button type="submit" class="btn btn-primary">Sign In</button>
            </form>
        </div>
    </div>
{{template "footer" .}}
{{end}}
```

### 编写接口

<strong>auth_page.proto</strong>

```protobuf
syntax = "proto3";

package frontend.auth;

import "api.proto";
import "frontend/common.proto";

option go_package = "frontend/auth";

message LoginReq {
  string email = 1 [(api.form)="email"];
  string password = 2 [(api.form)="password"];
}

service AuthService {
  rpc login(LoginReq) returns(common.Empty) {
    option (api.post) = "/auth/login";
  }
}
```

### 使用代码生成能力生成代码

```powershell
C:\ProgramData\chocolatey\bin\make.exe -f F:/goShop/goShop/Makefile -C F:\goShop\goShop gen-frontend
make: Entering directory 'F:/goShop/goShop'
make: Leaving directory 'F:/goShop/goShop'

进程已结束，退出代码为 0

C:\ProgramData\chocolatey\bin\make.exe -f F:/goShop/goShop/Makefile -C F:\goShop\goShop gen-frontend
make: Entering directory 'F:/goShop/goShop'
make: Leaving directory 'F:/goShop/goShop'

进程已结束，退出代码为 0
```

### 写

使用Hertz的中间件[sessions](https://github.com/hertz-contrib/sessions)

```powershell
PS F:\goShop\goShop\app\frontend> go get github.com/hertz-contrib/sessions
go: added github.com/gorilla/context v1.1.2
go: added github.com/gorilla/securecookie v1.1.2
go: added github.com/gorilla/sessions v1.2.2
go: added github.com/hertz-contrib/sessions v1.0.3
```

选择基于Redis的

```go
func registerMiddleware(h *server.Hertz) {
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	h.Use(sessions.New("BirdShop", store))

	···
}
```
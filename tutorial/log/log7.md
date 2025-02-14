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


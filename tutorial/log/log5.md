# 2025/2/3

## 编码指南与开发规范

### 代码规范

- [Uber Go 语言编码规范](https://github.com/xxjwxc/uber_go_guide_cn)

- [protobuf 编码风格指南](https://protobuf.dev/programming-guides/style/)

### 错误码设计

#### [HTTP状态响应代码](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)

- Information responses (100-199)
- Successful responses (200-299)
- Redirection responses (300-399)
- Client error responses (400-499)
- Server error responses (500-599)

#### 自定义错误码

### 日志实践

[常见日志库对hertz做的适配](https://github.com/hertz-contrib/logger)

### 提交规范与版本命名规范

- [提交规范](https://www.conventionalcommits.org/zh-hans/v1.0.0/)

- [版本管理](https://semver.org/lang/zh-CN/)

## 微服务之间的通信

### RPC(Remote Procedure Call)通信

### HTTP请求(RESTful API)

[代码实例](https://github.com/cloudwego/hertz-examples)

### 消息中间件(Message Middleware)

```go
// auth/middleware/middleware
package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"time"
)

func Middleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		begin := time.Now()
		err = next(ctx, req, resp)
		fmt.Println("middleware end", time.Since(begin))
		return err
	}
}
```

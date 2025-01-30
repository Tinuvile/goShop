package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default()

	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		// 使用 RequestContext 的 Data 方法
		c.Data(consts.StatusOK, consts.MIMETextPlain, []byte("hello world"))
	})

	h.Spin()
}

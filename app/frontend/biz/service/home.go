package service

import (
	"context"
	"github.com/Tinuvile/goShop/app/frontend/hertz_gen/frontend/common"

	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	resp = make(map[string]any)
	items := []map[string]any{
		{"Name": "T-shirt 1", "Price": 100, "Picture": "/static/image/T-shirt1.png"},
		{"Name": "T-shirt 2", "Price": 150, "Picture": "/static/image/T-shirt2.png"},
		{"Name": "T-shirt 3", "Price": 200, "Picture": "/static/image/T-shirt3.png"},
		{"Name": "cup 1", "Price": 50, "Picture": "/static/image/cup1.png"},
		{"Name": "cup 2", "Price": 55, "Picture": "/static/image/cup2.png"},
		{"Name": "cup 3", "Price": 60, "Picture": "/static/image/cup3.png"},
		{"Name": "cup 4", "Price": 80, "Picture": "/static/image/cup4.png"},
	}

	resp["Title"] = "Hot Sales"
	resp["Items"] = items
	return resp, nil
}

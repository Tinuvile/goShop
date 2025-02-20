package auth

import (
	"context"
	"github.com/Tinuvile/goShop/app/frontend/biz/service"
	"github.com/Tinuvile/goShop/app/frontend/biz/utils"
	auth "github.com/Tinuvile/goShop/app/frontend/hertz_gen/frontend/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Login .
// @router /auth/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	//resp := &common.Empty{}
	//resp, err = service.NewLoginService(ctx, c).Run(&req)
	_, err = service.NewLoginService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.Redirect(consts.StatusOK, []byte("/"))

	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
	//utils.SendSuccessResponse(ctx, c, consts.StatusOK, "done!")
}

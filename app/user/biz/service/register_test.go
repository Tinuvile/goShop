package service

import (
	"context"
	"github.com/Tinuvile/goShop/app/user/biz/dal/mysql"

	user "github.com/Tinuvile/goShop/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
	"reflect"
	"testing"
)

func TestRegister_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "test@gmail.com",
		Password:        "123456",
		ConfirmPassword: "123456",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

func TestRegisterService_Run(t *testing.T) {
	type fields struct {
		ctx context.Context
	}
	type args struct {
		req *user.RegisterReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *user.RegisterResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RegisterService{
				ctx: tt.fields.ctx,
			}
			gotResp, err := s.Run(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Run() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}


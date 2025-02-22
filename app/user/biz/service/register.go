package service

import (
	"context"
	"errors"
	"github.com/Tinuvile/goShop/app/user/biz/model"
	"github.com/Tinuvile/goShop/demo/auth/biz/dal/mysql"
	user "github.com/Tinuvile/goShop/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("password not match")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(emailRegex, req.Email)
	if err != nil || !matched {
		return nil, errors.New("invalid email format")
	}

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(passwordHashed),
	}

	err = model.Create(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}

	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}

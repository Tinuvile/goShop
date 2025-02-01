package service

import (
	"context"
	"errors"
	"fmt"

	auth "github.com/Tinuvile/goShop/demo/auth/kitex_gen/auth"
	"github.com/golang-jwt/jwt/v5"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

func verifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		// 处理具体错误类型
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return false, fmt.Errorf("malformed token")
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			return false, fmt.Errorf("token expired or not active yet")
		}
		return false, fmt.Errorf("couldn't handle this token: %w", err)
	}

	return token.Valid, nil
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// Finish your business logic.
	// 1. 输入验证
	if req.Token == "" {
		return &auth.VerifyResp{Res: false}, fmt.Errorf("empty token")
	}

	// 2. 执行Token验证
	valid, err := verifyToken(req.Token)
	if err != nil {
		// 记录详细错误日志（实际项目应使用日志框架）
		fmt.Printf("Token验证失败: %v\n", err)
		return &auth.VerifyResp{Res: false}, nil
	}

	// 3. 返回验证结果
	return &auth.VerifyResp{Res: valid}, nil
}

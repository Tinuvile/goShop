package service

import (
	"context"
	"errors"
	"sync"
	"time"

	auth "github.com/Tinuvile/goShop/demo/auth/kitex_gen/auth"
	"github.com/golang-jwt/jwt/v5"
)

// 从配置获取参数
type JWTConfig struct {
	SecretKey   string `yaml:"secret_key"`
	ExpireHours int    `yaml:"expire_hours"`
}

// 自定义错误类型
var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired token")
)

// 缓存解析结果（适用于高频验证场景）
var tokenCache sync.Map

// 定义JWT签名密匙
const (
	SecretKey   = "your-secret-key-12345" // 通过配置文件引入而非硬编码
	TokenExpire = time.Hour * 24          // Token有效期
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

func generateToken(userID int32) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   string(rune(userID)),                            // 用户ID
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpire)), // 过期时间
		IssuedAt:  jwt.NewNumericDate(time.Now()),                  // 签发时间
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SecretKey))
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// Finish your business logic.
	// 1. 生产Token
	tokenString, err := generateToken(req.UserId)
	if err != nil {
		return nil, err
	}
	// 2. 构造响应
	return &auth.DeliveryResp{
		Token: tokenString,
	}, nil
}

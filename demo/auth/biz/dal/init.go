package dal

import (
	"github.com/Tinuvile/goShop/demo/auth/biz/dal/mysql"
	"github.com/Tinuvile/goShop/demo/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

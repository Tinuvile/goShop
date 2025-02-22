package dal

import (
	"github.com/Tinuvile/goShop/app/user/biz/dal/mysql"
	"github.com/Tinuvile/goShop/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

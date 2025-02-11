package dal

import (
	"github.com/Tinuvile/goShop/app/frontend/biz/dal/mysql"
	"github.com/Tinuvile/goShop/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

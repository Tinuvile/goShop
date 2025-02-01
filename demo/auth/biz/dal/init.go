package dal

import (
	"github.com/Tinuvile/goShop/demo/auth/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}

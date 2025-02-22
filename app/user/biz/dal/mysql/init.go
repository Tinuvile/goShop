package mysql

import (
	"fmt"
	"github.com/Tinuvile/goShop/app/user/biz/model"
	"github.com/Tinuvile/goShop/app/user/conf"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)

	if err != nil {
		panic(fmt.Sprintf("连接失败: %v | DSN: %s", err, dsn))
	}

	if err := DB.AutoMigrate(model.User{}); err != nil {
		panic(fmt.Sprintf("迁移失败: %v", err))
	}
}

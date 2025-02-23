package mysql

import (
	"fmt"
	"github.com/Tinuvile/goShop/app/user/biz/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(
		//conf.GetConf().MySQL.DSN,
		"%s:%s@tcp(%s:3306)/user?charset=utf8mb4&parseTime=True&loc=Local",
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

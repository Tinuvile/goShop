package mysql

import (
	"fmt"
	"github.com/Tinuvile/goShop/demo/auth/conf"
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
		os.Getenv("MYSQL_ROOT_DATABASE"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	// 测试使用
	type Version struct {
		Version string
	}

	var v Version

	err = DB.Raw("select version() as version").Scan(&v).Error

	if err != nil {
		panic(err)
	}

	fmt.Println(v)
}

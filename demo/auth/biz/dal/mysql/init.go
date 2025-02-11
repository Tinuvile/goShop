package mysql

import (
	"fmt"
	"github.com/Tinuvile/goShop/demo/auth/biz/model"
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
		os.Getenv("MYSQL_DATABASE"))

	fmt.Println("DSN:", dsn) // 打印 DSN 以便调试

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

	// 自动迁移
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		return
	}

	fmt.Println(v)
}

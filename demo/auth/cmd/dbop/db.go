package main

import (
	"fmt"
	"github.com/Tinuvile/goShop/demo/auth/biz/dal"
	"github.com/Tinuvile/goShop/demo/auth/biz/dal/mysql"
	"github.com/Tinuvile/goShop/demo/auth/biz/model"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	dal.Init()
	// CURD
	mysql.DB.Create(&model.User{Email: "admin@gmail.com", Password: "jszhiewdsh"})                    // 创建
	mysql.DB.Model(&model.User{}).Where("email=?", "admin@gmail.com").Update("password", "zzhdshalt") // 修改

	var row model.User
	mysql.DB.Model(&model.User{}).Where("email=?", "admin@gmail.com").First(&row) // 查找

	fmt.Println("row:%+v\n", row)
	
	mysql.DB.Where("email=?", "admin@gmail.com").Delete(&model.User{}) // 删除

	mysql.DB.Unscoped().Where("email=?", "admin@gmail.com").Delete(&model.User{}) // 强制删除
}

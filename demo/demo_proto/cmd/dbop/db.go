package main

import (
	"github.com/joho/godotenv"
	"github.com/yuefan-mo/studymall/demo/demo_proto/biz/dal"
	"github.com/yuefan-mo/studymall/demo/demo_proto/biz/dal/mysql"
	"github.com/yuefan-mo/studymall/demo/demo_proto/biz/model"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()

	//数据库表中增加用户
	//mysql.DB.Create(&model.User{Email: "demo@example.com", Password: "weksdfdsf"})

	// //更改表中的信息
	// mysql.DB.Model(&model.User{}).Where("email=?", "demo@example.com").Update("password", "23432543")

	// //读取表中的数据
	// var row model.User
	// mysql.DB.Model(&model.User{}).Where("email=?", "demo@example.com").First(&row)
	// fmt.Printf("row:%+v\n", row)

	//删除表中的数据
	mysql.DB.Where("email=?", "demo@example.com").Delete(&model.User{})

	//强制删除
	mysql.DB.Unscoped().Where("email=?", "demo@example.com").Delete(&model.User{})

}

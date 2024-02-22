package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name  string
	Users []User `gorm:"many2many:groups_user"` // many2many
}

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	// 使用sqlite
	dsn := "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db.Debug()
	if err != nil {
		panic("failed to connect database")
	}
	// mock data
	db.AutoMigrate(&User{}, &Group{})
	user := User{
		Model: gorm.Model{ID: 1},
		Name:  "jinzhu", Age: 18}
	db.Create(&user)
	// 创建
	// 先删除Group表中的数据
	db.Exec("delete from group")
	db.Exec("delete from groups_user")
	// 创建一个group
	user.Name = "jinzhu2"
	db.Save(&Group{
		Name: "group1",
		Users: []User{
			user,
		},
	})
	// 查询
	var u =User{
		Model: gorm.Model{ID: 1},
	}
	db.First(&u)	
	fmt.Println(u)
}

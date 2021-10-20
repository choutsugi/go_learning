package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model // 包含字段：ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string
	Gender     string
	Hobby      string
}

func main() {
	dst := "root:root1234@tcp(192.168.99.100:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	if err != nil {
		panic("failed to connect database!")
	}

	// 自动迁移：根据结构体生成表。
	_ = db.AutoMigrate(&UserInfo{})

	// 增
	db.Create(&UserInfo{
		Model:  gorm.Model{},
		Name:   "shinin",
		Gender: "Mr.",
		Hobby:  "game",
	})

	// 查
	var user UserInfo
	db.First(&user, 1)                   // 根据主键查询
	db.First(&user, "hobby = ?", "game") // 查询 hobby 字段值为 game 的记录

	// 改
	// 更新一个字段
	db.Model(&user).Update("hobby", "code") // 将查询结果中的 hobby 字段的值修改为 code
	// 更新多个字段
	db.Model(&user).Updates(UserInfo{
		Model:  gorm.Model{},
		Name:   "lettredamour",
		Gender: "Mr.",
		Hobby:  "Rain",
	})
	db.Model(&user).Updates(map[string]interface{}{
		"Name":   "shinrin",
		"Gender": "Mrs.",
		"Hobby":  "Wind",
	})

	// 删
	db.Delete(&user, 1)

}

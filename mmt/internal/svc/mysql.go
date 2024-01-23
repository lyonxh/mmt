package svc

import (
	"mmt/mmt/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(DataSource string) *gorm.DB {
	if db, err := gorm.Open(mysql.Open(DataSource), &gorm.Config{}); err != nil {
		panic("Error: connect to mysql fail")
	} else {
		// TODO 数据库初始化优化
		if err := db.AutoMigrate(
			model.MmtUsers{},
		); err != nil {
			panic("初始化表失败")
		}
		return db
	}
}

package model

import "gorm.io/gorm"

type MmtProject struct {
	gorm.Model
	Name string `json:"name" gorm:"common:项目名"`
}

func (MmtProject) TableName() string {
	return "mmt_project"
}

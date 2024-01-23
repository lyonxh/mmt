package model

import "gorm.io/gorm"

type MmtConfiguration struct {
	gorm.Model
}

func (MmtConfiguration) TableName() string {
	return "mmt_configuration"
}

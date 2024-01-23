package model

import "gorm.io/gorm"

type MmtUsers struct {
	gorm.Model
	UserName string `json:"username" gorm:"comment:用户名"`
	NickName string `json:"nickname" gorm:"comment:昵称"`
	PassWord string `json:"password" gorm:"comment:密码"`
	Mobile   string `json:"mobile" gorm:"comment:手机"`
	IdCard   string `json:"idcard" gorm:"comment:身份证"`
	Email    string `json:"email" gorm:"comment:邮箱"`
}

func (MmtUsers) TableName() string {
	return "mmt_user"
}

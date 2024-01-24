package model

import "gorm.io/gorm"

type MmtWebHook struct {
	gorm.Model
	HookType string `json:"hook_type" gorm:"comment:hook类型"`
	Url      string `json:"url" gorm:"comment:hookUrl"`
	Status   int    `json:"status" gorm:"comment:hook状态"`
	Token    string `json:"token" gorm:"comment:token"`
	AppId    uint	`json:"appId"`
}

func (MmtWebHook) TableNname() string {
	return "mmt_webhook"
}

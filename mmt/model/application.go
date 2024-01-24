package model

type MmtApplicationn struct {
	Name               string             `json:"name"`
	AppIcon            string             `json:"app_icon"` // app图标
	NickName           string             `json:"nick_name"`
	NameSpace          string             `json:"namespace" gorm:"default:default"` //命名空间
	Description        string             `json:"description"`
	Status             int                `json:"status" gorm:"default:0"` // 状态
	Configurations     []MmtConfiguration `json:"configurations"`          // 配置 cd
	CurrentEnvReplicas map[string]int     `json:"currentEnvReplicas"`      // 当前环境应用部署实例数，避免与配置页面数据冲突
	Webhook            []MmtWebHook       `json:"webhook" gorm:"foreignKey:AppId"`
}

func (MmtApplicationn) TableNname() string {
	return "mmt_application"
}

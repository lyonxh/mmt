package model

type MmtApplicationn struct {
	Name               string          `json:"name"`
	AppIcon            string          `json:"appIcon"` // app图标
	Code               string          `json:"code"`
	Alias              string          `json:"alias"`
	Project            string          `json:"project"`
	Description        string          `json:"description"`
	Status             int             `json:"status"`             // 状态
	Configurations     []Configuration `json:"configurations"`     // 配置
	CurrentEnvReplicas map[string]int  `json:"currentEnvReplicas"` // 当前环境应用部署实例数，避免与配置页面数据冲突
	Webhook            []MmtWebHook    `json:"webhook"`
}

func (MmtApplicationn) TableNname() string {
	return "mmt_application"
}

type Configuration struct {

}

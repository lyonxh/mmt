package model

import "gorm.io/gorm"

type MmtConfiguration struct {
	gorm.Model
	Type string `json:"type"` //配置类型
	// 基础配置
	Kind       string      `json:"kind"`          // App类型，deployment无状态；statefulset有状态
	GitRepo    *GitRepo    `json:"gitRepo"`       // 代码仓库
	GitBranch  []GitBranch `json:"autoGitBranch"` // 自动构建代码分支
	ImageRepo  string      `json:"imageRepo" `    // 镜像仓库地址
	ImageName  string      `json:"imageName" `    // 镜像名称
	Dockerfile string      `json:"dockerfile"`
	Replicas   int         `json:"replicas"` // 实例数量

}

func (MmtConfiguration) TableName() string {
	return "mmt_configuration"
}

type GitBranch struct {
	MatchType  string `json:"matchType,omitempty" validate:"oneof=exactMatch regMatch"` // 匹配类型，exactMatch精确匹配和regularMatch正则匹配
	MatchValue string `json:"matchValue,omitempty" validate:"required"`                 // 匹配值，例如：devel-.*
}

type GitRepo struct {
	GitType     string `json:"gitType,omitempty" description:"代码仓库类型" validate:"omitempty,oneof=gitlab gitee gitea github"` // 代码仓库类型，gitlab gitee gittea github
	GitURL      string `json:"gitURL,omitempty" description:"git地址"`                                                        // 代码仓库地址
	GitUser     string `json:"gitUser,omitempty"  description:"git地址"`                                                      // 代码仓库用户
	GitPassword string `json:"gitPassword,omitempty" description:"git仓库用户密码"`                                               // 代码仓库密码
}

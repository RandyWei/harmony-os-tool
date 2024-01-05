package models

type Module struct {
	Id   string `json:"id"`
	Name string `json:"name"` // 模块名称
	Type string `json:"type"` // 模块类型
	Apps []App  `json:"apps"` // 模块应用列表
}

package models

type Device struct {
	Id      string `json:"id"`
	Product string `json:"product"`
	Model   string `json:"model"` //型号
	Brand   string `json:"brand"` //品牌
}

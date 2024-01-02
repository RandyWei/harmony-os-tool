package models

type App struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	RelatedIds  []string `json:"related_ids"` //相关联的package id
	Installed   bool     `json:"installed"`
	Icon        string   `json:"icon"`
	Version     string   `json:"version"`
}

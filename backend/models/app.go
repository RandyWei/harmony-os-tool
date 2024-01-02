package models

type App struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Installed   bool   `json:"installed"`
	Icon        string `json:"icon"`
	Version     string `json:"version"`
}

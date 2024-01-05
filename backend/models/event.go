package models

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type EventName string

const (
	Event_Refresh_App_List EventName = "refresh_app_list"
)

// 事件
type Event struct {
	Ctx  context.Context
	Name EventName
	Data EventData
}

type EventData struct {
	Data interface{} `json:"data"`
	Type string      `json:"type"`
}

func (event *Event) Send() {
	runtime.EventsEmit(event.Ctx, string(event.Name), event.Data)
}

package utils

import (
	"context"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// 日志打印
func Log(ctx context.Context, message string) {
	runtime.LogInfo(ctx, message)
}

func LogE(ctx context.Context, message string) {
	runtime.LogError(ctx, message)
}

func Exits(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// 如果目录不存在则创建目录
// 成功返回nil，不成功返回目录
func MkDir(path string) error {
	if !Exits(path) {
		err := os.MkdirAll(path, 0755)
		return err
	}
	return nil
}

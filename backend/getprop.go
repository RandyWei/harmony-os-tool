package backend

import "context"

// 通过adb shell获取设备型号
func GetProp(ctx context.Context, appDir string) (string, error) {
	return AdbShellCommand(ctx, "getprop", "|", "grep", "product")
}

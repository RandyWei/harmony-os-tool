package backend

import "context"

// 查看当前设备内存情况
// ```free -h```
func Free(ctx context.Context) (string, error) {
	return AdbShellCommand(ctx, "free", "-h")
}

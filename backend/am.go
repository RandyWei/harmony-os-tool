package backend

import (
	"context"
	"fmt"
)

// 启动进程
// ```am start com.huawei.android.hwouc/com.huawei.android.hwouc.ui.activities.MainEntranceActivity```
func StartApp(ctx context.Context, packageName string, className string) (string, error) {
	return AdbShellCommand(ctx, "am", "start", fmt.Sprintf("%s/%s", packageName, className))
}

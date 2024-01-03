package backend

import "fmt"

// 启动进程
// ```am start com.huawei.android.hwouc/com.huawei.android.hwouc.ui.activities.MainEntranceActivity```
func StartApp(packageName string, className string) (string, error) {
	return AdbShellCommand("am", "start", fmt.Sprintf("%s/%s", packageName, className))
}

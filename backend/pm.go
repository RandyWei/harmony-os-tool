package backend

import "fmt"

// 检测应用是否安装
//
// ```pm list packages | grep "fastapp"```
func CheckInstalled(packageName string) (bool, error) {
	result, err := AdbShellCommand("pm", "list", "packages", "|", "grep", packageName)
	if err != nil {
		return false, err
	}
	fmt.Printf("result:%s\n", result)
	return result != "", nil
}

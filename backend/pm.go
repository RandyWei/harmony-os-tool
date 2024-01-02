package backend

import (
	"fmt"
	"strings"

	"icu.bughub.app/harmonyos-tool/backend/constants"
	"icu.bughub.app/harmonyos-tool/backend/models"
)

// 检测应用是否安装
//
// ```pm list packages | grep "fastapp"```
func CheckInstalled(packageName string) (bool, error) {
	result, err := AdbShellCommand("pm", "list", "packages", "|", "grep", packageName)
	if err != nil {
		return false, err
	}
	fmt.Printf("result:%s\n", result)
	return strings.Contains(result, packageName), nil
}

func ListApps() ([]models.App, error) {
	apps := make([]models.App, 0)
	for _, v := range constants.Apps1 {
		installed, _ := CheckInstalled(v.Id)
		v.Installed = installed
		apps = append(apps, v)
	}
	return apps, nil
}

// 卸载应用
// ```pm uninstall --user 0 com.huawei.fastapp```
func UninstallApp(packageName string) (bool, error) {
	result, err := AdbShellCommand("pm", "uninstall", "--user", "0", packageName)
	fmt.Printf("UninstallApp:%s\n", result)
	return result == "Success", err
}

// 装回应用
// ```pm install-existing --user 0 com.huawei.fastapp```
func InstallExistingApp(packageName string) (bool, error) {
	result, err := AdbShellCommand("pm", "install-existing", "--user", "0", packageName)
	fmt.Printf("InstallExistingApp:%s\n", result)
	return strings.Contains(result, "installed"), err
}

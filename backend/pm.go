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
	return strings.Contains(result, packageName), nil
}

func CheckEnabled(packageName string) (bool, error) {
	result, err := AdbShellCommand("dumpsys", "meminfo", packageName)
	if err != nil {
		return false, err
	}
	fmt.Printf("CheckEnabled:%s\n", result)
	return !strings.Contains(result, "No process"), nil
}

// 系统服务需要通过```dumpsys meminfo com.huawei.android.hwouc```查询是否占用内存来判断是否启用禁用
// No process 找不到进程的时候就说明已经禁用
func ListApps0() ([]models.App, error) {

	apps := make([]models.App, 0)
	for _, v := range constants.Apps0 {
		installed, _ := CheckEnabled(v.Id)
		v.Installed = installed
		apps = append(apps, v)
	}
	return apps, nil
}

func ListApps1() ([]models.App, error) {

	apps := make([]models.App, 0)
	for _, v := range constants.Apps1 {
		installed, _ := CheckInstalled(v.Id)
		v.Installed = installed
		apps = append(apps, v)
	}

	return apps, nil
}

func ListApps2() ([]models.App, error) {

	apps := make([]models.App, 0)
	for _, v := range constants.Apps2 {
		installed, _ := CheckInstalled(v.Id)
		v.Installed = installed
		apps = append(apps, v)
	}
	return apps, nil
}

func ListApps3() ([]models.App, error) {
	apps := make([]models.App, 0)
	for _, v := range constants.Apps3 {
		installed, _ := CheckInstalled(v.Id)
		v.Installed = installed
		apps = append(apps, v)
	}
	return apps, nil
}

func ListApps4() ([]models.App, error) {
	apps := make([]models.App, 0)
	for _, v := range constants.Apps4 {
		installed, _ := CheckInstalled(v.Id)
		v.Installed = installed
		apps = append(apps, v)
	}

	return apps, nil
}

func ListApps5() ([]models.App, error) {
	apps := make([]models.App, 0)
	for _, v := range constants.Apps5 {
		installed, _ := CheckInstalled(v.Id)
		v.Installed = installed
		apps = append(apps, v)
	}

	return apps, nil
}

// 卸载应用
// ```pm uninstall --user 0 com.huawei.fastapp```
// TODO 这里应该判断卸载了所有应用
func UninstallApp(packageName string, relatedIds []string) (bool, error) {
	result, err := AdbShellCommand("pm", "uninstall", "--user", "0", packageName)
	for _, v := range relatedIds {
		_, err = AdbShellCommand("pm", "uninstall", "--user", "0", v)
	}
	fmt.Printf("UninstallApp:%s\n", result)
	return result == "Success", err
}

// 装回应用
// ```pm install-existing --user 0 com.huawei.fastapp```
func InstallExistingApp(packageName string, relatedIds []string) (bool, error) {
	result, err := AdbShellCommand("pm", "install-existing", "--user", "0", packageName)
	for _, v := range relatedIds {
		_, err = AdbShellCommand("pm", "install-existing", "--user", "0", v)
	}
	fmt.Printf("InstallExistingApp:%s\n", result)
	return strings.Contains(result, "installed"), err
}

// 禁用应用
// ```pm disable-user com.huawei.android.hwouc```
func DisableApp(packageName string) (bool, error) {
	result, err := AdbShellCommand("pm", "disable-user", packageName)
	fmt.Printf("DisableApp:%s\n", result)
	return strings.Contains(result, "disabled"), err
}

// 启用应用
// ```pm enable com.huawei.android.hwouc```
// com.huawei.android.hwouc/com.huawei.android.hwouc.ui.activities.MainEntranceActivity
func EnableApp(packageName string) (bool, error) {
	result, err := AdbShellCommand("pm", "enable", packageName)
	//启动更新进程，以确保服务运行起来，这样就可以更新前端状态
	StartApp("com.huawei.android.hwouc", "com.huawei.android.hwouc.ui.activities.MainEntranceActivity")
	return strings.Contains(result, "enabled"), err
}

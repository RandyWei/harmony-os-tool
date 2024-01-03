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

func ListApps1() ([]models.App, error) {

	apps := make([]models.App, 0)
	for _, v := range constants.Apps0 {
		installed, _ := CheckInstalled(v.Id)
		v.Installed = installed
		apps = append(apps, v)
	}
	return apps, nil
}

func ListApps2() ([]models.App, error) {

	apps := make([]models.App, 0)
	for _, v := range constants.Apps1 {
		installed, _ := CheckInstalled(v.Id)
		v.Installed = installed
		apps = append(apps, v)
	}

	return apps, nil
}

func ListApps3() ([]models.App, error) {

	apps := make([]models.App, 0)
	for _, v := range constants.Apps2 {
		installed, _ := CheckInstalled(v.Id)
		v.Installed = installed
		apps = append(apps, v)
	}
	return apps, nil
}

func ListApps4() ([]models.App, error) {
	apps := make([]models.App, 0)
	for _, v := range constants.Apps3 {
		installed, _ := CheckInstalled(v.Id)
		v.Installed = installed
		apps = append(apps, v)
	}
	return apps, nil
}

func ListApps5() ([]models.App, error) {
	apps := make([]models.App, 0)
	for _, v := range constants.Apps4 {
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

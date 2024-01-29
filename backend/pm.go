package backend

import (
	"context"
	"fmt"
	"strings"

	"icu.bughub.app/harmonyos-tool/backend/constants"
	"icu.bughub.app/harmonyos-tool/backend/models"
	"icu.bughub.app/harmonyos-tool/backend/utils"
)

// 检测应用是否安装
//
// ```pm list packages | grep "fastapp"```
func CheckInstalled(ctx context.Context, packageName string) (bool, error) {
	result, err := AdbShellCommand(ctx, "pm", "list", "packages", "|", "grep", packageName)
	if err != nil {
		utils.LogE(ctx, err.Error())
		return false, err
	}
	// utils.Log(ctx, result)
	return strings.Contains(result, packageName), nil
}

// 系统服务需要通过```dumpsys meminfo com.huawei.android.hwouc```查询是否占用内存来判断是否启用禁用
func CheckEnabled(ctx context.Context, packageName string) (bool, error) {
	result, err := AdbShellCommand(ctx, "dumpsys", "meminfo", packageName)
	if err != nil {
		utils.LogE(ctx, err.Error())
		return false, err
	}
	return !strings.Contains(result, "No process"), nil
}

// 系统服务需要通过```dumpsys meminfo com.huawei.android.hwouc```查询是否占用内存来判断是否启用禁用
// No process 找不到进程的时候就说明已经禁用
func ListApps0(ctx context.Context) ([]models.App, error) {

	apps := make([]models.App, 0)
	for _, v := range constants.Apps0 {
		installed, _ := CheckEnabled(ctx, v.Id)
		v.Installed = installed
		apps = append(apps, v)
	}
	return apps, nil
}

func ListApps1(ctx context.Context) ([]models.App, error) {

	apps := make([]models.App, 0)
	for _, v := range constants.Apps1 {
		installed, _ := CheckInstalled(ctx, v.Id)
		v.Installed = installed
		apps = append(apps, v)
	}

	return apps, nil
}

func ListApps2(ctx context.Context) ([]models.App, error) {

	apps := make([]models.App, 0)
	for _, v := range constants.Apps2 {
		installed, _ := CheckInstalled(ctx, v.Id)
		v.Installed = installed
		apps = append(apps, v)
	}
	return apps, nil
}

func ListApps3(ctx context.Context) ([]models.App, error) {
	apps := make([]models.App, 0)
	for _, v := range constants.Apps3 {
		installed, _ := CheckInstalled(ctx, v.Id)
		v.Installed = installed
		apps = append(apps, v)
	}
	return apps, nil
}

func ListApps4(ctx context.Context) ([]models.App, error) {
	apps := make([]models.App, 0)
	for _, v := range constants.Apps4 {
		installed, _ := CheckInstalled(ctx, v.Id)
		v.Installed = installed
		apps = append(apps, v)
	}

	return apps, nil
}

func ListApps5(ctx context.Context) ([]models.App, error) {
	apps := make([]models.App, 0)
	for _, v := range constants.Apps5 {
		installed, _ := CheckInstalled(ctx, v.Id)
		v.Installed = installed
		apps = append(apps, v)
	}

	return apps, nil
}

func ListModules(ctx context.Context, brand string) ([]models.Module, error) {
	modules := make([]models.Module, len(constants.Data[brand]))
	copy(modules, constants.Data[brand][:])
	return modules, nil
}

func ListModuleApps(ctx context.Context, brand string, id string) ([]models.App, error) {
	apps := make([]models.App, 0)
	for _, v := range constants.Data[brand] {
		if v.Id == id {
			if v.Type == "disable" {
				for _, v := range v.Apps {
					installed, _ := CheckEnabled(ctx, v.Id)
					v.Installed = installed
					apps = append(apps, v)
				}
			} else {
				for _, v := range v.Apps {
					installed, _ := CheckInstalled(ctx, v.Id)
					v.Installed = installed
					apps = append(apps, v)
				}
			}
			break
		}
	}
	return apps, nil
}

// 卸载应用
// ```pm uninstall --user 0 com.huawei.fastapp```
// TODO 这里应该判断卸载了所有应用
func UninstallApp(ctx context.Context, packageName string, relatedIds []string) (bool, error) {
	result, err := AdbShellCommand(ctx, "pm", "uninstall", "--user", "0", packageName)
	if err != nil {
		utils.LogE(ctx, err.Error())
	}
	for _, v := range relatedIds {
		_, err = AdbShellCommand(ctx, "pm", "uninstall", "--user", "0", v)
		if err != nil {
			utils.LogE(ctx, err.Error())
		}
	}
	fmt.Printf("UninstallApp:%s\n", result)
	return result == "Success", err
}

// 装回应用
// ```pm install-existing --user 0 com.huawei.fastapp```
func InstallExistingApp(ctx context.Context, packageName string, relatedIds []string) (bool, error) {
	result, err := AdbShellCommand(ctx, "pm", "install-existing", "--user", "0", packageName)
	for _, v := range relatedIds {
		_, err = AdbShellCommand(ctx, "pm", "install-existing", "--user", "0", v)
	}
	utils.Log(ctx, result)
	return strings.Contains(result, "installed"), err
}

// 禁用应用
// ```pm disable-user com.huawei.android.hwouc```
func DisableApp(ctx context.Context, packageName string) (bool, error) {
	result, err := AdbShellCommand(ctx, "pm", "disable-user", packageName)
	fmt.Printf("DisableApp:%s\n", result)
	utils.Log(ctx, result)
	return strings.Contains(result, "disabled"), err
}

// 启用应用
// ```pm enable com.huawei.android.hwouc```
// com.huawei.android.hwouc/com.huawei.android.hwouc.ui.activities.MainEntranceActivity
func EnableApp(ctx context.Context, packageName string) (bool, error) {
	result, err := AdbShellCommand(ctx, "pm", "enable", packageName)
	utils.Log(ctx, result)
	//启动更新进程，以确保服务运行起来，这样就可以更新前端状态
	StartApp(ctx, "com.huawei.android.hwouc", "com.huawei.android.hwouc.ui.activities.MainEntranceActivity")
	return strings.Contains(result, "enabled"), err
}

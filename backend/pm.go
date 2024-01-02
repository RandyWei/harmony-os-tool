package backend

import (
	"fmt"

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
	return result != "", nil
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

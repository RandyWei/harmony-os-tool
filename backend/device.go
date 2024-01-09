package backend

import (
	"context"
	"errors"
	"os/exec"
	"strings"

	"icu.bughub.app/harmonyos-tool/backend/models"
	"icu.bughub.app/harmonyos-tool/backend/utils"
)

/**
 * 等待设备连接
 */
func WaitForDevice(ctx context.Context, appDir string) ([]models.Device, error) {

	devices := []models.Device{}
	cmd := exec.Command("adb", "devices", "-l")
	PrepareBackgroundCommand(cmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		utils.LogE(ctx, err.Error())
		e := errors.New("发生了错误")
		if errors.Is(err, exec.ErrNotFound) {
			//没有安装adb
			e = errors.New("程序依赖adb服务，需要安装adb方可使用")
		}
		return devices, e
	}

	execResult := string(output)
	//cannot connect to daemon
	//adb: no devices/emulators found
	utils.Log(ctx, execResult)

	if strings.Contains(execResult, "List of devices attached") {
		//按行分割
		for index, line := range strings.Split(execResult, "\n") {
			if index == 0 {
				continue
			}
			if strings.Contains(line, "device") {
				deviceInfos := strings.Split(line, " ")
				device := models.Device{}
				for index, info := range deviceInfos {
					//获取设备名称
					if index == 0 {
						device.Id = info
						prop, _ := GetProp(ctx, info)
						if strings.Contains(prop, "huawei") {
							device.Brand = "HuaWei"
						} else if strings.Contains(prop, "honor") {
							device.Brand = "Honor"
						}
					}

					if strings.Contains(info, "product") {
						device.Product = strings.Split(info, ":")[1]
					}
					if strings.Contains(info, "model") {
						device.Model = strings.Split(info, ":")[1]
					}
				}

				devices = append(devices, device)
			}
		}
		return devices, nil
	}

	return devices, nil
}

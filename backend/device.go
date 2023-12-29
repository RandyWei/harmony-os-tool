package backend

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"icu.bughub.app/harmonyos-tool/backend/models"
)

/**
 * 等待设备连接
 */
func WaitForDevice(appDir string) ([]models.Device, error) {

	devices := []models.Device{}
	cmd := exec.Command("adb", "devices", "-l")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		e := errors.New("发生了错误")
		if errors.Is(err, exec.ErrNotFound) {
			//没有安装adb
			e = errors.New("程序依赖adb服务，需要安装adb方可使用")
		}
		return devices, e
	}

	execResult := string(output)
	//cannot connect to daemon
	fmt.Println(execResult)

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

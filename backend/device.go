package backend

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

/**
 * 等待设备连接
 */
func WaitForDevice(appDir string) ([]string, error) {

	devices := []string{}
	cmd := exec.Command("adb", "devices")
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
				//获取设备名称
				deviceName := strings.Split(line, "\t")[0]
				fmt.Println(deviceName)
				devices = append(devices, deviceName)
			}
		}
		return devices, nil
	}

	return devices, nil
}

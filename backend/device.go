package backend

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/**
 * 等待设备连接
 */
func WaitForDevice(appDir string) (bool, error) {

	cmd := exec.Command("adb", "kill-server", "&&", "adb", "wait-for-device")
	cmd.Dir = appDir + string(os.PathSeparator) + "platform-tools"
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		e := errors.New("发生了错误")
		if errors.Is(err, exec.ErrNotFound) {
			//没有安装adb
			e = errors.New("程序依赖adb服务，需要安装adb方可使用")
		}
		return false, e
	}

	execResult := string(output)

	fmt.Println(execResult)

	return strings.Contains(execResult, "Android Debug Bridge"), nil
}

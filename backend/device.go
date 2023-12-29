package backend

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func WaitForDevice() (bool, error) {
	cmd := exec.Command("adb1", "kill-server", "&&", "adb", "wait-for-device")

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

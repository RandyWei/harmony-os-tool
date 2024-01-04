package backend

import (
	"context"
	"errors"
	"fmt"
	"os/exec"

	"icu.bughub.app/harmonyos-tool/backend/utils"
)

func AdbShellCommand(ctx context.Context, commands ...string) (string, error) {
	commands = append([]string{"shell"}, commands...)
	cmd := exec.Command("adb", commands...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		utils.LogE(ctx, err.Error())
		e := errors.New("发生了错误")
		if errors.Is(err, exec.ErrNotFound) {
			//没有安装adb
			e = errors.New("程序依赖adb服务，需要安装adb方可使用")
		}
		return "", e
	}

	return string(output), nil
}

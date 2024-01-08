//go:build windows
// +build windows

package backend

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"syscall"

	"icu.bughub.app/harmonyos-tool/backend/utils"
)

func AdbShellCommand(ctx context.Context, commands ...string) (string, error) {
	commands = append([]string{"shell"}, commands...)
	cmd := exec.Command("adb", commands...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000,
	}
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
	result := string(output)
	utils.Log(ctx, result)
	return result, nil
}

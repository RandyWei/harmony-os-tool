//go:build windows
// +build windows

package backend

import (
	"os/exec"
	"syscall"
)

// Windows环境下设置后台隐藏
func PrepareBackgroundCommand(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000,
	}
}

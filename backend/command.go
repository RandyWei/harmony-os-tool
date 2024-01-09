//go:build !windows
// +build !windows

package backend

import "os/exec"

// Windows环境下设置后台隐藏
func PrepareBackgroundCommand(cmd *exec.Cmd) {

}

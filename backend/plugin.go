package backend

import (
	"fmt"
	"os/exec"
	"strings"
)

func TryRunAdb() bool {
	cmd := exec.Command("adb", "version")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	execResult := string(output)

	fmt.Println(execResult)

	return strings.Contains(execResult, "Android Debug Bridge")
}

package backend

// 查看当前设备内存情况
// ```free -h```
func Free() (string, error) {
	return AdbShellCommand("free", "-h")
}

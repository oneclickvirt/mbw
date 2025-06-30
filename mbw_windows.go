//go:build windows && !(amd64 || 386 || arm64)
// +build windows,!amd64,!386,!arm64

package mbw

import (
	"fmt"
	"os"
	"os/exec"
)

// GetMBW 获取与当前系统匹配的 mbw 二进制文件并返回路径
func GetMBW() (string, string, error) {
	// binaryName := "mbw-windows-amd64.exe"
	// 检查系统是否有原生 mbw 命令
	if _, err := exec.LookPath("mbw"); err == nil {
		return "mbw", "", nil // 返回系统原生命令
	}
	return "", "", fmt.Errorf("Can not use mbw")
}

// ExecuteMBW 执行 mbw 命令
func ExecuteMBW(mbwPath string, args []string) error {
	var cmd *exec.Cmd
	if mbwPath == "mbw" {
		// 使用系统 mbw
		cmd = exec.Command(mbwPath, args...)
	} else {
		// 在 Windows 上直接执行并传递 mbw 作为第一个参数
		cmd = exec.Command(mbwPath, append([]string{"mbw"}, args...)...)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

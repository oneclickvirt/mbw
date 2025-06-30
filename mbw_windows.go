//go:build windows && !(amd64 || 386 || arm64)
// +build windows,!amd64,!386,!arm64

package mbw

import (
	"fmt"
	"os"
	"os/exec"
)

// GetDD 获取与当前系统匹配的 dd 二进制文件并返回路径
func GetDD() (string, string, error) {
	// binaryName := "mbw-windows-amd64.exe"
	// 检查系统是否有原生 dd 命令
	if _, err := exec.LookPath("mbw"); err == nil {
		return "mbw", "", nil // 返回系统原生命令
	}
	return "", "", fmt.Errorf("Can not use dd")
}

// ExecuteDD 执行 dd 命令
func ExecuteDD(ddPath string, args []string) error {
	var cmd *exec.Cmd
	if ddPath == "mbw" {
		// 使用系统 dd
		cmd = exec.Command(ddPath, args...)
	} else {
		// 在 Windows 上直接执行并传递 dd 作为第一个参数
		cmd = exec.Command(ddPath, append([]string{"mbw"}, args...)...)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

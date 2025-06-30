//go:build darwin && !(amd64 || arm64)
// +build darwin,!amd64,!arm64

package mbw

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// GetMBW 获取与当前系统匹配的 mbw 二进制文件并返回路径
func GetMBW() (string, string, error) {
	binaryName := "mbw-darwin-arm64"
	// 优先尝试 sudo mbw 是否可用
	if path, err := exec.LookPath("mbw"); err == nil {
		testCmd := exec.Command("sudo", path, "-h")
		if err := testCmd.Run(); err == nil {
			return "sudo mbw", "", nil
		}
		// 如果 sudo mbw 不可用，则尝试直接使用 mbw
		testCmd = exec.Command(path, "-h")
		if err := testCmd.Run(); err == nil {
			return "mbw", "", nil
		}
	}
	// 创建临时目录存放二进制文件
	tempDir, err := os.MkdirTemp("", "mbwwrapper")
	if err != nil {
		return "", "", fmt.Errorf("创建临时目录失败: %v", err)
	}
	return "", "", fmt.Errorf("无法找到可用的 mbw 命令")
}

// ExecuteMBW 执行 mbw 命令
func ExecuteMBW(mbwPath string, args []string) error {
	var cmd *exec.Cmd
	if mbwPath == "mbw" {
		// 使用系统 mbw
		cmd = exec.Command(mbwPath, args...)
	} else {
		// 使用提取的 mbw mbw
		mbwCmd := fmt.Sprintf("%s mbw %s", mbwPath, strings.Join(args, " "))
		cmd = exec.Command("sh", "-c", mbwCmd)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

//go:build freebsd
// +build freebsd

package mbw

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// GetMBW 返回适用于当前系统的 mbw 命令字符串（带不带 sudo）和临时文件路径（用于清理）
func GetMBW() (mbwCmd string, tempFile string, err error) {
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

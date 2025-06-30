//go:build linux && arm64
// +build linux,arm64

package mbw

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//go:embed bin/mbw-linux-arm64
var binFiles embed.FS

// GetMBW 返回适用于当前系统的 mbw 命令字符串（带不带 sudo）和临时文件路径（用于清理）
func GetMBW() (mbwCmd string, tempFile string, err error) {
	var errors []string
	// 1. 尝试系统自带 mbw
	if path, lookErr := exec.LookPath("mbw"); lookErr == nil {
		// 确保 testCmd 被初始化
		testCmd := exec.Command("sudo", path, "-h")
		if runErr := testCmd.Run(); runErr == nil {
			return "sudo mbw", "", nil
		} else {
			errors = append(errors, fmt.Sprintf("sudo mbw 测试失败: %v", runErr))
		}
		// 直接尝试 mbw
		testCmd = exec.Command(path, "-h")
		if runErr := testCmd.Run(); runErr == nil {
			return "mbw", "", nil
		} else {
			errors = append(errors, fmt.Sprintf("mbw 直接运行失败: %v", runErr))
		}
	} else {
		errors = append(errors, fmt.Sprintf("无法找到 mbw: %v", lookErr))
	}
	// 2. 创建临时目录
	tempDir, tempErr := os.MkdirTemp("", "mbwwrapper")
	if tempErr != nil {
		return "", "", fmt.Errorf("创建临时目录失败: %v", tempErr)
	}
	// 3. 尝试使用 glibc 版本 mbw
	binName := "mbw-linux-arm64"
	binPath := filepath.Join("bin", binName)
	fileContent, readErr := binFiles.ReadFile(binPath)
	if readErr == nil {
		tempFile = filepath.Join(tempDir, binName)
		writeErr := os.WriteFile(tempFile, fileContent, 0755)
		if writeErr == nil {
			// 确保 testCmd 被初始化
			testCmd := exec.Command("sudo", tempFile, "-h")
			if runErr := testCmd.Run(); runErr == nil {
				return fmt.Sprintf("sudo %s", tempFile), tempFile, nil
			} else {
				errors = append(errors, fmt.Sprintf("sudo %s 运行失败: %v", tempFile, runErr))
			}
			// 直接尝试
			testCmd = exec.Command(tempFile, "-h")
			if runErr := testCmd.Run(); runErr == nil {
				return tempFile, tempFile, nil
			} else {
				errors = append(errors, fmt.Sprintf("%s 运行失败: %v", tempFile, runErr))
			}
		} else {
			errors = append(errors, fmt.Sprintf("写入临时文件失败 (%s): %v", tempFile, writeErr))
		}
	} else {
		errors = append(errors, fmt.Sprintf("读取嵌入的 mbw glibc 版本失败: %v", readErr))
	}
	// 返回所有错误信息
	return "", "", fmt.Errorf("无法找到可用的 mbw 命令:\n%s", strings.Join(errors, "\n"))
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

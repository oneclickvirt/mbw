//go:build windows && 386
// +build windows,386

package mbw

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

//go:embed bin/mbw-windows-386.exe
var binFiles embed.FS

// GetMBW 获取与当前系统匹配的 mbw 二进制文件并返回路径
func GetMBW() (string, string, error) {
	binaryName := "mbw-windows-386.exe"
	// 检查系统是否有原生 mbw 命令
	if _, err := exec.LookPath("mbw"); err == nil {
		return "mbw", "", nil // 返回系统原生命令
	}
	// 创建临时目录存放二进制文件
	tempDir, err := os.MkdirTemp("", "mbwwrapper")
	if err != nil {
		return "", "", fmt.Errorf("创建临时目录失败: %v", err)
	}
	// 读取嵌入的二进制文件
	binPath := filepath.Join("bin", binaryName)
	fileContent, err := binFiles.ReadFile(binPath)
	if err != nil {
		return "", "", fmt.Errorf("读取嵌入的 mbw 二进制文件失败: %v", err)
	}
	// 写入临时文件
	tempFile := filepath.Join(tempDir, binaryName)
	if err := os.WriteFile(tempFile, fileContent, 0755); err != nil {
		return "", "", fmt.Errorf("写入临时文件失败: %v", err)
	}
	return tempFile, tempFile, nil
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

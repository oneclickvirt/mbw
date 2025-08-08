//go:build windows && arm64
// +build windows,arm64

package mbw

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//go:embed bin/mbw-windows-amd64.exe
var binFiles embed.FS

func GetMBW() (string, string, error) {
	binaryName := "mbw-windows-arm64.exe"
	if path, err := exec.LookPath("mbw"); err == nil {
		cmd := exec.Command(path, "-h")
		output, runErr := cmd.CombinedOutput()
		if runErr == nil && strings.Contains(string(output), "Usage: mbw") {
			return "mbw", "", nil
		}
	}
	tempDir, err := os.MkdirTemp("", "mbwwrapper")
	if err != nil {
		return "", "", fmt.Errorf("创建临时目录失败: %v", err)
	}
	binPath := fmt.Sprintf("bin/%s", binaryName)
	fileContent, err := binFiles.ReadFile(binPath)
	if err != nil {
		return "", "", fmt.Errorf("读取嵌入的 mbw 二进制文件失败: %v", err)
	}
	tempFile := filepath.Join(tempDir, binaryName)
	if err := os.WriteFile(tempFile, fileContent, 0755); err != nil {
		return "", "", fmt.Errorf("写入临时文件失败: %v", err)
	}
	return tempFile, tempFile, nil
}

func ExecuteMBW(mbwPath string, args []string) error {
	var cmd *exec.Cmd
	if mbwPath == "mbw" {
		cmd = exec.Command(mbwPath, args...)
	} else {
		cmd = exec.Command(mbwPath, args...)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

//go:build windows && amd64
// +build windows,amd64

package mbw

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

//go:embed bin/mbw-windows-amd64.exe
var binFiles embed.FS

func GetMBW() (string, string, error) {
	binaryName := "mbw-windows-amd64.exe"
	if _, err := exec.LookPath("mbw"); err == nil {
		return "mbw", "", nil
	}
	tempDir, err := os.MkdirTemp("", "mbwwrapper")
	if err != nil {
		return "", "", fmt.Errorf("创建临时目录失败: %v", err)
	}
	binPath := filepath.Join("bin", binaryName)
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
		cmd = exec.Command(mbwPath, append([]string{"mbw"}, args...)...)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

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

func GetMBW() (string, string, error) {
	binaryName := "mbw-windows-386.exe"
	if path, err := exec.LookPath("mbw"); err == nil {
		cmd := exec.Command(path, "-h")
		output, runErr := cmd.CombinedOutput()
		if runErr == nil && (len(output) > 0 && (string(output) != "")) {
			return "mbw", "", nil
		}
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

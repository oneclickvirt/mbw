//go:build darwin && amd64
// +build darwin,amd64

package mbw

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//go:embed bin/mbw-darwin-amd64
var binFiles embed.FS

func GetMBW() (mbwCmd string, tempFile string, err error) {
	var errors []string
	if path, lookErr := exec.LookPath("mbw"); lookErr == nil {
		output, runErr := exec.Command(path, "-h").CombinedOutput()
		if strings.Contains(string(output), "Usage: mbw") {
			return "mbw", "", nil
		} else {
			errors = append(errors, fmt.Sprintf("mbw 直接运行失败: %v\n输出: %s", runErr, string(output)))
		}
	} else {
		errors = append(errors, fmt.Sprintf("无法找到 mbw: %v", lookErr))
	}
	tempDir, tempErr := os.MkdirTemp("", "mbwwrapper")
	if tempErr != nil {
		return "", "", fmt.Errorf("创建临时目录失败: %v", tempErr)
	}
	binName := "mbw-darwin-amd64"
	binPath := filepath.Join("bin", binName)
	fileContent, readErr := binFiles.ReadFile(binPath)
	if readErr == nil {
		tempFile = filepath.Join(tempDir, binName)
		writeErr := os.WriteFile(tempFile, fileContent, 0755)
		if writeErr == nil {
			output, runErr := exec.Command(tempFile, "-h").CombinedOutput()
			if strings.Contains(string(output), "Usage: mbw") {
				return tempFile, tempFile, nil
			} else {
				errors = append(errors, fmt.Sprintf("%s 运行失败: %v\n输出: %s", tempFile, runErr, string(output)))
			}
		} else {
			errors = append(errors, fmt.Sprintf("写入临时文件失败 (%s): %v", tempFile, writeErr))
		}
	} else {
		errors = append(errors, fmt.Sprintf("读取嵌入的 mbw 版本失败: %v", readErr))
	}
	return "", "", fmt.Errorf("无法找到可用的 mbw 命令:\n%s", strings.Join(errors, "\n"))
}

func ExecuteMBW(mbwPath string, args []string) error {
	var cmd *exec.Cmd
	if mbwPath == "mbw" {
		cmd = exec.Command(mbwPath, args...)
	} else {
		cmd = exec.Command("sh", "-c", fmt.Sprintf("%s %s", mbwPath, strings.Join(args, " ")))
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

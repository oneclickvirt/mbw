//go:build darwin && !(amd64 || arm64)
// +build darwin,!amd64,!arm64

package mbw

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetMBW() (string, string, error) {
	if path, err := exec.LookPath("mbw"); err == nil {
		output, runErr := exec.Command("sudo", path, "-h").CombinedOutput()
		if strings.Contains(string(output), "Usage: mbw") {
			return "sudo mbw", "", nil
		} else {
			output, runErr = exec.Command(path, "-h").CombinedOutput()
			if strings.Contains(string(output), "Usage: mbw") {
				return "mbw", "", nil
			}
		}
	}
	tempDir, err := os.MkdirTemp("", "mbwwrapper")
	if err != nil {
		return "", "", fmt.Errorf("创建临时目录失败: %v", err)
	}
	return "", "", fmt.Errorf("无法找到可用的 mbw 命令")
}

func ExecuteMBW(mbwPath string, args []string) error {
	var cmd *exec.Cmd
	if mbwPath == "mbw" || mbwPath == "sudo mbw" {
		cmd = exec.Command("sh", "-c", fmt.Sprintf("%s %s", mbwPath, strings.Join(args, " ")))
	} else {
		cmd = exec.Command("sh", "-c", fmt.Sprintf("%s %s", mbwPath, strings.Join(args, " ")))
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

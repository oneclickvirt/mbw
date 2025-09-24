//go:build !(linux || freebsd || darwin || openbsd)
// +build !linux,!freebsd,!darwin,!openbsd

package mbw

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetMBW() (mbwCmd string, tempFile string, err error) {
	if path, err := exec.LookPath("mbw"); err == nil {
		output, runErr := exec.Command(path, "-h").CombinedOutput()
		if runErr == nil && strings.Contains(string(output), "Usage: mbw") {
			return "mbw", "", nil
		}
	}
	return "", "", fmt.Errorf("无法找到可用的 mbw 命令")
}

func ExecuteMBW(mbwPath string, args []string) error {
	var cmd *exec.Cmd
	if mbwPath == "mbw" {
		cmd = exec.Command(mbwPath, args...)
	} else {
		fullCmd := fmt.Sprintf("%s %s", mbwPath, strings.Join(args, " "))
		cmd = exec.Command("sh", "-c", fullCmd)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

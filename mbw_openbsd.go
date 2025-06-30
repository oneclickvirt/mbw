//go:build openbsd
// +build openbsd

package mbw

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetMBW() (mbwCmd string, tempFile string, err error) {
	if path, err := exec.LookPath("mbw"); err == nil {
		output, _ := exec.Command(path, "-h").CombinedOutput()
		if strings.Contains(string(output), "Usage: mbw") {
			return "mbw", "", nil
		}
	}
	return "", "", fmt.Errorf("无法找到可用的 mbw 命令")
}

func ExecuteMBW(mbwPath string, args []string) error {
	var cmd *exec.Cmd
	if mbwPath == "mbw" {
		cmd = exec.Command("sh", "-c", fmt.Sprintf("%s %s", mbwPath, strings.Join(args, " ")))
	} else {
		cmd = exec.Command("sh", "-c", fmt.Sprintf("%s %s", mbwPath, strings.Join(args, " ")))
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

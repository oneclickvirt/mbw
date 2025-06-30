//go:build windows && !(amd64 || 386 || arm64)
// +build windows,!amd64,!386,!arm64

package mbw

import (
	"fmt"
	"os"
	"os/exec"
)

func GetMBW() (string, string, error) {
	if _, err := exec.LookPath("mbw"); err == nil {
		return "mbw", "", nil
	}
	return "", "", fmt.Errorf("无法使用 mbw")
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

package mbw

import (
	"os"
	"testing"
)

// TestGetMBW 测试 GetMBW 函数是否能够正确返回可用的 mbw 命令
func TestGetMBW(t *testing.T) {
	mbwCmd, tempFile, err := GetMBW()
	if err != nil {
		t.Fatalf("GetMBW 失败: %v", err)
	}
	t.Logf("返回的命令: %s", mbwCmd)
	if tempFile != "" {
		// 检查文件是否存在
		if _, statErr := os.Stat(tempFile); os.IsNotExist(statErr) {
			t.Errorf("临时文件 %s 不存在", tempFile)
		}
	}
}

// TestExecuteMBW 测试 ExecuteMBW 函数是否能够执行 mbw -h（帮助信息）
func TestExecuteMBW(t *testing.T) {
	mbwCmd, tempFile, err := GetMBW()
	if err != nil {
		t.Skipf("跳过测试 ExecuteMBW，因为无法获取 mbw 命令: %v", err)
		return
	}
	defer func() {
		if tempFile != "" {
			os.Remove(tempFile)
		}
	}()

	args := []string{"-h"} // 安全测试
	err = ExecuteMBW(mbwCmd, args)
	if err != nil {
		t.Errorf("ExecuteMBW 执行失败: %v", err)
	}
}

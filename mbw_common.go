package mbw

import (
	"os"
	"path/filepath"
)

// CleanMBW 删除临时提取出的 mbw/mbw 文件
func CleanMBW(tempFile string) error {
	if tempFile == "" {
		return nil // 不需要清理
	}
	// 删除整个临时目录
	return os.RemoveAll(filepath.Dir(tempFile))
}

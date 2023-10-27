package common

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetRootProject(path string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("[common/path/getRootProject()]Lỗi khi lấy thư mục làm việc:", err)
		return "", err
	}
	fmt.Println("[common/path/getRootProject()]: " + dir)
	return filepath.Join(dir, path), nil
}

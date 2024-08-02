package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetLeetcodeDir カレントディレクトリから leetcode/ のパスを特定し返す
func GetLeetcodeDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	parts := strings.Split(dir, string(filepath.Separator))
	for i := range parts {
		if parts[i] == "leetcode" {
			p := filepath.Join(parts[:i+1]...)
			return filepath.Join("/", p), nil
		}
	}
	return "", fmt.Errorf("leetcode directory not found in path")
}

package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	LeetcodeDirPath       string `json:"leetcodeDirPath"`
	ProblemDirPath        string `json:"problemDirPath"`
	VscodeLeetcodeVersion string `json:"vscodeLeetcodeVersion"`
}

// CreateLeetcodeDirAndConfig leetcode/ とそのなかに config.json を生成
func CreateLeetcodeDirAndConfig(leetcodeDirPath string, vscodeLeetcodeVersion string) error {
	err := os.Mkdir(leetcodeDirPath, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	configFilePath := filepath.Join(leetcodeDirPath, "config.json")
	config := Config{
		LeetcodeDirPath:       leetcodeDirPath,
		VscodeLeetcodeVersion: vscodeLeetcodeVersion,
	}

	file, err := os.Create(configFilePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Printf("Failed to close file: %v\n", cerr)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(config)
	if err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}

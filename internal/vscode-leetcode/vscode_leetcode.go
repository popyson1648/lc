package vscode_leetcode

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

// Embed the files
//
//go:embed vscode-leetcode-files/package.json
var packageJSON []byte

// getCurrentUsername 現在ログインしているユーザ名を返す
func getCurrentUsername() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.Username, nil
}

// getVscodeLeetCodeDir vscode-leetcode のディレクトリのパスを返す
func getVscodeLeetCodeDir(leetcodeDirPath string) (string, error) {
	configFilePath := filepath.Join(leetcodeDirPath, "config.json")
	configData, err := os.ReadFile(configFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read leetcode/config.json: %w", err)
	}

	var config map[string]interface{}
	err = json.Unmarshal(configData, &config)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal leetcode/config.json: %w", err)
	}

	version, ok := config["vscodeLeetcodeVersion"].(string)
	if !ok || version == "" {
		return "", fmt.Errorf("vscodeLeetcodeVersion not found in config.json")
	}

	username, err := getCurrentUsername()
	if err != nil {
		return "", err
	}

	vscodeLeetCodeDir := filepath.Join("/Users", username, ".vscode", "extensions", fmt.Sprintf("leetcode.vscode-leetcode-%s", version))
	return vscodeLeetCodeDir, nil
}

// OverWritePackageJson vscode-leetcode ディレクトリの package.json を上書き
func OverWritePackageJson(leetcodeDirPath string) error {
	vscodeLeetCodeDirPath, err := getVscodeLeetCodeDir(leetcodeDirPath)
	if err != nil {
		return err
	}

	filePath := filepath.Join(vscodeLeetCodeDirPath, "package.json")
	err = os.WriteFile(filePath, packageJSON, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	fmt.Printf("written package.json to %s\n", filePath)
	return nil
}

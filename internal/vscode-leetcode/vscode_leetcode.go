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

//go:embed vscode-leetcode-files/problemConfig.js
var problemConfigJS []byte

//go:embed vscode-leetcode-files/show.js
var showJS []byte

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

// WriteFiles vscode-leetcode が lc と連携するためのファイル群を vscode-leetcode ディレクトリに生成
func WriteFiles(leetcodeDirPath string) error {
	vscodeLeetCodeDirPath, err := getVscodeLeetCodeDir(leetcodeDirPath)
	if err != nil {
		return err
	}

	files := map[string][]byte{
		"package.json":     packageJSON,
		"problemConfig.js": problemConfigJS,
		"show.js":          showJS,
	}

	for name, data := range files {
		filePath := filepath.Join(vscodeLeetCodeDirPath, name)
		err := os.WriteFile(filePath, data, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to write file %s: %w", filePath, err)
		}
		fmt.Printf("written %s to %s\n", name, filePath)
	}

	// config-path.json を生成
	configPathData := map[string]string{
		"leetcodeDirPath": leetcodeDirPath,
	}

	configPathJSON, err := json.MarshalIndent(configPathData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to generate config-path.json: %w", err)
	}

	configPathFile := filepath.Join(vscodeLeetCodeDirPath, "config-path.json")
	err = os.WriteFile(configPathFile, configPathJSON, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write config-path.json: %w", err)
	}
	fmt.Printf("written config-path.json to %s\n", configPathFile)

	return nil
}

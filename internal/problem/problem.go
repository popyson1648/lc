package problem

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetBranchName 現在のブランチ名を取得
func GetBranchName() (string, error) {
	ecmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := ecmd.Output()
	if err != nil {
		return "", fmt.Errorf("could not determine branch name: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}

// ExtractProblemDirName ブランチ名から､問題ディレクトリ名を抽出して返す
func ExtractProblemDirName(branchName string) (string, string, error) {
	parts := strings.Split(branchName, "-")
	if len(parts) < 3 {
		return "", "", fmt.Errorf("invalid branch name format")
	}
	workbook := parts[0]
	problemDir := strings.Join(parts[1:], "-")
	return workbook, problemDir, nil
}

// CreateDir leetcode/ 配下に problems/ワークブック名/問題名 のディレクトリを生成する
func CreateDir(leetcodeDirPath string) error {
	branchName, err := GetBranchName()
	if err != nil {
		return err
	}

	workbook, problemDirName, err := ExtractProblemDirName(branchName)
	if err != nil {
		return err
	}

	err = os.Chdir(leetcodeDirPath)
	if err != nil {
		return fmt.Errorf("failde to change directory path: %w", err)
	}

	err = os.MkdirAll(filepath.Join("problems", workbook, problemDirName), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failde to create directory: %w", err)
	}

	fmt.Println("created problem directory:", filepath.Join("problems", workbook, problemDirName))
	return nil
}

// SetDirPath leetcode/config.json に問題ディレクトリのパスを設定する
func SetDirPath(leetcodeDirPath string) error {

	branchName, err := GetBranchName()
	if err != nil {
		return err
	}

	workbook, problemDirName, err := ExtractProblemDirName(branchName)
	if err != nil {
		return err
	}

	err = os.Chdir(leetcodeDirPath)
	if err != nil {
		return fmt.Errorf("failde to change directory path: %w", err)
	}

	configData, err := os.ReadFile("config.json")
	if err != nil {
		return fmt.Errorf("failed to read leetcode/config.json: %w", err)
	}

	var config map[string]interface{}
	err = json.Unmarshal(configData, &config)
	if err != nil {
		return fmt.Errorf("failed to unmarshal leetcode/config.json: %w", err)
	}

	problemDirPath := filepath.Join(leetcodeDirPath, "problems", workbook, problemDirName)
	config["problemDirPath"] = problemDirPath

	newConfigData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal leetcode/config.json: %w", err)
	}

	err = os.WriteFile("config.json", newConfigData, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write leetcode/config.json: %w", err)
	}

	fmt.Println("updated problemDirPath in leetcode/config.json to:", problemDirPath)
	return nil
}

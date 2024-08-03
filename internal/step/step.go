package step

import (
	"encoding/json"
	"errors"
	"fmt"
	"lc/internal/config"
	"os"
	"path/filepath"
	"strconv"
)

// UpdateStep step_count.json をもとに step_x.py を step_<number>.py にリネーム
func UpdateStep() error {
	cur, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	configPath := filepath.Join(cur, "step_count.json")

	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read step_count.json: %w", err)
	}

	var stepCount config.StepCount
	err = json.Unmarshal(configFile, &stepCount)
	if err != nil {
		return fmt.Errorf("failed to parse step_count.json: %w", err)
	}

	oldStepFilePath := filepath.Join(cur, "step_x.py")
	newStepFilePath := filepath.Join(cur, "step_"+strconv.Itoa(stepCount.StepNumber)+".py")

	err = os.Rename(oldStepFilePath, newStepFilePath)
	if err != nil {
		return fmt.Errorf("failed to rename file: %w", err)
	}

	stepCount.StepNumber++

	newConfigFile, err := json.MarshalIndent(stepCount, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to generate new step_count.json: %w", err)
	}

	err = os.WriteFile(configPath, newConfigFile, 0644)
	if err != nil {
		return fmt.Errorf("failed to save new step_count.json: %w", err)
	}

	fmt.Println("step updated successfully.")
	return nil
}

func Move(leetcodeDirPath string, problemDirPath string) error {
	tmpDirPath := filepath.Join(leetcodeDirPath, "tmp")

	// カレントディレクトリに step_x.py があるかチェック
	if !Exists(tmpDirPath) {
		return errors.New("step_x.py does not exist")
	}

	// step_x.py を 問題ディレクトリに移動
	err := os.Rename(filepath.Join(tmpDirPath, "step_x.py"), filepath.Join(problemDirPath, "step_x.py"))
	if err != nil {
		return fmt.Errorf("failed to rename step_x.py: %w", err)
	}

	return nil
}

func Exists(path string) bool {
	fileFound := false

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("error reading directory: %w", err)
	}

	targetFile := "step_x.py"
	for _, file := range files {
		if file.Name() == targetFile {
			fileFound = true
			break
		}
	}

	return fileFound
}

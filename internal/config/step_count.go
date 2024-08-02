package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type StepCount struct {
	StepNumber int `json:"stepNumber"`
}

// CreateStepCountJson 問題ディレクトリの中に step-count.json を生成する
func CreateStepCountJson(leetcodeDirPath string) error {
	configFilePath := filepath.Join(leetcodeDirPath, "config.json")

	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return fmt.Errorf("failed to read leetcode/config.json: %w", err)
	}

	var config map[string]interface{}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return fmt.Errorf("failed to parse leetcode/config.json: %w", err)
	}

	problemDirPath, ok := config["problemDirPath"].(string)
	if !ok || problemDirPath == "" {
		return errors.New("'problemDirPath' not found in config.json")
	}

	stepCountFilePath := filepath.Join(problemDirPath, "step_count.json")

	stepCount := StepCount{StepNumber: 1}

	jsonData, err := json.MarshalIndent(stepCount, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to generate step_count.json: %w", err)
	}

	err = os.WriteFile(stepCountFilePath, jsonData, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write step_count.json: %w", err)
	}

	fmt.Printf("step_count.json has been generated at %s\n", stepCountFilePath)
	return nil
}

package markdown

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// analysisSourceCode content を解析し､コメントとコードを文字列で返す
func analysisSourceCode(content string) (string, string) {
	var comments strings.Builder
	var code strings.Builder

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "#") {
			if !strings.HasPrefix(trimmedLine, "# @lc") {
				comments.WriteString(strings.TrimSpace(trimmedLine[1:]) + "\n")
			}
		} else {
			code.WriteString(line + "\n")
		}
	}

	return comments.String(), code.String()
}

// GenerateProblemMD step_<number>.py から problem.md を生成する
func GenerateProblemMD(problemDir string) error {
	outputFile := filepath.Join(problemDir, "problem.md")

	files, err := os.ReadDir(problemDir)
	if err != nil {
		return fmt.Errorf("failed to read problem directory: %w", err)
	}

	var content strings.Builder
	var stepFiles []string

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "step_") && strings.HasSuffix(file.Name(), ".py") && file.Name() != "step_x.py" {
			stepFiles = append(stepFiles, file.Name())
		}
	}

	sort.Strings(stepFiles)

	for _, stepFile := range stepFiles {
		stepPath := filepath.Join(problemDir, stepFile)
		stepContent, err := os.ReadFile(stepPath)
		if err != nil {
			return fmt.Errorf("failed to read step-count.json %s: %w", stepPath, err)
		}

		comments, code := analysisSourceCode(string(stepContent))

		// Extract step number including sub-steps
		stepNumber := strings.TrimSuffix(strings.TrimPrefix(stepFile, "step_"), ".py")
		stepNumber = strings.ReplaceAll(stepNumber, "_", "-")

		content.WriteString(fmt.Sprintf("\n--------------------\nSTEP %s\n--------------------\n\n", stepNumber))
		content.WriteString(comments)
		content.WriteString("\n```py\n")
		content.WriteString(code)
		content.WriteString("\n```\n\n")
	}

	if err = os.WriteFile(outputFile, []byte(content.String()), 0644); err != nil {
		return fmt.Errorf("failed to write text to %s: %w", outputFile, err)
	}

	fmt.Printf("%s has been generated.\n", outputFile)
	return nil
}

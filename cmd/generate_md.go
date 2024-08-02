package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"lc/internal/markdown" // モジュールのインポート
	"os"
	"regexp"
)

var generateMDCmd = &cobra.Command{
	Use:   "generate-md",
	Short: "Generate a problem.md file from the step_*.py files in the problem directory",
	Run: func(cmd *cobra.Command, args []string) {

		cur, err := os.Getwd()
		if err != nil {
			fmt.Errorf("failed to getting current directory: %w", err)
			return
		}

		files, err := os.ReadDir(cur)
		if err != nil {
			fmt.Println("failed to reading directory:", err)
			return
		}

		// step_<任意の数字>.py が存在するかチェック
		pattern := `^step_\d+\.py$`
		regex := regexp.MustCompile(pattern)
		fileFound := false
		for _, file := range files {
			if regex.MatchString(file.Name()) {
				fileFound = true
				break
			}
		}

		if !fileFound {
			fmt.Println("error: step files not found")
			return
		}

		err = markdown.GenerateProblemMD(cur)
		if err != nil {
			fmt.Println("failed to generating problem.md:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(generateMDCmd)
}

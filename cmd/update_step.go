package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"lc/internal/config"
	"lc/internal/dir"
	"lc/internal/problem"
	"lc/internal/step"
	"os"
	"path/filepath"
)

var updateStepCmd = &cobra.Command{
	Use:   "update-step",
	Short: "Update the step number and rename the step_x.py file",
	Long:  `Refer to the value of "stepNumber" in step_number.json, rename the step_x.py file accordingly, and then update the value of "stepNumber`,
	Run: func(cmd *cobra.Command, args []string) {

		var branchName, workbook, problemDir string
		var cur string
		var err error

		branchName, err = problem.GetBranchName()
		if err != nil {
			fmt.Println("error while getting branch name:", err)
			return
		}

		workbook, problemDir, err = problem.ExtractProblemDirName(branchName)
		if err != nil {
			fmt.Println("error while extracting problem dir:", err)
			return
		}

		cur, err = os.Getwd()
		if err != nil {
			fmt.Errorf("failed to get current directory: %w", err)
			return
		}

		leetcodeDirPath, err = config.GetLeetcodeDir()
		if err != nil {
			fmt.Printf("Error getting leetcode directory path: %s\n", err)
			return
		}

		problemDirPath := filepath.Join(leetcodeDirPath, "problems", workbook, problemDir)
		if !dir.IsSubDirOf(cur, filepath.Base(problemDirPath)) {
			fmt.Println("error the current directory is not the problem directory")
			return
		}

		// カレントディレクトリに step_x.py があるかチェック
		var files []os.DirEntry
		// ディレクトリ内のファイルを取得
		files, err = os.ReadDir(cur)
		if err != nil {
			fmt.Println("error reading directory:", err)
			return
		}

		targetFile := "step_x.py"
		fileFound := false
		for _, file := range files {
			if file.Name() == targetFile {
				fileFound = true
				break
			}
		}

		if !fileFound {
			fmt.Println("error step_x file not found")
			return
		}

		// step-count.json の StepNumber をアップデート
		err = step.UpdateStep()
		if err != nil {
			fmt.Println("error updating step:", err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(updateStepCmd)
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"lc/internal/config"
	"lc/internal/problem" // モジュールのインポート
)

var changeProblemCmd = &cobra.Command{
	Use:   "change-problem",
	Short: "Change the problem to work on",
	Long:  `Set the path to the current problem directory in the "problemDirPath" of the config.json.`,
	Run: func(cmd *cobra.Command, args []string) {

		var err error
		leetcodeDirPath, err = config.GetLeetcodeDir()
		if err != nil {
			fmt.Printf("Error getting leetcode directory path: %s\n", err)
			return
		}

		err = problem.SetDirPath(leetcodeDirPath)
		if err != nil {
			fmt.Println("failed to setting problem directory path in config.json:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(changeProblemCmd)
}

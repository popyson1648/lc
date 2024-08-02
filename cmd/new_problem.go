package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"lc/internal/config"
	"lc/internal/problem"
)

var newProblemCmd = &cobra.Command{
	Use:   "new-problem",
	Short: "Create a new problem file",
	Run: func(cmd *cobra.Command, args []string) {

		var err error
		leetcodeDirPath, err = config.GetLeetcodeDir()
		if err != nil {
			fmt.Printf("Error getting leetcode directory path: %s\n", err)
			return
		}

		err = problem.CreateDir(leetcodeDirPath)
		if err != nil {
			fmt.Println("failed to creating problem directory:", err)
			return
		}

		err = problem.SetDirPath(leetcodeDirPath)
		if err != nil {
			fmt.Println("failed to setting problem directory path in config.json:", err)
			return
		}

		err = config.CreateStepCountJson(leetcodeDirPath)
		if err != nil {
			fmt.Println("failed to creating step-count.json:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(newProblemCmd)

}

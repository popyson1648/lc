package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"lc/internal/config"
	vscode_leetcode "lc/internal/vscode-leetcode"
	"os"
	"path/filepath"
)

var vscodeLeetcodeVersion string
var leetcodeDirPath string

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the setup to use lc",
	Long:  "Generate the necessary files for lc to integrate with the LeetCode extension, and overwrite the contents of the existing files in the LeetCode extension",
	Run: func(cmd *cobra.Command, args []string) {

		// --version オプションが指定されていない場合
		if !cmd.Flags().Changed("version") {
			fmt.Println("error: please provide the version using --version flag")
			return
		}

		// --version で受け取った値が空の場合
		if vscodeLeetcodeVersion == "" {
			fmt.Println("error: the --version flag requires a value")
			return
		}

		cur, err := os.Getwd()
		if err != nil {
			fmt.Errorf("failed to get current directory: %w", err)
			return
		}

		leetcodeDirPath = filepath.Join(cur, "leetcode")

		err = config.CreateLeetcodeDirAndConfig(leetcodeDirPath, vscodeLeetcodeVersion)
		if err != nil {
			fmt.Println("Error creating files:", err)
			return
		}

		err = vscode_leetcode.WriteFiles(leetcodeDirPath)
		if err != nil {
			fmt.Println("failed to write file to vscode-leetcode directory:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVar(&vscodeLeetcodeVersion, "version", "", "Specify the version for vscode-leetcode")
}

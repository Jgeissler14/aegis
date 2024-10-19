package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	source string
	target string
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Zip files and backup to a target location",
	Run: func(cmd *cobra.Command, args []string) {
		if source == "" || target == "" {
			fmt.Println("Error: --source and --target are required.")
			os.Exit(1)
		}
		fmt.Printf("Starting backup...\nSource: %s\nTarget: %s\n", source, target)

		files, err := os.Open(source)
		if err != nil {
			fmt.Println("Error opening directory:", err)
			return
		}
		defer files.Close()

		fileInfo, err := files.Readdir(-1)
		if err != nil {
			fmt.Println("Error reading directory:", err)
			return
		}

		for _, file := range fileInfo {
			fmt.Println(file.Name())
		}
	},
}

func init() {
	// Add flags for the backup command
	backupCmd.Flags().StringVarP(&source, "source", "s", "", "Source directory to backup")
	backupCmd.Flags().StringVarP(&target, "target", "t", "", "Target location for the backup")

	// Mark --source and --target as required
	backupCmd.MarkFlagRequired("source")
	backupCmd.MarkFlagRequired("target")

	rootCmd.AddCommand(backupCmd)
}

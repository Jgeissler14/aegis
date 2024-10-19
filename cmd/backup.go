package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	source string
	target string
	tozip  string
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Zip files and backup to a target location",
	Run: func(cmd *cobra.Command, args []string) {
		if source == "" || target == "" {
			fmt.Println("Error: --source and --target are required.")
			os.Exit(1)
		}

		// Check if --tozip is provided and if it ends with ".zip"
		if tozip != "" {
			if !strings.HasSuffix(tozip, ".zip") {
				fmt.Println("Error: --tozip must specify a file with a .zip extension.")
				os.Exit(1)
			}
			fmt.Printf("Zipping files to %s...\n", tozip)
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

		if len(tozip) > 0 {
			fmt.Printf("Zipping Files...\n")

			archive, err := os.Create(tozip)
			defer archive.Close()
			zipWriter := zip.NewWriter(archive)
			if err != nil {
				panic(err)
			}

			for _, file := range fileInfo {
				if file.IsDir() {
					return
				}
				fmt.Printf("writing %s to archive...\n", file.Name())

				f, err := os.Open(file.Name())
				if err != nil {
					panic(err)
				}
				defer f.Close()

				w, err := zipWriter.Create(file.Name() + ".zip")
				if err != nil {
					panic(err)
				}

				if _, err := io.Copy(w, f); err != nil {
					panic(err)
				}

				fmt.Println("closing zip archive...")
				zipWriter.Close()
			}
		}
	},
}

func init() {
	// Add flags for the backup command
	backupCmd.Flags().StringVarP(&source, "source", "s", "", "Source directory to backup")
	backupCmd.Flags().StringVarP(&target, "target", "t", "", "Target location for the backup")
	backupCmd.Flags().StringVarP(&tozip, "tozip", "z", "", "Zip the source before backup")

	// Mark --source and --target as required
	backupCmd.MarkFlagRequired("source")
	backupCmd.MarkFlagRequired("target")

	rootCmd.AddCommand(backupCmd)
}

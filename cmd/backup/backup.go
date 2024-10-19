package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Print the version number of Aegis",
	Long:  `All software has versions. This is Aegis'`,
	Run: func(cmd *cobra.Command, args []string) {
		name := flag.String("name", "world", "The name to greet.")
		flag.Parse()

		if flag.NArg() == 0 {
			fmt.Printf("Hello, %s!\n", *name)
		} else if flag.Arg(0) == "list" {
			files, _ := os.Open(".")
			defer files.Close()

			fileInfo, _ := files.Readdir(-1)

			for _, file := range fileInfo {
				fmt.Println(file.Name())
			}
		} else {
			fmt.Printf("Hello, %s!\n", *name)
		}
	},
}

// func init() {
// 	rootCmd.AddCommand(backupCmd)
// }

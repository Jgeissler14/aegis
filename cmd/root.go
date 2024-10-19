package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string
	Verbose bool

	rootCmd = &cobra.Command{
		Use:   "aegis",
		Short: "A secure and automated backup CLI built in Go for seamless data protection.",
		Long: `Aegis is a Golang-powered CLI tool designed for automating backups of files, 
directories, and cloud resources. With Aegis, users can securely store and manage backups, 
handle multiple backup versions, and restore data effortlessly. Featuring encryption, 
cloud integration, and customizable retention policies, Aegis ensures that your critical data 
remains safe and accessible with minimal effort.`,
	}
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Aegis",
	Long:  `All software has versions. This is Aegis'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Aegis v0.1")
	},
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(backupCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

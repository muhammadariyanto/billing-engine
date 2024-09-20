package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	}
)

// Execute executes the Go function.
//
// It has no parameters and returns an error.
func Execute() error {
	return rootCmd.Execute()
}

// init initializes the application configuration.
//
// No parameters.
// No return type.
func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config.yaml)")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
}

// initConfig initializes the configuration for the application.
//
// No parameters.
func initConfig() {
	// Find current working directory (project root).
	cwd, err := os.Getwd()
	cobra.CheckErr(err)

	if cfgFile == "" {
		cfgFile = filepath.Join(cwd, ".config.yaml")
	}
}

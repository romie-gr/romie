package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	homedir "github.com/mitchellh/go-homedir"
)

var cfgFile string
var debugFlag bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "romie",
	Short: "Game ROM Manager",
	Long: `romie is a command line interface for managing game ROMs.

	Find more information at: https://romie-gr.github.io/romie/`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.romie.yaml)")
	rootCmd.PersistentFlags().BoolVar(&debugFlag, "debug", false, "set verbose logs")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Activate debug messages if --debug flag is used.
	if debugFlag {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// Check config file
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".romie" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".romie")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logrus.Info("Using configuration file:", viper.ConfigFileUsed())
	} else {
		logrus.Warnf("Configuration file not found ($HOME/.romie)")
	}
}

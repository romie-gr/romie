package cmd

import (
	"fmt"
	"os"

	"github.com/romie-gr/romie/internal/config"
	"github.com/romie-gr/romie/internal/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var debugFlag bool
var Config config.Configuration

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
		// Search config in home directory with name ".romie" (without extension).
		viper.AddConfigPath(config.GetDefaultConfigPath())
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// Try to read the configuration file twice
	// If no config is found the first time, the auto-generated config should be read the next time
	for i := 0; i < 2; i++ {
		// If a config file is found, read it in.
		err := viper.ReadInConfig()

		switch {
		case err == nil:
			logrus.Info("Using configuration file: ", viper.ConfigFileUsed())
			// linter thinks this break is aimed at the switch statement and calls it redundant
			// so I break form the for loop by making the counter out of bounds
			i = 2
		case i == 0:
			logrus.Warnf("Configuration file not found ($HOME/.romie/config.yml)")
			logrus.Warnf("Creating default config.yml")
			createDefaultConfig()
		default:
			logrus.Errorf("Problem creating configuration file")
			os.Exit(1)
		}
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		logrus.Fatal(err)
	}

	Config.AdjustDirectories()
}

// createDefaultConfig creates a default ./romie/config.yml entry
func createDefaultConfig() {
	romieDir := config.GetDefaultConfigPath()

	if err := downloadFile("config.yml", config.GetDefaultConfigURL(), romieDir, 1, 1); err != nil {
		logrus.Errorf("Failed to download config.yml\n")
		logrus.Errorf("Error: %q\n", err)

		if err = utils.RemoveFile(config.GetDefaultConfigPath() + "/config.yml"); err != nil {
			logrus.Error(err)
		}
	}
}

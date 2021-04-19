package config

import (
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
)

type Configuration struct {
	Download string
	Database string
}

func (c *Configuration) AdjustDirectories() {
	c.Download = expandDirectory(c.Download)
	c.Database = expandDirectory(c.Database)
}

func expandDirectory(location string) string {
	dir, err := homedir.Expand(location)
	if err != nil {
		logrus.Errorf("config entry %s is not a valid location", location)
		os.Exit(1)
	}

	return dir
}

// GetDefaultConfigPath returns the directory where the default config file can be found
func GetDefaultConfigPath() string {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		logrus.Errorf("Error gettins the home directory %q", err)
		os.Exit(1)
	}

	return home + "/.romie"
}

// GetDefaultConfigURL returns the url of the default config
func GetDefaultConfigURL() string {
	return "https://raw.githubusercontent.com/ge0r/romie/add-config/internal/config/config.yml"
}

package cmd

import (
	"github.com/romie-gr/romie/internal/utils"
	"github.com/romie-gr/romie/pkg/websites/emulatorgames"
	log "github.com/sirupsen/logrus"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Retrieve new list of ROM metadata",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the current filepath where the binary of romie is running
		// TODO: To read the config file and use the correct PATH
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}

		// ---- Code Duplication with search.go ---- //
		// EmulatorGames.Net
		if !utils.FileExists(emulatorgames.DBFile) {
			log.Fatalf("There is no database for EmulatorGames.Net: %s", emulatorgames.DBFile)
		}

		link := "https://raw.githubusercontent.com/romie-gr/romie/master/database.emulatorgames.json"

		server, _ := url.Parse(link)
		err = utils.IsOnline(*server)
		if err != nil {
			log.Fatal("Host is not accessible. Skipping update ...")
		} else {
			log.Debug("Host is accessible")
		}

		name := "database.emulatorgames.json"
		if err := downloadFile(name, link, path, 0, 0); err != nil {
			log.Errorf("Failed to update romie\n")
			log.Errorf("Error: %q\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

package cmd

import (
	"net/url"

	"github.com/romie-gr/romie/internal/utils"
	"github.com/romie-gr/romie/pkg/websites/emulatorgames"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Retrieve new list of ROM metadata",
	Run: func(cmd *cobra.Command, args []string) {
		// ---- Code Duplication with search.go ---- //
		// EmulatorGames.Net
		DownloadDB(emulatorgames.DBFilename, emulatorgames.DBLink)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func DownloadDB(db string, dblink string) {
	server, _ := url.Parse(emulatorgames.DBLink)
	err := utils.IsOnline(*server)

	if err != nil {
		log.Fatal("Host is not accessible. Skipping update ...")
	} else {
		log.Debug("Host is accessible")
	}

	if err := downloadFile(db, dblink, Config.Database, 0, 0); err != nil {
		log.Errorf("Failed to update romie\n")
		log.Errorf("Error: %q\n", err)
	}
}

package cmd

import (
	"os"
	"path/filepath"

	"github.com/romie-gr/romie/internal/scraper"
	"github.com/romie-gr/romie/internal/utils"
	"github.com/romie-gr/romie/pkg/websites/emulatorgames"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove game ROMs",
	Run: func(cmd *cobra.Command, args []string) {
		// ---- Code Duplication with search.go ---- //
		// EmulatorGames.Net
		dbPath := filepath.Join(Config.Database, emulatorgames.DBFilename)

		if !utils.FileExists(dbPath) {
			log.Warnf("There is no database for EmulatorGames.Net: %s", dbPath)
			log.Warnf("Updating database")
			DownloadDB(emulatorgames.DBFilename, emulatorgames.DBLink)
		}

		jsonToEmuDB(dbPath)

		var foundGames []scraper.Rom

		notFound := true

		for _, rom := range emulatorgames.Roms {
			if utils.StringContains(rom.Name, Title) {
				notFound = false
				foundGames = append(foundGames, rom)
			}
		}

		if notFound {
			log.Fatal("No games matching your title")
		}

		log.Infof("Removing %d games ...\n", len(foundGames))

		for _, game := range foundGames {
			dirPath := filepath.Join(Config.Download, game.Console, game.Name)
			log.Debugf("Checking if folder exists: %s\n", dirPath)

			if !utils.FolderExists(dirPath) {
				log.Errorf("%s is not installed. Skip removing ...\n", game.Name)
				continue
			}

			log.Debugf("Folder exist. Removing it now!\n")
			if err := os.RemoveAll(dirPath); err != nil {
				log.Errorf("couldn't remove the game '%s': %q", game.Name, err)
			}
			log.Infof("%s has been successfully removed\n", game.Name)
		}
	},
}

func init() {
	removeCmd.Flags().StringVarP(&Title, "title", "t", "", "Title of the game you want to install")
	_ = removeCmd.MarkFlagRequired("title")
	rootCmd.AddCommand(removeCmd)
}

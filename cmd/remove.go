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
		// Get the current filepath where the binary of romie is running
		// TODO: To read the config file and use the correct PATH
		// TODO: Also set this from one place, not multiple
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}

		// ---- Code Duplication with search.go ---- //
		// EmulatorGames.Net
		if !utils.FileExists(emulatorgames.DBFile) {
			log.Fatalf("There is no database for EmulatorGames.Net: %s", emulatorgames.DBFile)
		}

		jsonToEmuDB(emulatorgames.DBFile)

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
			dirPath := filepath.Join(path, game.Name)
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

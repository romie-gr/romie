package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/romie-gr/romie/internal/utils"
	"github.com/romie-gr/romie/pkg/websites/emulatorgames"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

var Title string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search game ROMs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")

		// EmulatorGames.Net
		if !utils.FileExists(emulatorgames.DBFile) {
			log.Fatalf("There is no database for EmulatorGames.Net: %s", emulatorgames.DBFile)
		}

		jsonToEmuDB(emulatorgames.DBFile)

		notFound := true
		for _, rom := range emulatorgames.Roms {
			if utils.StringContains(rom.Name, Title){
				notFound = false
				fmt.Println(rom.Console, rom.Name, rom.Version, rom.Language, rom.DownloadLink)
			}
		}

		if notFound {
			log.Println("No games matching your title")
		}
	},
}

func init() {
	searchCmd.Flags().StringVarP(&Title, "title", "t", "", "Title of the game you are looking for")
	_ = searchCmd.MarkFlagRequired("title")
	rootCmd.AddCommand(searchCmd)
}

// readDBFile reads the JSON file and returns the JSON data
func readDBFile(file string) []byte {
	fileJSON, err := ioutil.ReadFile(file) // Read the file
	if err != nil {
		log.Fatalf("Could not read the file.\nError: %s\n", err)
	}
	return fileJSON
}

// jsonToEmuDB reads the JSON file and writes the information to the EmulatorsGames.Net database
func jsonToEmuDB(file string) {
	fileJSON := readDBFile(file)
	err := json.Unmarshal(fileJSON, &emulatorgames.Roms)
	if err != nil {
		log.Fatalf("The %s is not a valid JSON format: %v", file, err)
	}
}

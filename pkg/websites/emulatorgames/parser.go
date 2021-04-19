package emulatorgames

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/romie-gr/romie/internal/utils"
	log "github.com/sirupsen/logrus"
)

const DBFilename string = "database.emulatorgames.json"
const DBLink string = "https://raw.githubusercontent.com/romie-gr/romie/master/database.emulatorgames.json"

func getFileDBPath(file string) string {
	if !utils.FileExists(file) {
		if err := utils.CreateFile(file); err != nil {
			log.Fatalf("couldn't create file %s. Error: %q", file, err)
		}
	}

	return file
}

func Parser(console string) {
	if err := supportedConsole(console); err != nil {
		log.Panic(err)
	}

	log.Debugf("Supported console found: %s", console)

	uri := fmt.Sprintf("https://www.emulatorgames.net/roms/%s/", console)
	log.Debugf("Romie will parse: %s", uri)

	// Fetch all the game links for every page
	for _, page := range getPaginationLinks(uri) {
		parsePageGames(page) // fills the collectedGames
	}

	for _, gameURL := range collectedGames {
		parseGame(gameURL, console) // creates Roms saves them into the database JSON file
	}

	for _, v := range Roms {
		log.Println(v)
	}

	saveDBFile(getFileDBPath(DBFilename))
}

func supportedConsole(console string) (err error) {
	if !utils.StringContains(console, emulatorGamesValidConsoles...) {
		err = fmt.Errorf("platform %s is not yet supported", console)
	}

	return err
}

func saveDBFile(filename string) {
	fileJSON, err := json.MarshalIndent(Roms, "", " ")

	if err != nil {
		log.Fatal("Couldn't encode to JSON")
	}

	err = ioutil.WriteFile(filename, fileJSON, 0600)

	if err != nil {
		log.Fatalf("Couldn't update the db file %s\nError: %s", filename, err)
	}
}

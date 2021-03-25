package emulatorgames

import (
	"github.com/romie-gr/romie/internal/scraper"
)

var (
	emulatorGamesValidConsoles = []string{"playstation"}
	collectedGames             []string
	collectedPages             []string
	Roms                       []scraper.Rom
)

package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/romie-gr/romie/internal/scraper/emulatorgames"
)

// scrapeCmd represents the scrape command
var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrapes ROM targets and updates DB",
	Run: func(cmd *cobra.Command, args []string) {
		source := args[0]
		console := args[1]

		switch source {
		case "emulatorgames":
			emulatorgames.Parse(console)
			break
		default:
			log.Fatalf("Source %s found", source)
			break
		}
	},
}

func init() {
	rootCmd.AddCommand(scrapeCmd)
}

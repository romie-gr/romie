package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/romie-gr/romie/internal/scraper"
	"github.com/romie-gr/romie/internal/scraper/emulatorgames"
)

// scrapeCmd represents the scrape command
var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrapes ROM targets and updates DB",
	Run: func(cmd *cobra.Command, args []string) {
		headless, _ := cmd.Flags().GetBool("headless")
		source := args[0]
		console := args[1]

		options := scraper.ScrapeOptions{
			Headless: headless,
		}

		switch source {
		case "emulatorgames":
			emulatorgames.Parse(options, console)
		default:
			log.Fatalf("Source %s found", source)
		}
	},
}

func init() {
	rootCmd.AddCommand(scrapeCmd)
	scrapeCmd.Flags().BoolP(
		"headless",
		"H",
		true,
		"If a browser is required to collect download link, this will configure it to run headless or not. Defaults to true",
	)
}

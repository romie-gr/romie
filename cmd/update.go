package cmd

import (
	"github.com/spf13/cobra"

	"github.com/romie-gr/romie/internal/scraper/emulatorgames"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Retrieve new list of ROM metadata",
	Run: func(cmd *cobra.Command, args []string) {
		emulatorgames.Parse("playstation")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

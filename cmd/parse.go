package cmd

import (
	"github.com/romie-gr/romie/pkg/websites/emulatorgames"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse new game ROMs",
	Run: func(cmd *cobra.Command, args []string) {
		emulatorgames.Parser("playstation")
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}

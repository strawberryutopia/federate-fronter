package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	federate "github.com/strawberryutopia/federate-fronter"
)

// fronterCmd represents the fronter command
var fronterCmd = &cobra.Command{
	Use:     "fronter",
	Aliases: []string{"front"},
	Short:   "Get current fronter",
	Long:    `Reads from the LMHD API and gets the current fronter`,
	Run: func(cmd *cobra.Command, args []string) {
		fronter, _ := federate.GetFronter()
		log.Infof("Fronter: %s\n", fronter.Members[0])
	},
}

func init() {
	rootCmd.AddCommand(fronterCmd)
}

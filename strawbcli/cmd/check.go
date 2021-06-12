package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	federate "github.com/strawberryutopia/federate-fronter"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Gets the current federation status",
	Long:  `Gets the current federation status from various sources`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("check called with args: %v", args)

		// For now, assume Slack workspace was specified

		client, err := federate.NewClient()
		if err != nil {
			log.Fatal(err.Error())
		}

		// TODO: Error handling for when invalid workspace is specified
		client.Slack.Workspaces[args[0]].PrintProfile()
	},
}

func init() {
	federateCmd.AddCommand(checkCmd)
}

package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	federate "github.com/strawberryutopia/federate-fronter"
)

// federateCmd represents the federate command
var federateCmd = &cobra.Command{
	Use:   "federate",
	Short: "Federates fronter details",
	Long:  `Takes current fronter from LMHD API (or flag) and updates various sources`,
	Run: func(cmd *cobra.Command, args []string) {

		nameFlag := cmd.PersistentFlags().Lookup("name")
		avatarFlag := cmd.PersistentFlags().Lookup("avatar")

		var fronter federate.Fronter
		var err error

		if !nameFlag.Changed || !avatarFlag.Changed {
			fronter, err = federate.GetFronter()
			if err != nil {
				log.Fatalf("%s", err)
			}

			log.Infof("Current API Fronter: %v, %v", fronter.Members[0].Name, fronter.Members[0].AvatarURL)
		} else {
			log.Infof("Custom Fronter: %v, %v", nameFlag.Value, avatarFlag.Value)

			fronter = federate.Fronter{
				Members: []federate.Member{
					{
						Name:      nameFlag.Value.String(),
						AvatarURL: avatarFlag.Value.String(),
					},
				},
			}
		}

		// Update Slack(s)
		var slacks *federate.Slack
		slacks, err = federate.NewSlack()
		if err != nil {
			log.Fatal(err.Error)
		}

		err = slacks.Update(
			fronter.Members[0].Name,
			fronter.Members[0].AvatarURL,
		)
		if err != nil {
			log.Fatal(err.Error)
		}
	},
}

func init() {
	rootCmd.AddCommand(federateCmd)

	// TODO: rename this to --fronter once we pull from Vault
	federateCmd.PersistentFlags().String("name", "", "Override name")

	// TODO: remove this once we pull from Vault
	federateCmd.PersistentFlags().String("avatar", "", "Override avatar")
}

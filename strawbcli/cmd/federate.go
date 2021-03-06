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

		var fronter string

		client, err := federate.NewClient()
		if err != nil {
			log.Fatal(err.Error())
		}

		nameFlag := cmd.PersistentFlags().Lookup("name")
		avatarFlag := cmd.PersistentFlags().Lookup("avatar")

		if !nameFlag.Changed || !avatarFlag.Changed {
			fronter, err = client.UpdateFromFront()
			if err != nil {
				log.Fatal(err.Error())
			}
		} else {
			log.Infof("Custom Fronter: %v, %v", nameFlag.Value, avatarFlag.Value)

			fronter = nameFlag.Value.String()

			err := client.Update(
				nameFlag.Value.String(),
				avatarFlag.Value.String(),
			)
			if err != nil {
				log.Fatal(err.Error())
			}
		}

		log.Infof("Federated fronter: %v", fronter)
	},
}

func init() {
	rootCmd.AddCommand(federateCmd)

	// TODO: rename this to --fronter once we pull from Vault
	federateCmd.PersistentFlags().String("name", "", "Override name")

	// TODO: remove this once we pull from Vault
	federateCmd.PersistentFlags().String("avatar", "", "Override avatar")
}

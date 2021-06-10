package cmd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// fronterCmd represents the fronter command
var fronterCmd = &cobra.Command{
	Use:     "fronter",
	Aliases: []string{"front"},
	Short:   "Get current fronter",
	Long:    `Reads from the LMHD API and gets the current fronter`,
	Run: func(cmd *cobra.Command, args []string) {
		fronter, _ := GetFronter()
		log.Infof("Fronter: %s\n", fronter.Members[0])
	},
}

func init() {
	rootCmd.AddCommand(fronterCmd)
}

type Member struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Pronouns    string `json:"pronouns"`
	AvatarURL   string `json:"avatar_url"`
}

type Fronter struct {
	Members []Member `json:"members"`
}

func GetFronter() (Fronter, error) {
	resp, err := http.Get("https://api.lmhd.me/v1/front.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read body
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var fronter Fronter
	err = json.Unmarshal(b, &fronter)
	if err != nil {
		log.Fatal(err)
	}

	return fronter, nil
}

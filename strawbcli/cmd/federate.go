package cmd

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	_ "github.com/joho/godotenv/autoload"

	"github.com/slack-go/slack"
)

// federateCmd represents the federate command
var federateCmd = &cobra.Command{
	Use:   "federate",
	Short: "Federates fronter details",
	Long:  `Takes current fronter from LMHD API (or flag) and updates various sources`,
	Run: func(cmd *cobra.Command, args []string) {

		nameFlag := cmd.PersistentFlags().Lookup("name")
		avatarFlag := cmd.PersistentFlags().Lookup("avatar")

		var fronter Fronter
		var err error

		if !nameFlag.Changed || !avatarFlag.Changed {
			fronter, err = GetFronter()
			if err != nil {
				log.Fatalf("%s", err)
			}

			log.Infof("Fronter: %s\n", fronter.Members[0])
		} else {
			log.Infof("Custom Fronter: %v, %v", nameFlag.Value, avatarFlag.Value)

			fronter = Fronter{
				Members: []Member{
					{
						Name:      nameFlag.Value.String(),
						AvatarURL: avatarFlag.Value.String(),
					},
				},
			}
		}

		client := slack.New(os.ExpandEnv("${SLACK_API_TOKEN}"))
		UpdateProfile(client, fronter.Members[0].Name)
		UpdateAvatar(client, fronter.Members[0].AvatarURL)

		PrintProfile(client)
	},
}

func init() {
	rootCmd.AddCommand(federateCmd)

	federateCmd.PersistentFlags().String("name", "", "Override name")
	federateCmd.PersistentFlags().String("avatar", "", "Override avatar")
}

func PrintProfile(client *slack.Client) {
	profile, err := client.GetUserProfile(&slack.GetUserProfileParameters{
		UserID: os.ExpandEnv("${SLACK_USER_ID}"),
	})
	if err != nil {
		log.Fatal(err)
	}

	profileJSON, _ := json.MarshalIndent(profile, "", "\t")
	log.Infof("Profile: %s\n", profileJSON)
}

func UpdateProfile(client *slack.Client, name string) {
	err := client.SetUserRealName(name)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateAvatar(client *slack.Client, avatar string) {
	file, err := ioutil.TempFile("", "profile.png")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	err = DownloadFile(file.Name(), avatar)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SetUserPhoto(file.Name(), slack.UserSetPhotoParams{})
	if err != nil {
		log.Fatal(err)
	}
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
// Source: https://golangcode.com/download-a-file-from-a-url/
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

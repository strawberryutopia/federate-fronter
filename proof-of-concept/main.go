package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/slack-go/slack"
)

type Fronter struct {
	Members []struct {
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		Pronouns    string `json:"pronouns"`
		AvatarURL   string `json:"avatar_url"`
	} `json:"members"`
}

func main() {

	client := slack.New(os.ExpandEnv("${SLACK_API_TOKEN}"))

	fronter, _ := GetFronter()
	fmt.Printf("Fronter: %s\n", fronter.Members[0])

	UpdateProfile(client, fronter.Members[0].Name)
	UpdateAvatar(client, fronter.Members[0].AvatarURL)

	PrintProfile(client)
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

func PrintProfile(client *slack.Client) {
	profile, err := client.GetUserProfile(&slack.GetUserProfileParameters{
		UserID: os.ExpandEnv("${SLACK_USER_ID}"),
	})
	if err != nil {
		log.Fatal(err)
	}

	profileJSON, _ := json.MarshalIndent(profile, "", "\t")
	fmt.Printf("Profile: %s\n", profileJSON)
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

package federate

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

type SlackWorkspace struct {
	Name   string
	UserID string
	client *slack.Client
}

func NewSlackWorkspace(name, token, userid string) *SlackWorkspace {
	ws := &SlackWorkspace{
		Name:   name,
		UserID: userid,

		client: slack.New(token),
	}

	return ws
}

func (w SlackWorkspace) UpdateProfile(name string) error {
	err := w.client.SetUserRealName(name)
	if err != nil {
		return err
	}

	return nil
}

// TODO: This is currently unused, but could be used to check federation status
// e.g. READ api.lmhd.me/v1/front/federate
func (w SlackWorkspace) PrintProfile(userid string) error {
	profile, err := w.client.GetUserProfile(&slack.GetUserProfileParameters{
		UserID: userid,
	})
	if err != nil {
		return err
	}

	profileJSON, _ := json.MarshalIndent(profile, "", "\t")
	log.Debugf("Profile: %s\n", profileJSON)

	return nil
}

func (w SlackWorkspace) UpdateAvatar(avatar string) error {
	file, err := ioutil.TempFile("", "profile.png")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())

	err = DownloadFile(file.Name(), avatar)
	if err != nil {
		return err
	}

	err = w.client.SetUserPhoto(file.Name(), slack.UserSetPhotoParams{})
	if err != nil {
		return err
	}

	return nil
}

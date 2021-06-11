package federate

import (
	"os"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	_ "github.com/joho/godotenv/autoload"
)

type Slack struct {
	Workspaces map[string]*SlackWorkspace
}

func NewSlack() (*Slack, error) {
	ws := make(map[string]*SlackWorkspace)
	s := &Slack{
		Workspaces: ws,
	}

	names := strings.Split(os.ExpandEnv("${SLACK_WORKSPACES}"), ",")
	tokens := strings.Split(os.ExpandEnv("${SLACK_API_TOKENS}"), ",")
	userIDs := strings.Split(os.ExpandEnv("${SLACK_USER_IDS}"), ",")

	// Assume each of the above has the same name
	// It'll throw a runtime error otherwise
	for i := range names {

		ws := NewSlackWorkspace(
			names[i],
			tokens[i],
			userIDs[i],
		)

		s.Workspaces[names[i]] = ws
	}

	// TODO: Read from Vault
	// Tokens: map[string]string
	// UIDs: map[string]string
	// Members: map[string]complicated struct goes here

	return s, nil
}

func (s Slack) Update(name, avatar string) error {

	var wg sync.WaitGroup

	// TODO: Pass back errors through channels

	for _, ws := range s.Workspaces {
		wg.Add(2)

		go func(wg *sync.WaitGroup, ws *SlackWorkspace) {
			defer wg.Done()
			log.Infof("Updating Name in Slack Workspace: %v", ws.Name)
			err := ws.UpdateProfile(name)
			if err != nil {
				log.Errorf("%v", err)
			}
		}(&wg, ws)

		go func(wg *sync.WaitGroup, ws *SlackWorkspace) {
			defer wg.Done()
			log.Infof("Updating Avatar in Slack Workspace: %v", ws.Name)
			err := ws.UpdateAvatar(avatar)
			if err != nil {
				log.Errorf("%v", err)
			}
		}(&wg, ws)
	}

	log.Info("Slack.Update: Waiting for workers to finish")
	wg.Wait()
	log.Info("Slack.Update: Completed")

	return nil
}

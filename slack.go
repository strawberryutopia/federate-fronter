package federate

import (
	"fmt"
	"os"
	"sync"

	vaultAPI "github.com/hashicorp/vault/api"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"

	_ "github.com/joho/godotenv/autoload"
)

type Slack struct {
	Workspaces map[string]*SlackWorkspace
	Members    SlackMembers
}

func NewSlack(v *vaultAPI.Client) (*Slack, error) {
	ws := make(map[string]*SlackWorkspace)
	s := &Slack{
		Workspaces: ws,
	}

	// TODO: Get slack members
	_, err := getSlackMembers(v)
	if err != nil {
		return nil, fmt.Errorf("error getting Slack Members: %s", err)
	}

	tokens, err := getSlackTokens(v)
	if err != nil {
		return nil, fmt.Errorf("error getting Slack Tokens: %s", err)
	}

	var names []string
	for k := range tokens {
		names = append(names, k)
	}

	userIDs, err := getSlackUserIDs(v)
	if err != nil {
		return nil, fmt.Errorf("error getting Slack User IDs: %s", err)
	}

	// Assume each of the above has the same name
	// It'll throw a runtime error otherwise
	for i := range names {

		ws := NewSlackWorkspace(
			names[i],
			tokens[names[i]],
			userIDs[names[i]],
		)

		s.Workspaces[names[i]] = ws
	}

	return s, nil
}

// Update takes a specific name and avatar, and updates each Slack workspace
func (s Slack) Update(name, avatar string) error {
	// TODO: Get name and avatar from Members

	// Run each Slack API command concurrently
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

func getSlackTokens(v *vaultAPI.Client) (map[string]string, error) {
	tokens := make(map[string]string)

	path := os.ExpandEnv("${VAULT_SECRET_SLACK_TOKENS}")

	log.Infof("Reading tokens from %v", path)
	s, err := v.Logical().Read(path)
	if err != nil {
		return nil, fmt.Errorf("could not read from Vault: %v", err)
	}

	err = mapstructure.Decode(s.Data["data"], &tokens)
	if err != nil {
		return nil, fmt.Errorf("could decode secret from Vault: %v", err)
	}

	return tokens, nil
}

func getSlackMembers(v *vaultAPI.Client) (SlackMembers, error) {
	members := make(SlackMembers)

	path := os.ExpandEnv("${VAULT_SECRET_SLACK_MEMBERS}")

	log.Infof("Reading members from %v", path)
	s, err := v.Logical().Read(path)
	if err != nil {
		return nil, fmt.Errorf("could not read from Vault: %v", err)
	}

	err = mapstructure.Decode(s.Data["data"], &members)
	if err != nil {
		return nil, fmt.Errorf("could decode secret from Vault: %v", err)
	}

	return members, nil
}

func getSlackUserIDs(v *vaultAPI.Client) (map[string]string, error) {
	userIDs := make(map[string]string)

	path := os.ExpandEnv("${VAULT_SECRET_SLACK_USER_IDS}")

	log.Infof("Reading User IDs from %v", path)
	s, err := v.Logical().Read(path)
	if err != nil {
		return nil, fmt.Errorf("could not read from Vault: %v", err)
	}

	err = mapstructure.Decode(s.Data["data"], &userIDs)
	if err != nil {
		return nil, fmt.Errorf("could decode secret from Vault: %v", err)
	}

	return userIDs, nil
}

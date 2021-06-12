package federate

import (
	"fmt"
	"sync"

	vaultAPI "github.com/hashicorp/vault/api"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	Slack *Slack
	//Vault *vaultAPI.Client

	// TODO: Discord
	// TODO: Other Things
}

func NewClient() (*Client, error) {
	// TODO: Vault auth, e.g. AWS or AppRole
	vaultClient, err := vaultAPI.NewClient(
		vaultAPI.DefaultConfig(),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating Vault Client: %v", err)
	}

	// TODO: AppRole/AWS/Etc. Auth

	// Update Slack(s)
	slacks, err := NewSlack(vaultClient)
	if err != nil {
		return nil, fmt.Errorf("error creating Slack Clients: %v", err)
	}

	c := &Client{
		Slack: slacks,
		//Vault: vaultClient,
	}
	return c, nil
}

func (c Client) UpdateFromFront() error {
	fronter, err := GetFronter()
	if err != nil {
		return err
	}

	log.Infof("Current API Fronter: %v, %v", fronter.Members[0].Name, fronter.Members[0].AvatarURL)

	return c.Update(fronter.Members[0].Name, fronter.Members[0].AvatarURL)
}

func (c Client) Update(name, avatar string) error {

	// Run each federation target concurrently
	var wg sync.WaitGroup

	// TODO: Pass back errors through channels

	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		log.Debugf("Updating Slack Workspaces")

		err := c.Slack.Update(
			name,
			avatar,
		)
		if err != nil {
			log.Fatal(err.Error)
		}
	}(&wg)

	// TODO: Discord etc.

	log.Debugf("Client.Update: Waiting for workers to finish")
	wg.Wait()
	log.Debugf("Client.Update: Completed")

	return nil
}

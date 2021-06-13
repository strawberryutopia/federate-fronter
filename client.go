package federate

import (
	"fmt"
	"os"
	"sync"

	vaultAPI "github.com/hashicorp/vault/api"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	Slack *Slack
	Vault *vaultAPI.Client

	// TODO: Discord
	// TODO: Other Things
}

func NewClient() (*Client, error) {

	c := &Client{}

	vaultClient, err := vaultAPI.NewClient(
		vaultAPI.DefaultConfig(),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating Vault Client: %v", err)
	}
	c.Vault = vaultClient
	err = c.VaultAuth()
	if err != nil {
		return nil, fmt.Errorf("error authenticating with Vault: %v", err)
	}

	// Update Slack(s)
	slacks, err := NewSlack(vaultClient)
	if err != nil {
		return nil, fmt.Errorf("error creating Slack Clients: %v", err)
	}
	c.Slack = slacks

	return c, nil
}

func (c Client) UpdateFromFront() (string, error) {
	fronter, err := GetFronter()
	if err != nil {
		return "", err
	}

	// If we have a fronter...
	if len(fronter.Members) > 0 {
		log.Infof("Current API Fronter: %v, %v", fronter.Members[0].Name, fronter.Members[0].AvatarURL)
		return fronter.Members[0].Name, c.Update(fronter.Members[0].Name, fronter.Members[0].AvatarURL)
	}

	// Otherwise, front is either empty, or we have a private fronter
	log.Infof("Current API Fronter empty or private")
	return "Clear", c.Update("_empty", "")
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

func (c Client) VaultAuth() error {
	// TODO: Check for an env var

	if os.ExpandEnv("${VAULT_APPROLE_ROLE_ID}") != "" {

		if os.ExpandEnv("${VAULT_APPROLE_SECRET_ID}") != "" {
			log.Debugf("AppRole Auth with Secret ID")

			token, err := c.Vault.Logical().Write("auth/approle/login", map[string]interface{}{
				"role_id":   os.ExpandEnv("${VAULT_APPROLE_ROLE_ID}"),
				"secret_id": os.ExpandEnv("${VAULT_APPROLE_SECRET_ID}"),
			})
			if err != nil {
				return fmt.Errorf("error authenticating with AppRole: %v", err)
			}

			c.Vault.SetToken(token.Auth.ClientToken)

			return nil
		}

		log.Debugf("AppRole Auth without Secret ID")

		token, err := c.Vault.Logical().Write("auth/approle/login", map[string]interface{}{
			"role_id": os.ExpandEnv("${VAULT_APPROLE_ROLE_ID}"),
		})
		if err != nil {
			return fmt.Errorf("error authenticating with AppRole: %v", err)
		}

		c.Vault.SetToken(token.Auth.ClientToken)

		return nil
	}

	// TODO: Check for a Token File

	// TODO: Do we want to do an auth/token/lookup-self to check if we have a token already from env?
	return nil
}

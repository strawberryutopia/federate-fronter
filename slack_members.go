package federate

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

type SlackMember struct {
	UseAPIAvatar       *bool   `mapstructure:"use_api_avatar,omitempty"`
	UseAPIName         *bool   `mapstructure:"use_api_name,omitempty"`
	Avatar             *string `mapstructure:"avatar,omitempty"`
	Name               *string `mapstructure:"name,omitempty"`
	WorkspaceOverrides map[string]struct {
		Avatar *string `mapstructure:"avatar,omitempty"`
		Name   *string `mapstructure:"name,omitempty"`
	} `mapstructure:"workspace_overrides,omitempty"`
}

type SlackMembers map[string]SlackMember

// TODO: These functions need tests!

func (sm SlackMembers) GetName(fronter, workspace string) (name string) {
	log.Debugf("Getting name for %s in %s", fronter, workspace)

	// Use lcase fronter as a key for our SlackMembers config
	fronterLcase := strings.ToLower(fronter)

	// Do we have specific member config for this fronter?
	if config, exists := sm[fronterLcase]; exists {
		log.Debugf("Fronter Config: %v", config)

		// Do we have specific workspace config for this fronter?
		workspaceConfig, wsConfigExists := config.WorkspaceOverrides[workspace]
		if wsConfigExists {
			if workspaceConfig.Name != nil {
				log.Debugf("Fronter name defined for %s: %s",
					workspace,
					*workspaceConfig.Name)

				return *workspaceConfig.Name
			}

			// So Workspace config exists, but does not specify a name
			// Fallback then to Member config
		}

		// Do we have a name defined at the Member level?
		if config.Name != nil {
			log.Debugf("Fronter name defined: %s", *config.Name)
			return *config.Name
		}

	}

	log.Debugf("No name configured for %s in %s", fronterLcase, workspace)

	if config, exists := sm["_default"]; exists {
		log.Debugf("Default Fronter Config: %v", config)

		// Do we have specific workspace config for this fronter?
		workspaceConfig, wsConfigExists := config.WorkspaceOverrides[workspace]
		if wsConfigExists {
			if workspaceConfig.Name != nil {
				log.Debugf("Fronter name defined for %s: %s",
					workspace,
					*workspaceConfig.Name)

				return *workspaceConfig.Name
			}

			// So Workspace config exists, but does not specify a name
			// Fallback then to Member config
		}

		// Do we have a name defined at the Member level?
		if config.Name != nil {
			log.Debugf("Fronter name defined: %s", *config.Name)
			return *config.Name
		}
	}

	log.Debugf("No defaults configured")

	// If we've got to this point, and we still haven't got a name
	// use the one specified when we called the function
	return fronter
}

func (sm SlackMembers) GetAvatar(fronter, workspace, api_avatar string) string {

	return api_avatar
}

package federate

import (
	"encoding/json"
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

func (m SlackMember) ToString() string {

	json, _ := json.Marshal(m)

	return string(json)
}

type SlackMembers map[string]SlackMember

// TODO: These functions need tests!

func (sm SlackMembers) GetName(fronter, workspace string) string {
	log.Debugf("Getting name for %s in %s", fronter, workspace)

	// Use lcase fronter as a key for our SlackMembers config
	fronterLcase := strings.ToLower(fronter)

	for _, fronterOrDefault := range []string{fronterLcase, "_default"} {
		// Do we have specific member config for this fronter?
		if config, exists := sm[fronterOrDefault]; exists {
			log.Debugf("Fronter %s Config: %v", fronterOrDefault, config.ToString())

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

			log.Debugf("No name configured for %s in %s", fronterOrDefault, workspace)
		}
	}

	// If we've got to this point, and we still haven't got a name
	// use the one specified when we called the function
	return fronter
}

func (sm SlackMembers) GetAvatar(fronter, workspace, api_avatar string) string {
	log.Debugf("Getting avatar for %s in %s", fronter, workspace)

	// Use lcase fronter as a key for our SlackMembers config
	fronterLcase := strings.ToLower(fronter)

	for _, fronterOrDefault := range []string{fronterLcase, "_default"} {
		// Do we have specific member config for this fronter?
		if config, exists := sm[fronterOrDefault]; exists {
			log.Debugf("Fronter %s Config: %v", fronterOrDefault, config.ToString())

			// Do we have specific workspace config for this fronter?
			workspaceConfig, wsConfigExists := config.WorkspaceOverrides[workspace]
			if wsConfigExists {
				if workspaceConfig.Avatar != nil {
					log.Debugf("Fronter avatar defined for %s: %s",
						workspace,
						*workspaceConfig.Avatar)

					return *workspaceConfig.Avatar
				}

				// So Workspace config exists, but does not specify an avatar
				// Fallback then to Member config
			}

			// Do we have an avatar defined at the Member level?
			if config.Avatar != nil {
				log.Debugf("Fronter avatar defined: %s", *config.Avatar)
				return *config.Avatar
			}

			log.Debugf("No avatar configured for %s in %s", fronterOrDefault, workspace)
		}
	}

	// If we've got to this point, and we still haven't got an avatar
	// use the one specified when we called the function
	return api_avatar
}

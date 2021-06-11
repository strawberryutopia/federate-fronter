package federate

type SlackMember struct {
	UseAPIAvatar       bool   `mapstructure:"use_api_avatar,omitempty"`
	UseAPIName         bool   `mapstructure:"use_api_name,omitempty"`
	Avatar             string `mapstructure:"avatar,omitempty"`
	Name               string `mapstructure:"name,omitempty"`
	WorkspaceOverrides map[string]struct {
		Avatar string `mapstructure:"avatar,omitempty"`
		Name   string `mapstructure:"name,omitempty"`
	} `mapstructure:"workspace_overrides,omitempty"`
}

type SlackMembers map[string]SlackMember

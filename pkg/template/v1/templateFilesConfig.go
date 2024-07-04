package v1

import (
	channelsV1 "github.com/deployix/deployed/pkg/channels/v1"
	configV1 "github.com/deployix/deployed/pkg/config/v1"
	promotionsV1 "github.com/deployix/deployed/pkg/promotions/v1"
)

type TemplateFilesConfig struct {
	Channels   *channelsV1.Channels
	Config     *configV1.Config
	Promotions *promotionsV1.Promotions
}

// GetGitType returns the git provider specified in the config.yml file
func (tc *TemplateFilesConfig) GetGitType() string {
	return tc.Config.GitConfig.Provider
}

func NewTemplateConfig() (*TemplateFilesConfig, error) {
	channels, err := channelsV1.GetChannels()
	if err != nil {
		return nil, err
	}

	config, err := configV1.GetConfig()
	if err != nil {
		return nil, err
	}

	promotions, err := promotionsV1.GetPromotions()
	if err != nil {
		return nil, err
	}

	return &TemplateFilesConfig{
		Channels:   channels,
		Config:     config,
		Promotions: promotions,
	}, nil
}

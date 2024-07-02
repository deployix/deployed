package v1

import (
	channelsV1 "github.com/deployix/deployed/pkg/channels/v1"
	promotionsV1 "github.com/deployix/deployed/pkg/promotions/v1"
)

type TemplateConfig struct {
	Channels   *channelsV1.Channels
	Config     *Config
	Promotions *promotionsV1.Promotions
}

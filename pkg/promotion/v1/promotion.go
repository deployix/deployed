package v1

import (
	"fmt"

	channelsV1 "github.com/deployix/deployed/pkg/channels/v1"
	configV1 "github.com/deployix/deployed/pkg/config/v1"
	utilsV1 "github.com/deployix/deployed/pkg/utils/v1"
)

type Promotion struct {
	Name        string `short:"n" long:"name" yaml:"name"`
	Description string `short:"d" long:"desc" yaml:"description,omitempty"`
	FromChannel string `short:"n" long:"name" yaml:"from_channel"`
	ToChannel   string `short:"n" long:"name" yaml:"to_channel"`
	Crontime    string `short:"c" long:"crontime" yaml:"crontime"`
}

// Promote promotes a promotion
func (p *Promotion) Promote() error {
	config, err := configV1.GetConfig()
	if err != nil {
		return err
	}

	// get toChannel
	channels, err := channelsV1.GetChannels()
	if err != nil {
		return err
	}

	// verify channels exist
	if !channels.ChannelExists(p.FromChannel) {
		return fmt.Errorf("channel %s does not exist", p.FromChannel)
	}

	if !channels.ChannelExists(p.ToChannel) {
		return fmt.Errorf("channel %s does not exist", p.ToChannel)
	}

	// promote channel TODO: refactor
	updatedChannel := channels.Channels[p.ToChannel]

	updatedChannel.AppendActionableVersion(updatedChannel.ActionableVersion)

	updatedChannel.ActionableVersion = channels.Channels[p.FromChannel].ActionableVersion
	updatedChannel.ActionableVersion.DateTime = utilsV1.GetCurrentDateTimeAsString(utilsV1.DateTimeLayoutFromTypeName(config.DateTimeFormat))

	channels.Channels[p.ToChannel] = updatedChannel

	// write to file
	return channels.WriteToFile()
}

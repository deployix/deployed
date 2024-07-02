package v1

import (
	"context"
	"fmt"

	channelsV1 "github.com/deployix/deployed/pkg/channels/v1"
	utilsV1 "github.com/deployix/deployed/pkg/utils/v1"
)

// PromotionType is the tyoe of promotion we want (i.e. manual or crontime)
type PromotionType string

type Promotion struct {
	Name        string        `short:"n" long:"name" yaml:"name"`
	Type        PromotionType `short:"t" long:"type" yaml:"type"`
	Description string        `short:"d" long:"desc" yaml:"description,omitempty"`
	FromChannel string        `short:"n" long:"name" yaml:"from_channel"`
	ToChannel   string        `short:"n" long:"name" yaml:"to_channel"`
	Crontime    string        `short:"c" long:"crontime" yaml:"crontime"`
}

type PromoteInput struct {
	// DateTimeFormat format user wants to use to display datetime
	DateTimeFormat string `yaml:"dateTimeFormat"`
}

// Promote promotes a promotion
func (p *Promotion) Promote(ctx context.Context, input PromoteInput) error {

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
	updatedChannel.ActionableVersion.DateTime = utilsV1.GetCurrentDateTimeAsString(utilsV1.DateTimeLayoutFromTypeName(input.DateTimeFormat))

	channels.Channels[p.ToChannel] = updatedChannel

	// write to file
	return channels.WriteToFile()
}

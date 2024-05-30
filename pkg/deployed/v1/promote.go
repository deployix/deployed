package v1

import (
	"fmt"

	channelsV1 "github.com/deployix/deployed/pkg/channels/v1"
	configV1 "github.com/deployix/deployed/pkg/config/v1"
	promotionV1 "github.com/deployix/deployed/pkg/promotion/v1"
	promotionsV1 "github.com/deployix/deployed/pkg/promotions/v1"
	utilsV1 "github.com/deployix/deployed/pkg/utils/v1"
)

type PromoteInput struct {
	// Name is the promotion name we are targeting
	Name string
}

type PromoteOutput struct {
	DateTime        string
	PreviousVersion string
	ActiveVersion   string
}

// GeneratePromotionArgs returns args to run a promotion
func (client Deployed) Promote(input *PromoteInput) (*PromoteOutput, error) {
	// 1. Open config file as a struct
	promotions, err := promotionsV1.GetPromotions()
	if err != nil {
		return nil, err
	}

	// Get promotion by name
	if _, found := promotions.Promotions[input.Name]; !found {
		return nil, fmt.Errorf(fmt.Sprintf("unable to find promotion with the name `%s`", input.Name))
	}

	targetedPromotion := promotions.Promotions[input.Name]

	// 2. update config file
	if err := executePromotion(&targetedPromotion); err != nil {
		return nil, err
	}

	// 3. commit config file to git
	if err := GitPush(); err != nil {
		return nil, err
	}

	promote := &PromoteOutput{}

	return promote, nil
}

// executePromotion performs the actual promotion in the config file
func executePromotion(p *promotionV1.Promotion) error {
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

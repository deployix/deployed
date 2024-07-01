package v1

import (
	"fmt"
	"os"

	actionableVersionV1 "github.com/deployix/deployed/pkg/actionableVersion/v1"
	channelV1 "github.com/deployix/deployed/pkg/channel/v1"
	channelsV1 "github.com/deployix/deployed/pkg/channels/v1"
	configV1 "github.com/deployix/deployed/pkg/config/v1"
	utilsV1 "github.com/deployix/deployed/pkg/utils/v1"
	"github.com/spf13/cobra"
)

var channelsCreateChannelName string
var channelsCreateDescription string
var channelsCreateActionableVersion string

func init() {
	// create required channel name flag
	channelsCreate.Flags().StringVarP(&channelsCreateChannelName, "name", "n", "", "(required) channel name")
	if err := channelsCreate.MarkFlagRequired("name"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// create channel description flag
	channelsCreate.Flags().StringVarP(&channelsCreateDescription, "desc", "d", "", "channel description")

	// create channel actionable-version flag
	channelsCreate.Flags().StringVarP(&channelsCreateActionableVersion, "actionable-version", "v", "", "version that is deployable for this channel")

	channels.AddCommand(channelsCreate)
}

var channelsCreate = &cobra.Command{
	Use:          "create",
	RunE:         channelCreateRun,
	SilenceUsage: true,
}

func channelCreateRun(cmd *cobra.Command, args []string) error {
	// validate flags
	if err := cmd.ValidateRequiredFlags(); err != nil {
		return err
	}

	// populate channels var from channels.yml file if it exists
	channels, err := channelsV1.GetChannels()
	if err != nil {
		return err
	}

	config, err := configV1.GetConfig()
	if err != nil {
		return err
	}
	// check if channel with the same name already exists
	if _, found := channels.Channels[channelsCreateChannelName]; found {
		return fmt.Errorf(fmt.Sprintf("channel with the name '%s' already exists", channelsCreateChannelName))
	}

	// add channel
	newChannel := channelV1.Channel{}

	if channelsCreateDescription != "" {
		newChannel.Description = channelsCreateDescription
	}

	if channelsCreateActionableVersion != "" {
		newChannel.ActionableVersion = actionableVersionV1.ActionableVersion{
			Version:  channelsCreateActionableVersion,
			DateTime: utilsV1.GetCurrentDateTimeAsString(utilsV1.DateTimeLayoutFromTypeName(config.DateTimeFormat)),
		}
	}

	channels.Channels[channelsCreateChannelName] = newChannel

	return channels.WriteToFile()
}

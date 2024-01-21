package deployed

import (
	"fmt"
	"os"

	"github.com/deployix/deployed/internal/utils"
	"github.com/spf13/cobra"
)

var channelsCreateChannelName string
var channelsCreateDescription string
var channelsCreateActionableVersion string

func init() {
	// create required channel name flag
	channelsCreate.Flags().StringVarP(&channelsCreateChannelName, "name", "n", "", "(required) channel name")
	if err := channelsCreate.MarkFlagRequired("name"); err != nil {
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
	channels, err := GetChannels()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	config, err := GetConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// check if channel with the same name already exists
	if _, found := channels.Channels[channelsCreateChannelName]; found {
		return fmt.Errorf(fmt.Sprintf("channel with the name '%s' already exists", channelsCreateChannelName))
	}

	// add channel
	newChannel := Channel{}

	if channelsCreateDescription != "" {
		newChannel.Description = channelsCreateDescription
	}

	if channelsCreateActionableVersion != "" {
		newChannel.ActionableVersion = ActionableVersion{
			Version:  channelsCreateActionableVersion,
			DateTime: utils.GetCurrentDateTimeAsString(config.DateTimeFormat), //TODO: convert string to datetype format
		}
	}

	channels.Channels[channelsCreateChannelName] = newChannel

	return channels.WriteToFile()
}

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
	if err := cmd.ValidateRequiredFlags(); err != nil {
		return err
	}
	
	// get channels from file if it exists
	if err := getChannels(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// check if channel with the same name already exists
	if _, found := chs.Channels[channelsCreateChannelName]; found {
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
			DateTime: utils.GetCurrentDateTimeAsString(cfg.DateTimeFormat),
		}
	}

	chs.Channels[channelsCreateChannelName] = newChannel

	// update file
	if err := CreateChannelsFile(); err != nil {
		return fmt.Errorf("unable to create channel. Try running `deployed init` to initialize")
	}

	return nil
}

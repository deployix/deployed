package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var channelsUpdateChannelName string
var channelsUpdateDescription string

func init() {
	update.Flags().StringVarP(&channelsUpdateChannelName, "name", "n", "", "(required) channel name")
	if err := update.MarkFlagRequired("name"); err != nil {
		os.Exit(1)
	}

	// create channel description flag
	update.Flags().StringVarP(&channelsUpdateDescription, "desc", "d", "", "channel description")
	channel.AddCommand(update)
}

var update = &cobra.Command{
	Use:          "update",
	RunE:         channelsUpdateRun,
	SilenceUsage: true,
}

func channelsUpdateRun(cmd *cobra.Command, args []string) error {
	if err := cmd.ValidateRequiredFlags(); err != nil {
		return err
	}

	// get channels from file if it exists
	if err := getChannels(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// check if channel already exists
	if _, found := chs.Channels[channelsUpdateChannelName]; !found {
		return fmt.Errorf(fmt.Sprintf("Channel with the name %s does not exist", channelsUpdateChannelName))
	}

	// update channel
	updatedChannel := chs.Channels[channelsUpdateChannelName]
	if channelsUpdateDescription != "" {
		updatedChannel.Description = channelsUpdateDescription
	}

	chs.Channels[channelsUpdateChannelName] = updatedChannel

	// update file
	if err := CreateChannelsFile(); err != nil {
		return fmt.Errorf("Unable to update channel. Try running `deployed init` to initialize")
	}

	return nil
}

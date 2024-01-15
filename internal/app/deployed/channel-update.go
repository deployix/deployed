package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var channelsUpdateChannelName string
var channelsUpdateDescription string

func init() {
	channelsUpdate.Flags().StringVarP(&channelsUpdateChannelName, "name", "n", "", "(required) channel name")
	if err := channelsUpdate.MarkFlagRequired("name"); err != nil {
		os.Exit(1)
	}

	// create channel description flag
	channelsUpdate.Flags().StringVarP(&channelsUpdateDescription, "desc", "d", "", "channel description")
	channels.AddCommand(channelsUpdate)
}

var channelsUpdate = &cobra.Command{
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
	if _, found := Chs.Channels[channelsUpdateChannelName]; !found {
		return fmt.Errorf(fmt.Sprintf("Channel with the name %s does not exist", channelsUpdateChannelName))
	}

	// update channel
	updatedChannel := Chs.Channels[channelsUpdateChannelName]
	if channelsUpdateDescription != "" {
		updatedChannel.Description = channelsUpdateDescription
	}

	Chs.Channels[channelsUpdateChannelName] = updatedChannel

	// update file
	if err := CreateChannelsFile(); err != nil {
		return fmt.Errorf("Unable to update channel. Try running `deployed init` to initialize")
	}

	return nil
}

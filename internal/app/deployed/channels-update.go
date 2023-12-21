package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const ()

func init() {
	channels.AddCommand(update)
}

var update = &cobra.Command{
	Use:  "update",
	RunE: channelsUpdateRun,
}

func channelsUpdateRun(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("You must specify at least 1 channel name to update")
	}

	// get channels from file if it exists
	if err := getChannels(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, channelName := range args {
		// check if channel already exists
		if _, found := chs.Channels[channelName]; !found {
			return fmt.Errorf(fmt.Sprintf("Channel with the name %s does not exist", channelName))
		}

		// update channel
		chs.Channels[channelName] = Channel{
			Name: channelName,
		}
	}

	// update file
	if err := CreateChannelsFile(); err != nil {
		return fmt.Errorf("Unable to update channels. Try running `deployed init` to initialize")
	}

	return nil
}

package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const ()

func init() {
	channel.AddCommand(del)
}

var del = &cobra.Command{
	Use:        "delete",
	ArgAliases: []string{"ChannelName"},
	RunE:       channelsDeleteRun,
}

func channelsDeleteRun(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("You must specify at least 1 channel name to delete")
	}

	// get channels from file if it exists
	if err := getChannels(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, channelName := range args {
		// delete channel
		delete(chs.Channels, channelName)
	}

	// update file
	if err := CreateChannelsFile(); err != nil {
		return fmt.Errorf("Unable to update channels config.")
	}

	return nil
}

package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var channelsDeleteChannelName string
var channelsDeleteVerifyChannel bool

func init() {
	channelsDelete.Flags().StringVarP(&channelsDeleteChannelName, "name", "n", "", "(required) channel name to delete")
	if err := channelsDelete.MarkFlagRequired("name"); err != nil {
		os.Exit(1)
	}

	channelsDelete.Flags().BoolVarP(&channelsDeleteVerifyChannel, "verify-channel", "v", false, "verify the channel exists before deleting, otherwise throw error")

	channels.AddCommand(channelsDelete)
}

var channelsDelete = &cobra.Command{
	Use:          "delete",
	ArgAliases:   []string{"ChannelName"},
	RunE:         channelsDeleteRun,
	SilenceUsage: true,
}

func channelsDeleteRun(cmd *cobra.Command, args []string) error {
	if err := cmd.ValidateRequiredFlags(); err != nil {
		return err
	}

	// get channels from file if it exists
	if err := getChannels(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if channelsDeleteVerifyChannel {
		if _, found := chs.Channels[channelsDeleteChannelName]; !found {
			return fmt.Errorf("channel %s does not exist", channelsDeleteChannelName)
		}
	}

	// delete channel
	delete(chs.Channels, channelsDeleteChannelName)

	// update file
	if err := CreateChannelsFile(); err != nil {
		return fmt.Errorf("Unable to update channels config.")
	}

	return nil
}

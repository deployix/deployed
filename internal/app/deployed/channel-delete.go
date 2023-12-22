package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var channelsDeleteChannelName string
var channelsDeleteVerifyChannel bool

func init() {
	del.Flags().StringVarP(&channelsDeleteChannelName, "name", "n", "", "(required) channel name to delete")
	if err := del.MarkFlagRequired("name"); err != nil {
		os.Exit(1)
	}

	del.Flags().BoolVarP(&channelsDeleteVerifyChannel, "verify-channel", "v", false, "verify the channel exists before deleting, otherwise throw error")

	channel.AddCommand(del)
}

var del = &cobra.Command{
	Use:          "delete",
	ArgAliases:   []string{"ChannelName"},
	RunE:         channelsDeleteRun,
	SilenceUsage: true,
}

func channelsDeleteRun(cmd *cobra.Command, args []string) error {
	if err := cmd.ValidateRequiredFlags(); err != nil {
		return err
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

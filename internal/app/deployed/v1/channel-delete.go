package v1

import (
	"fmt"
	"os"

	channelsV1 "github.com/deployix/deployed/pkg/channels/v1"
	"github.com/spf13/cobra"
)

var channelsDeleteChannelName string
var channelsDeleteVerifyChannel bool

func init() {
	channelsDelete.Flags().StringVarP(&channelsDeleteChannelName, "name", "n", "", "(required) channel name to delete")
	if err := channelsDelete.MarkFlagRequired("name"); err != nil {
		fmt.Println(err.Error())
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

	// populate channels var from channels.yml file if it exists
	channels, err := channelsV1.GetChannels()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if channelsDeleteVerifyChannel {
		if _, found := channels.Channels[channelsDeleteChannelName]; !found {
			return fmt.Errorf("channel %s does not exist", channelsDeleteChannelName)
		}
	}

	// delete channel
	delete(channels.Channels, channelsDeleteChannelName)

	return channels.WriteToFile()
}

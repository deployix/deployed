package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var channelsCreateChannelName string
var channelsCreateDescription string

func init() {
	// create required channel name flag
	create.Flags().StringVarP(&channelsCreateChannelName, "name", "n", "", "(required) channel name")
	if err := create.MarkFlagRequired("name"); err != nil {
		os.Exit(1)
	}

	// create required channel description flag
	create.Flags().StringVarP(&channelsCreateDescription, "desc", "d", "", "channel description")

	channel.AddCommand(create)
}

var create = &cobra.Command{
	Use:  "create",
	RunE: channelCreateRun,
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
	chs.Channels[channelsCreateChannelName] = Channel{
		Name: channelsCreateChannelName,
	}

	// update file
	if err := CreateChannelsFile(); err != nil {
		return fmt.Errorf("unable to create channel. Try running `deployed init` to initialize")
	}

	return nil
}

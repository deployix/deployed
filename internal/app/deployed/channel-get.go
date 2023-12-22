package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var channelsGetChannelName string

func init() {
	get.Flags().StringVarP(&channelsGetChannelName, "name", "n", "", "(required) channel name")
	if err := get.MarkFlagRequired("name"); err != nil {
		os.Exit(1)
	}
	channel.AddCommand(get)
}

var get = &cobra.Command{
	Use:          "get",
	RunE:         channelsGetRun,
	SilenceUsage: true,
}

func channelsGetRun(cmd *cobra.Command, args []string) error {
	if err := cmd.ValidateRequiredFlags(); err != nil {
		return err
	}

	// get channels from file if it exists
	if err := getChannels(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// check if channel already exists
	if _, found := chs.Channels[channelsGetChannelName]; !found {
		return fmt.Errorf(fmt.Sprintf("Channel with the name %s does not exist", channelsGetChannelName))
	}

	// get channel
	fmt.Println(chs.Channels[channelsGetChannelName])
	return nil
}

package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var channelsGetChannelName string

func init() {
	channelsGet.Flags().StringVarP(&channelsGetChannelName, "name", "n", "", "(required) channel name")
	if err := channelsGet.MarkFlagRequired("name"); err != nil {
		os.Exit(1)
	}
	channels.AddCommand(channelsGet)
}

var channelsGet = &cobra.Command{
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
	if _, found := Chs.Channels[channelsGetChannelName]; !found {
		return fmt.Errorf(fmt.Sprintf("Channel with the name %s does not exist", channelsGetChannelName))
	}

	// get channel
	fmt.Println(Chs.Channels[channelsGetChannelName])
	return nil
}

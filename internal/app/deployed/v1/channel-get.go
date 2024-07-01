package v1

import (
	"fmt"
	"os"

	channelsV1 "github.com/deployix/deployed/pkg/channels/v1"
	"github.com/spf13/cobra"
)

var channelsGetChannelName string

func init() {
	channelsGet.Flags().StringVarP(&channelsGetChannelName, "name", "n", "", "(required) channel name")
	if err := channelsGet.MarkFlagRequired("name"); err != nil {
		fmt.Println(err.Error())
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

	// populate channels var from channels.yml file if it exists
	channels, err := channelsV1.GetChannels()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// check if channel already exists
	if _, found := channels.Channels[channelsGetChannelName]; !found {
		return fmt.Errorf(fmt.Sprintf("Channel with the name %s does not exist", channelsGetChannelName))
	}

	// get channel
	fmt.Println(channels.Channels[channelsGetChannelName])
	return nil
}

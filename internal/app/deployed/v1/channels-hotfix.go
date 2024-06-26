package v1

import (
	"fmt"
	"os"

	channelsV1 "github.com/deployix/deployed/pkg/channels/v1"
	"github.com/spf13/cobra"
)

var channelsHotfixChannelName string
var channelsHotfixVersion string

func init() {
	channelsHotfix.Flags().StringVarP(&channelsHotfixChannelName, "name", "n", "", "(required) channel name to hotfix")
	if err := channelsHotfix.MarkFlagRequired("name"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	channelsHotfix.Flags().StringVarP(&channelsHotfixVersion, "hotfix-version", "v", "", "(required) channels actionable version to apply")
	if err := channelsHotfix.MarkFlagRequired("hotfix-version"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	channels.AddCommand(channelsHotfix)
}

var channelsHotfix = &cobra.Command{
	Use:          "hotfix",
	RunE:         channelsHotfixRun,
	SilenceUsage: true,
}

func channelsHotfixRun(cmd *cobra.Command, args []string) error {
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
	if _, found := channels.Channels[channelsHotfixChannelName]; !found {
		return fmt.Errorf(fmt.Sprintf("Channel with the name %s does not exist", channelsHotfixChannelName))
	}

	// hotfix channels
	hotfixchannel := channels.Channels[channelsHotfixChannelName]
	// hotfixchannel.ActionableVersion = channelsHotfixVersion //TODO: fix
	//TODO: add previous channels to history + check length
	channels.Channels[channelsHotfixChannelName] = hotfixchannel

	return channels.WriteToFile()
}

package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var channelsHotfixChannelName string
var channelsHotfixVersion string

func init() {
	channelsHotfix.Flags().StringVarP(&channelsHotfixChannelName, "name", "n", "", "(required) channel name to hotfix")
	if err := channelsHotfix.MarkFlagRequired("name"); err != nil {
		os.Exit(1)
	}

	channelsHotfix.Flags().StringVarP(&channelsHotfixVersion, "hotfix-version", "v", "", "(required) channels actionable version to apply")
	if err := channelsHotfix.MarkFlagRequired("hotfix-version"); err != nil {
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

	// get channels from file if it exists
	if err := getChannels(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// check if channel already exists
	if _, found := Chs.Channels[channelsHotfixChannelName]; !found {
		return fmt.Errorf(fmt.Sprintf("Channel with the name %s does not exist", channelsHotfixChannelName))
	}

	// hotfix channels
	hotfixchannel := Chs.Channels[channelsHotfixChannelName]
	// hotfixchannel.ActionableVersion = channelsHotfixVersion //TODO: fix
	//TODO: add previous channels to history + check length
	Chs.Channels[channelsHotfixChannelName] = hotfixchannel

	// update channels file
	if err := CreateChannelsFile(); err != nil {
		return fmt.Errorf("Unable to update channels config.")
	}
	return nil
}

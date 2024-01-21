package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const ()

func init() {
	channels.AddCommand(channelsList)
}

var channelsList = &cobra.Command{
	Use:          "list",
	Args:         cobra.ExactArgs(0),
	RunE:         channelsListRun,
	SilenceUsage: true,
}

func channelsListRun(cmd *cobra.Command, args []string) error {
	if err := cmd.ValidateRequiredFlags(); err != nil {
		return err
	}

	// populate channels var from channels.yml file if it exists
	channels, err := GetChannels()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// list channels
	for channelName, _ := range channels.Channels {
		fmt.Println(channelName)
	}
	return nil
}

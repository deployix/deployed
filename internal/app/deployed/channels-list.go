package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const ()

func init() {
	channels.AddCommand(list)
}

var list = &cobra.Command{
	Use:  "list",
	Args: cobra.ExactArgs(0),
	RunE: channelsListRun,
}

func channelsListRun(cmd *cobra.Command, args []string) error {
	// get channels from file if it exists
	if err := getChannels(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// list channels
	for channelName, _ := range chs.Channels {
		fmt.Println(channelName)
	}
	return nil
}

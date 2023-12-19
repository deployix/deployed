package deployed

import (
	"os"

	"github.com/deployix/deployed/configs"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var channels Channels

const ()

type Channels struct {
}

func init() {
	rootCmd.AddCommand(channel)
}

var channel = &cobra.Command{
	Use:     "channel",
	Aliases: []string{"ch"},
	Short:   "",
	Long:    "",
	Run:     channelRun,
}

func channelRun(cmd *cobra.Command, args []string) {
}

func CreateChannelsFile() error {
	channelsYmlData, err := yaml.Marshal(&channels)
	if err != nil {
		return err
	}

	f, err := os.Create(configs.Cfg.GetChannelsPath())
	defer f.Close()
	if err != nil {
		return err
	}
	f.Write(channelsYmlData)
	f.Sync()
	return nil
}

package deployed

import (
	"os"

	"github.com/deployix/deployed/configs"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var chs Channels

const ()

type Channels struct {
	Channels map[string]Channel
}

type Channel struct {
}

func init() {
	rootCmd.AddCommand(channels)
}

var channels = &cobra.Command{
	Use:     "channels",
	Aliases: []string{"chs"},
	Short:   "",
	Long:    "",
	Run:     channelsRun,
}

func channelsRun(cmd *cobra.Command, args []string) {
}

func CreateChannelsFile() error {
	channelsYmlData, err := yaml.Marshal(&chs)
	if err != nil {
		return err
	}

	f, err := os.Create(configs.Cfg.GetChannelsPath())
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(channelsYmlData)
	if err != nil {
		return err
	}

	if err = f.Sync(); err != nil {
		return err
	}
	return nil
}

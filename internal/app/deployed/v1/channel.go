package v1

import (
	"fmt"
	"os"

	channelsV1 "github.com/deployix/deployed/pkg/channels/v1"
	utilsV1 "github.com/deployix/deployed/pkg/utils/v1"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

const ()

var channelsConfigSetDescription string

func init() {
	rootCmd.AddCommand(channels)

	// try to get initialize channels if config file exists.
	// if not don't error out
	_ = getChannels()
}

var channels = &cobra.Command{
	Use:          "channels",
	Short:        "",
	Long:         "",
	Run:          channelsRun,
	SilenceUsage: true,
}

func channelsRun(cmd *cobra.Command, args []string) {
	fmt.Println("RUNNNING")
}

// SetChannelsConfigFlags applies flags for `deployed config set` to use
func SetChannelsConfigFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&channelsConfigSetDescription, "desc", "d", "", "channel description")

}

func CreateChannelsFile() error {
	channels := channelsV1.Channels{}
	channelsYmlData, err := yaml.Marshal(&channels)
	if err != nil {
		return err
	}

	f, err := os.Create(utilsV1.FilePaths.GetChannelsFilePath())
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

func getChannels() error {
	channels := channelsV1.Channels{}
	if _, err := os.Stat(utilsV1.FilePaths.GetChannelsFilePath()); err == nil {
		yamlFile, err := os.ReadFile(utilsV1.FilePaths.GetChannelsFilePath())
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(yamlFile, &channels)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Channels config file does not exists. Make sure the file %s exists", utilsV1.FilePaths.GetChannelsFilePath())
	}
	return nil
}

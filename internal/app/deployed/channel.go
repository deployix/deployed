package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var Chs Channels

const ()

var channelsConfigSetDescription string

type Channels struct {
	Channels map[string]Channel
}

type Channel struct {
	Description       string `yaml:"description,omitempty"`
	ActionableVersion ActionableVersion
	History           []History
}

type History struct {
	Version string
	Date    string
}

type ActionableVersion struct {
	Version  string `yaml:"version,omitempty"`
	DateTime string `yaml:"datetime,omitempty"`
}

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
	channelsYmlData, err := yaml.Marshal(&Chs)
	if err != nil {
		return err
	}

	f, err := os.Create(FilePaths.GetChannelsFilePath())
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

// ChannelExists validates a channel exists and returns true is found otherwise returns false
func ChannelExists(name string) bool {
	if _, found := Chs.Channels[name]; found {
		return true
	}
	return false
}

func getChannels() error {
	if _, err := os.Stat(FilePaths.GetChannelsFilePath()); err == nil {
		yamlFile, err := os.ReadFile(FilePaths.GetChannelsFilePath())
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(yamlFile, &Chs)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Channels config file does not exists. Make sure the file %s exists", FilePaths.GetChannelsFilePath())
	}
	return nil
}

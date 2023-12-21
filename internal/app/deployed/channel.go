package deployed

import (
	"fmt"
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
	Name        string
	Description string
}

func init() {
	rootCmd.AddCommand(channel)
}

var channel = &cobra.Command{
	Use:   "channel",
	Short: "",
	Long:  "",
	Run:   channelsRun,
}

func channelsRun(cmd *cobra.Command, args []string) {
	fmt.Println("RUNNNING")
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

func getChannels() error {
	if _, err := os.Stat(configs.Cfg.GetChannelsPath()); err == nil {
		yamlFile, err := os.ReadFile(configs.Cfg.GetChannelsPath())
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(yamlFile, &chs)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Channels config file does not exists. Make sure the file %s exists", configs.Cfg.GetChannelsPath())
	}
	return nil
}

// validate config file for channels exists before trying to manipulate
func channelsFileExists() bool {
	if _, err := os.Stat(configs.Cfg.GetChannelsPath()); err != nil {
		return false
	}
	return true
}

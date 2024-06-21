package v1

import (
	"fmt"
	"os"

	channelV1 "github.com/deployix/deployed/pkg/channel/v1"
	utilsV1 "github.com/deployix/deployed/pkg/utils/v1"
	"gopkg.in/yaml.v3"
)

type Channels struct {
	Channels map[string]channelV1.Channel
}

// ChannelExists validates a channel exists and returns true is found otherwise returns false
func (c *Channels) ChannelExists(name string) bool {
	if _, found := c.Channels[name]; found {
		return true
	}
	return false
}

func GetChannels() (*Channels, error) {
	if _, err := os.Stat(utilsV1.FilePaths().GetChannelsFilePath()); err == nil {
		channelsConfigFile := &Channels{}
		yamlFile, err := os.ReadFile(utilsV1.FilePaths().GetChannelsFilePath())
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(yamlFile, channelsConfigFile)
		if err != nil {
			return nil, err
		}
		return channelsConfigFile, nil
	}
	return nil, fmt.Errorf("Channels config file does not exists. Make sure the file %s exists", utilsV1.FilePaths().GetChannelsFilePath())
}

func (c *Channels) WriteToFile() error {
	channelsYmlData, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}

	f, err := os.Create(utilsV1.FilePaths().GetChannelsFilePath())
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

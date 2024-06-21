package v1

import (
	"fmt"
	"os"

	utilsV1 "github.com/deployix/deployed/pkg/utils/v1"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ApplicationName string `yaml:"applicationName"`
	// DateTime format user wants as string
	DateTimeFormat string `yaml:"datetimeFormat"`
	// DefaultBranch: The branch that always pushes the latest version
	DefaultBranch string `yaml:"defaultBranch"`
	// DefaultChannel: The channel that always gets the latest version
	DefaultChannel   string           `yaml:"defaultChannel"`
	ChannelsConfig   ChannelsConfig   `yaml:"channels"`
	PromotionsConfig PromotionsConfig `yaml:"promotions"`
	GitConfig        GitConfig        `yaml:"git"`
}

func (c *Config) WriteToFile() error {
	configYmlData, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}

	f, err := os.Create(utilsV1.FilePaths().GetConfigFilePath())
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(configYmlData)
	if err != nil {
		return err
	}

	if err = f.Sync(); err != nil {
		return err
	}
	return nil
}

func GetConfig() (*Config, error) {
	if _, err := os.Stat(utilsV1.FilePaths().GetConfigFilePath()); err == nil {
		configFile := &Config{}
		yamlFile, err := os.ReadFile(utilsV1.FilePaths().GetConfigFilePath())
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(yamlFile, configFile)
		if err != nil {
			return nil, err
		}
		return configFile, nil
	}
	return nil, fmt.Errorf("Config file does not exists. Make sure the file %s exists", utilsV1.FilePaths().GetConfigFilePath())

}

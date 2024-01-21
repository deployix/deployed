package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// ConfigFiles manages all the viper config files that makeup deployed
type ConfigFiles struct {
	ChannelsFile   *viper.Viper
	ConfigFile     *viper.Viper
	PromotionsFile *viper.Viper
}

func init() {
}

func GetConfig() (*Config, error) {
	if _, err := os.Stat(FilePaths.GetConfigFilePath()); err == nil {
		configFile := &Config{}
		yamlFile, err := os.ReadFile(FilePaths.GetConfigFilePath())
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(yamlFile, configFile)
		if err != nil {
			return nil, err
		}
		return configFile, nil
	}
	return nil, fmt.Errorf("Config file does not exists. Make sure the file %s exists", FilePaths.GetConfigFilePath())

}

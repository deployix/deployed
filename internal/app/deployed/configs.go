package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var CfgFiles ConfigFiles

// ConfigFiles manages all the viper config files that makeup deployed
type ConfigFiles struct {
	ChannelsFile   *viper.Viper
	ConfigFile     *viper.Viper
	PromotionsFile *viper.Viper
}

func init() {
}

func InitConfigFiles() {
	// setup channels viper config
	channelsViper, err := InitChannelsConfigFile()
	if err != nil {
		fmt.Errorf(err.Error())
	}

	// setup config viper config
	configViper, err := InitConfigFile()
	if err != nil {
		fmt.Errorf(err.Error())
	}

	// setup config viper config
	promotionsViper, err := InitPromotionsConfigFile()
	if err != nil {
		fmt.Errorf(err.Error())
	}

	// set global file
	CfgFiles = ConfigFiles{
		ChannelsFile:   channelsViper,
		ConfigFile:     configViper,
		PromotionsFile: promotionsViper,
	}
}

func InitChannelsConfigFile() (*viper.Viper, error) {
	if _, err := os.Stat(FilePaths.GetChannelsFilePath()); err != nil {
		return nil, err
	}

	channelsViper := viper.New()
	fmt.Println(FilePaths.GetChannelsFilePath())
	channelsViper.SetConfigFile(FilePaths.GetChannelsFilePath())
	if err := channelsViper.ReadInConfig(); err != nil {
		fmt.Println("error using channels config file:", channelsViper.ConfigFileUsed())
		return nil, err
	}
	channelsViper.AutomaticEnv()
	channelsViper.OnConfigChange(OnChannelsFileChange)

	// populate var with channel config
	if err := channelsViper.Unmarshal(&Chs); err != nil {
		fmt.Printf("error Unmarshal channels config %v", err)
	}

	return channelsViper, nil
}

func InitConfigFile() (*viper.Viper, error) {
	if _, err := os.Stat(FilePaths.GetConfigFilePath()); err != nil {
		return nil, err
	}

	// setup config viper config
	configViper := viper.New()
	configViper.SetConfigFile(FilePaths.GetConfigFilePath())
	if err := configViper.ReadInConfig(); err != nil {
		fmt.Println("error using config file:", configViper.ConfigFileUsed())
		return nil, err
		// TODO: change to print as debug log type
	}
	configViper.AutomaticEnv()
	configViper.OnConfigChange(OnConfigFileChange)

	// populate var with config
	if err := configViper.Unmarshal(&Cfg); err != nil {
		fmt.Printf("error Unmarshal config %v", err)
	}

	return configViper, nil
}

func InitPromotionsConfigFile() (*viper.Viper, error) {
	if _, err := os.Stat(FilePaths.GetPromotionsFilePath()); err != nil {
		return nil, err
	}

	// setup promotions viper config
	promotionsViper := viper.New()
	promotionsViper.SetConfigFile(FilePaths.GetPromotionsFilePath())
	if err := promotionsViper.ReadInConfig(); err != nil {
		fmt.Println("error using config file:", promotionsViper.ConfigFileUsed())
		return nil, err
		// TODO: change to print as debug log type
	}
	promotionsViper.AutomaticEnv()
	promotionsViper.OnConfigChange(OnPromotionsFileChange)

	// populate var with promotion config
	if err := promotionsViper.Unmarshal(&Promos); err != nil {
		fmt.Printf("error Unmarshal promotion config %v", err)
	}

	return promotionsViper, nil
}

package configs

import (
	"fmt"
	"os"

	"github.com/deployix/deployed/internal/constants"
)

type Config struct {
	// path to working directory
	path string

	// Name of directory that contains deployed configurations
	dirName string

	configFileName string

	channelsFileName string

	promotionsFileName string

	versionsFileName string
}

var Cfg Config

func init() {
	// set default Path as current working dir
	if cwd, err := os.Getwd(); err != nil {
		fmt.Println(err)
	} else {
		Cfg.path = cwd
	}

	// set DirName to default value
	Cfg.dirName = constants.DEFAULT_DIR_NAME
	// set ConfigFileName to default value
	Cfg.configFileName = constants.DEFAULT_CONFIG_FILENAME
	// set ChannelsFileName to default value
	Cfg.channelsFileName = constants.DEFAULT_CHANNELS_FILENAME
	// set PromotionsFileName to default value
	Cfg.promotionsFileName = constants.DEFAULT_PROMOTIONS_FILENAME
	// set VersionsFileName to default values
	Cfg.versionsFileName = constants.DEFAULT_VERSIONS_FILENAME

}

func (cfg *Config) GetConfigFileName() string {
	return cfg.configFileName
}

func (cfg *Config) GetPath() string {
	return cfg.path
}

func (cfg *Config) GetDirectoryPath() string {
	return fmt.Sprintf("%s/%s", cfg.path, cfg.dirName)
}

func (cfg *Config) GetConfigPath() string {
	return fmt.Sprintf("%s/%s/%s", cfg.path, cfg.dirName, cfg.configFileName)
}

func (cfg *Config) GetChannelsPath() string {
	return fmt.Sprintf("%s/%s/%s", cfg.path, cfg.dirName, cfg.channelsFileName)
}

func (cfg *Config) GetPromotionsPath() string {
	return fmt.Sprintf("%s/%s/%s", cfg.path, cfg.dirName, cfg.promotionsFileName)
}

func (cfg *Config) GetVersionsPath() string {
	return fmt.Sprintf("%s/%s/%s", cfg.path, cfg.dirName, cfg.versionsFileName)
}

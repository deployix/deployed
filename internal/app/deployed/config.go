package deployed

import (
	"fmt"
	"os"

	"github.com/deployix/deployed/internal/constants"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

const ()

var cfg Config

type Config struct {
	// path to working directory
	path string

	// Name of directory that contains deployed configurations
	dirName string

	configFileName string

	channelsFileName string

	promotionsFileName string

	versionsFileName string
	// MaxHistory: Max length of history versions struct should keep
	// DefaultBranch: The branch that always pushes the latest version
	// DefaultChannel: The channel that always gets the latest version
	// Pause Promotions: Bool flag to pause promotion
	// Notification: struct field to allow users to recive notifications when promotions or updates are done
	// GitHub vs Gitlab code source config
	// DateTime format user wants as string
	DateTimeFormat string         `yaml:"datetimeFormat"`
	Channels       ChannelsConfig `yaml:"channels"`
	Promotions     PromotionsConfig
	Git            GitConfig
}

type ChannelsConfig struct {
	DateTimeFormat string `yaml:"datetimeFormat"`
}

type PromotionsConfig struct {
}

type GitConfig struct {
}

type NotificationConfig struct {
}

type DateTimeFormatType int

// valid datetime formats
const (
	ANSIC DateTimeFormatType = iota
	UnixDate
	RubyDate
	RFC822
	RFC822Z
	RFC850
	RFC1123
	RFC1123Z
	RFC3339
	RFC3339Nano
)

func (d DateTimeFormatType) String() string {
	return [...]string{"ansic", "unixdate", "rubydate", "rfc822", "rfc822z", "rfc850", "rfc1123", "rfc1123z", "rfc3339", "rfc3339nano"}[d]
}

func (d DateTimeFormatType) EnumIndex() int {
	return int(d)
}

func init() {
	rootCmd.AddCommand(config)

	// get config if setup
	_ = getConfig()

	// set default Path as current working dir
	if cwd, err := os.Getwd(); err != nil {
		fmt.Println(err)
	} else {
		cfg.path = cwd
	}

	// set DirName to default value
	if cfg.dirName == "" {
		cfg.dirName = constants.DEFAULT_DIR_NAME
	}

	// set ConfigFileName to default value
	if cfg.configFileName == "" {
		cfg.configFileName = constants.DEFAULT_CONFIG_FILENAME
	}

	// set ChannelsFileName to default value
	if cfg.channelsFileName == "" {
		cfg.channelsFileName = constants.DEFAULT_CHANNELS_FILENAME
	}

	// set PromotionsFileName to default value
	if cfg.promotionsFileName == "" {
		cfg.promotionsFileName = constants.DEFAULT_PROMOTIONS_FILENAME
	}

	// set VersionsFileName to default values
	if cfg.versionsFileName == "" {
		cfg.versionsFileName = constants.DEFAULT_VERSIONS_FILENAME
	}

	if cfg.DateTimeFormat == "" {
		cfg.DateTimeFormat = RFC1123.String()
	}

	// // populate cfg var with whats in .deployed/config.yml file
	// if _, err := os.Stat(cfg.GetConfigPath()); err == nil {
	// 	cfgYmlFile, err := os.ReadFile(cfg.GetConfigPath())
	// 	if err != nil {
	// 		log.Printf("yamlFile.Get err   #%v ", err)
	// 	}
	// 	err = yaml.Unmarshal(cfgYmlFile, cfg)
	// 	if err != nil {
	// 		log.Fatalf("Unmarshal: %v", err)
	// 	}
	// }
}

var config = &cobra.Command{
	Use:     "config",
	Aliases: []string{"cfg"},
	Short:   "",
	Long:    "",
	Run:     configRun,
}

func configRun(cmd *cobra.Command, args []string) {
}

func CreateConfigFile() error {
	configYmlData, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}

	f, err := os.Create(cfg.GetConfigPath())
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

func getConfig() error {
	if _, err := os.Stat(cfg.GetConfigPath()); err == nil {
		yamlFile, err := os.ReadFile(cfg.GetConfigPath())
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(yamlFile, &cfg)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Config file does not exists. Make sure the file %s exists", cfg.GetConfigPath())
	}
	return nil
}

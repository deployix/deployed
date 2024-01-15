package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

const ()

var Cfg Config

type Config struct {
	// // path to working directory
	// path string

	// // Name of directory that contains deployed configurations
	// dirName string

	// configFileName string

	// channelsFileName string

	// promotionsFileName string

	// versionsFileName string

	// GitHub vs Gitlab code source config
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

type ChannelsConfig struct {

	// MaxVersionHistoryLength: The maximum historical versions that should be keep
	MaxVersionHistoryLength int `yaml:"maxVersionHistoryLength"`

	// Pause Promotions: Bool flag to indicate if promotion should be paused
	PausePromotion bool `yaml:"pausePromotion"`

	NotificationConfig NotificationConfig `yaml:"notificationConfig"`
}

type PromotionsConfig struct {
	NotificationConfig NotificationConfig `yaml:"notificationConfig"`
}

type GitType int

const (
	Github = iota
	Gitlab
)

func (t GitType) String() string {
	return [...]string{"github", "gitlab"}[t]
}

func (t GitType) EnumIndex() int {
	return int(t)
}

type GitConfig struct {
	// GitType is the type of git being used (i.e. GitHub, Gitlab)
	GitType string `yaml:"gitType"`
}

type NotificationConfig struct {
	SlackConfig NotificationSlackConfig `yaml:"slackConfig"`
}

type NotificationSlackConfig struct {
	SlackChannel string `yaml:"slackChannel"`
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

	// // set default Path as current working dir
	// if cwd, err := os.Getwd(); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	cfg.path = cwd
	// }

	// if cfg.DateTimeFormat == "" {
	// 	cfg.DateTimeFormat = RFC1123.String()
	// }

	// // // populate cfg var with whats in .deployed/config.yml file
	// // if _, err := os.Stat(cfg.GetConfigPath()); err == nil {
	// // 	cfgYmlFile, err := os.ReadFile(cfg.GetConfigPath())
	// // 	if err != nil {
	// // 		log.Printf("yamlFile.Get err   #%v ", err)
	// // 	}
	// // 	err = yaml.Unmarshal(cfgYmlFile, cfg)
	// // 	if err != nil {
	// // 		log.Fatalf("Unmarshal: %v", err)
	// // 	}
	// // }
}

var config = &cobra.Command{
	Use:   "config",
	Short: "",
	Long:  "",
	Run:   configRun,
}

func configRun(cmd *cobra.Command, args []string) {
}

func CreateConfigFile() error {
	configYmlData, err := yaml.Marshal(&Cfg)
	if err != nil {
		return err
	}

	f, err := os.Create(FilePaths.GetConfigFilePath())
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

// func (cfg *Config) GetConfigFileName() string {
// 	return cfg.configFileName
// }

// func (cfg *Config) GetPath() string {
// 	return cfg.path
// }

// func (cfg *Config) GetDirectoryPath() string {
// 	return fmt.Sprintf("%s/%s", cfg.path, cfg.dirName)
// }

// func (cfg *Config) GetConfigPath() string {
// 	return fmt.Sprintf("%s/%s/%s", cfg.path, cfg.dirName, cfg.configFileName)
// }

// func (cfg *Config) GetChannelsPath() string {
// 	return fmt.Sprintf("%s/%s/%s", cfg.path, cfg.dirName, cfg.channelsFileName)
// }

// func (cfg *Config) GetPromotionsPath() string {
// 	return fmt.Sprintf("%s/%s/%s", cfg.path, cfg.dirName, cfg.promotionsFileName)
// }

// func (cfg *Config) GetVersionsPath() string {
// 	return fmt.Sprintf("%s/%s/%s", cfg.path, cfg.dirName, cfg.versionsFileName)
// }

func getConfig() error {
	if _, err := os.Stat(FilePaths.GetConfigFilePath()); err == nil {
		yamlFile, err := os.ReadFile(FilePaths.GetConfigFilePath())
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(yamlFile, &Cfg)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Config file does not exists. Make sure the file %s exists", FilePaths.GetConfigFilePath())
	}
	return nil
}

package deployed

import (
	"fmt"
	"os"
	"time"

	"github.com/deployix/deployed/internal/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
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

type ChannelsConfig struct {
	// MaxVersionHistoryLength: The maximum historical versions that should be keep
	MaxVersionHistoryLength int `yaml:"maxVersionHistoryLength"`

	NotificationConfig NotificationConfig `yaml:"notificationConfig"`
}

type PromotionsConfig struct {
	// Pause Promotions: Bool flag to indicate if promotion should be paused
	PausePromotion     bool               `yaml:"pausePromotion"`
	NotificationConfig NotificationConfig `yaml:"notificationConfig"`
}

const (
	GitHub = iota
	GitLab
)

type GitConfig struct {
	// GitType is the type of git being used (i.e. GitHub, Gitlab)
	Provider string `yaml:"provider"`
	Domain   string `yaml:"domain"`
	RepoName string `yaml:"repoName"`
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
	return [...]string{time.ANSIC, time.UnixDate, time.RubyDate, time.RFC822, time.RFC822Z, time.RFC850, time.RFC1123, time.RFC1123Z, time.RFC3339, time.RFC3339Nano}[d]
}

func (d DateTimeFormatType) EnumIndex() int {
	return int(d)
}

func GetConfig() (*Config, error) {
	if _, err := os.Stat(utils.FilePaths.GetConfigFilePath()); err == nil {
		configFile := &Config{}
		yamlFile, err := os.ReadFile(utils.FilePaths.GetConfigFilePath())
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(yamlFile, configFile)
		if err != nil {
			return nil, err
		}
		return configFile, nil
	}
	return nil, fmt.Errorf("Config file does not exists. Make sure the file %s exists", utils.FilePaths.GetConfigFilePath())

}

func (c *Config) WriteToFile() error {
	configYmlData, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}

	f, err := os.Create(utils.FilePaths.GetConfigFilePath())
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

func init() {
	rootCmd.AddCommand(config)
}

var config = &cobra.Command{
	Use:   "config",
	Short: "",
	Long:  "",
	Run:   configRun,
}

func configRun(cmd *cobra.Command, args []string) {
}

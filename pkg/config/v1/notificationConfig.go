package v1

type NotificationConfig struct {
	SlackConfig NotificationSlackConfig `yaml:"slackConfig"`
}

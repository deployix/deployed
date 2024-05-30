package v1

type ChannelsConfig struct {
	// MaxVersionHistoryLength: The maximum historical versions that should be keep
	MaxVersionHistoryLength int `yaml:"maxVersionHistoryLength"`

	NotificationConfig NotificationConfig `yaml:"notificationConfig"`
}

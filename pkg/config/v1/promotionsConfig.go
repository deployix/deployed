package v1

type PromotionsConfig struct {
	DeployedGithubAVersions string             `yaml:"deployedGithubAVersions"`
	NotificationConfig      NotificationConfig `yaml:"notificationConfig"`
}

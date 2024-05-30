package v1

type GitConfig struct {
	// GitType is the type of git being used (i.e. GitHub, Gitlab)
	Provider string `yaml:"provider"`
	Domain   string `yaml:"domain"`
	RepoName string `yaml:"repoName"`
}

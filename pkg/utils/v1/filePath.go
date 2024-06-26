package v1

import (
	"path/filepath"
	"strings"

	constantsV1 "github.com/deployix/deployed/pkg/constants/v1"
)

// FilePathConfig manages all filepath related data for the cli
type FilePathsConfig struct {
	// path to working directory
	path string

	// Name of directory that contains deployed configurations
	dirName string

	configFileName string

	channelsFileName string

	promotionsFileName string

	versionsFileName string
}

// FilePaths returns a new instance of FilePathsConfig
func FilePaths() *FilePathsConfig {
	return &FilePathsConfig{
		path:               constantsV1.DEFAULT_FILEPATH,
		dirName:            constantsV1.DEFAULT_DEPLOYED_DIRECTORY,
		configFileName:     constantsV1.DEFAULT_CONFIG_FILENAME,
		channelsFileName:   constantsV1.DEFAULT_CHANNELS_FILENAME,
		promotionsFileName: constantsV1.DEFAULT_PROMOTIONS_FILENAME,
		versionsFileName:   constantsV1.DEFAULT_VERSIONS_FILENAME,
	}
}

// SetRootDir sets the root directory and returns *FilePathsConfig
func (fpc *FilePathsConfig) SetRootDir(rootDir string) *FilePathsConfig {
	fpc.path = filepath.Join(rootDir, fpc.path)
	return fpc
}

func (fpc *FilePathsConfig) GetConfigFileName() string {
	return fpc.configFileName
}

func (fpc *FilePathsConfig) GetPath() string {
	return fpc.path
}

func (fpc *FilePathsConfig) GetDirectoryPath() string {
	return filepath.Join(fpc.GetPath(), fpc.dirName)
}

func (fpc *FilePathsConfig) GetGitDirectoryPath(gitType string) string {
	dir := constantsV1.DEFAULT_GITLAB_TEMPLATES_DIRECTORY_PATH
	if strings.EqualFold(gitType, "gitlab") {
		dir = constantsV1.DEFAULT_GITHUB_TEMPLATES_DIRECTORY_PATH
	}
	return dir
}

// GetGitDirectoryOutputPath returns the template output file path
func (fpc *FilePathsConfig) GetGitDirectoryOutputPath(gitType string) string {
	dir := constantsV1.DEFAULT_GITHUB_ACTIONS_DIRECTORY_PATH
	if strings.EqualFold(gitType, "gitlab") {
		dir = constantsV1.DEFAULT_GITLAB_PIPELINE_DIRECTORY_PATH
	}
	return filepath.Join(fpc.GetPath(), dir)
}

func (fpc *FilePathsConfig) GetConfigFilePath() string {
	return filepath.Join(fpc.GetPath(), fpc.dirName, fpc.configFileName)
}

func (fpc *FilePathsConfig) GetChannelsFilePath() string {
	return filepath.Join(fpc.GetPath(), fpc.dirName, fpc.channelsFileName)
}

func (fpc *FilePathsConfig) GetPromotionsFilePath() string {
	return filepath.Join(fpc.GetPath(), fpc.dirName, fpc.promotionsFileName)
}

func (fpc *FilePathsConfig) GetVersionsFilePath() string {
	return filepath.Join(fpc.GetPath(), fpc.dirName, fpc.versionsFileName)
}

func (fpc *FilePathsConfig) GetTemplatesDirectoryPath() string {
	return filepath.Join(fpc.GetPath(), constantsV1.DEFAULT_TEMPLATES_DIRECTORY_PATH)
}

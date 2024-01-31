package utils

import (
	"fmt"
	"strings"

	"github.com/deployix/deployed/internal/constants"
)

// FilePaths handles file paths used by the CLI
var FilePaths FilePathsConfig = FilePathsConfig{
	path:               constants.DEFAULT_FILEPATH,
	dirName:            constants.DEFAULT_DEPLOYED_DIRECTORY,
	configFileName:     constants.DEFAULT_CONFIG_FILENAME,
	channelsFileName:   constants.DEFAULT_CHANNELS_FILENAME,
	promotionsFileName: constants.DEFAULT_PROMOTIONS_FILENAME,
	versionsFileName:   constants.DEFAULT_VERSIONS_FILENAME,
}

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

func (fpc *FilePathsConfig) GetConfigFileName() string {
	return fpc.configFileName
}

func (fpc *FilePathsConfig) GetPath() string {
	return fpc.path
}

func (fpc *FilePathsConfig) GetDirectoryPath() string {
	return fmt.Sprintf("%s/%s", fpc.path, fpc.dirName)
}

func (fpc *FilePathsConfig) GetGitDirectoryPath(gitType string) string {
	dir := constants.DEFAULT_GITHUB_TEMPLATES_DIRECTORY_PATH
	if strings.EqualFold(gitType, "gitlab") {
		dir = constants.DEFAULT_GITHUB_TEMPLATES_DIRECTORY_PATH
	}
	return dir
}

// GetGitDirectoryOutputPath returns the template output file path
func (fpc *FilePathsConfig) GetGitDirectoryOutputPath(gitType string) string {
	dir := constants.DEFAULT_GITHUB_ACTIONS_DIRECTORY_PATH
	if strings.EqualFold(gitType, "gitlab") {
		dir = constants.DEFAULT_GITLAB_PIPELINE_DIRECTORY_PATH
	}
	return fmt.Sprintf("%s/%s", fpc.path, dir)
}

func (fpc *FilePathsConfig) GetConfigFilePath() string {
	return fmt.Sprintf("%s/%s/%s", fpc.path, fpc.dirName, fpc.configFileName)
}

func (fpc *FilePathsConfig) GetChannelsFilePath() string {
	return fmt.Sprintf("%s/%s/%s", fpc.path, fpc.dirName, fpc.channelsFileName)
}

func (fpc *FilePathsConfig) GetPromotionsFilePath() string {
	return fmt.Sprintf("%s/%s/%s", fpc.path, fpc.dirName, fpc.promotionsFileName)
}

func (fpc *FilePathsConfig) GetVersionsFilePath() string {
	return fmt.Sprintf("%s/%s/%s", fpc.path, fpc.dirName, fpc.versionsFileName)
}

func (fpc *FilePathsConfig) GetTemplatesDirectoryPath() string {
	return fmt.Sprintf("%s/%s", fpc.path, constants.DEFAULT_TEMPLATES_DIRECTORY_PATH)
}

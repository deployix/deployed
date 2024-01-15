package deployed

import (
	"fmt"

	"github.com/deployix/deployed/internal/constants"
)

// FilePaths handles file paths used by the CLI
var FilePaths FilePathsConfig = FilePathsConfig{
	path:               constants.DEFAULT_FILEPATH,
	dirName:            constants.DEFAULT_DIR_NAME,
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

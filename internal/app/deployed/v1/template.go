package v1

import (
	"errors"
	"os"
	"path/filepath"
	textTemplate "text/template"

	channelsV1 "github.com/deployix/deployed/pkg/channels/v1"
	configV1 "github.com/deployix/deployed/pkg/config/v1"
	promotionsV1 "github.com/deployix/deployed/pkg/promotions/v1"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(promotion)
}

var template = &cobra.Command{
	Use: "template",
	Run: templatesRun,
}

func templatesRun(cmd *cobra.Command, args []string) {
}

type TemplateConfig struct {
	Channels   *channelsV1.Channels
	Config     *configV1.Config
	Promotions *promotionsV1.Promotions
}

// GetGitType returns the git provider specified in the config.yml file
func (tc *TemplateConfig) GetGitType() string {
	return tc.Config.GitConfig.Provider
}

func NewTemplateConfig() (*TemplateConfig, error) {
	channels, err := channelsV1.GetChannels()
	if err != nil {
		return nil, err
	}

	config, err := configV1.GetConfig()
	if err != nil {
		return nil, err
	}

	promotions, err := promotionsV1.GetPromotions()
	if err != nil {
		return nil, err
	}

	return &TemplateConfig{
		Channels:   channels,
		Config:     config,
		Promotions: promotions,
	}, nil
}

// func GenerateGitPromotionTemplate() error {
// 	templateConfig, err := NewTemplateConfig()
// 	if err != nil {
// 		return err
// 	}

// 	gitProvider := templateConfig.GetGitType()

// 	path := utilsV1.FilePaths().GetGitDirectoryPath(gitProvider)

// 	files, err := os.ReadDir(path)
// 	if err != nil {
// 		return err
// 	}

// 	for _, f := range files {
// 		filePath := filepath.Join(path, f.Name())
// 		fileDir := utilsV1.FilePaths().GetGitDirectoryOutputPath(gitProvider)
// 		fileName := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
// 		templ, err := template.ParseFiles(filePath)
// 		if err != nil {
// 			return err
// 		}
// 		if err := createFileWithTemplate(templ, fileDir, fileName, templateConfig); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

func createFileWithTemplate(t *textTemplate.Template, dir, filename string, data interface{}) error {
	// create directory if it doesnt exist
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	filePath := filepath.Join(dir, filename)
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, data)
	if err != nil {
		return err
	}

	return nil
}

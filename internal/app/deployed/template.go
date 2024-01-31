package deployed

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/deployix/deployed/internal/utils"
)

type TemplateConfig struct {
	Channels   *Channels
	Config     *Config
	Promotions *Promotions
}

// GetGitType returns the git provider specified in the config.yml file
func (tc *TemplateConfig) GetGitType() string {
	return tc.Config.GitConfig.Provider
}

func NewTemplateConfig() (*TemplateConfig, error) {
	channels, err := GetChannels()
	if err != nil {
		return nil, err
	}

	config, err := GetConfig()
	if err != nil {
		return nil, err
	}

	promotions, err := GetPromotions()
	if err != nil {
		return nil, err
	}

	return &TemplateConfig{
		Channels:   channels,
		Config:     config,
		Promotions: promotions,
	}, nil
}

func GenerateGitPromotionTemplate() error {
	templateConfig, err := NewTemplateConfig()
	if err != nil {
		return err
	}

	gitProvider := templateConfig.GetGitType()

	path := utils.FilePaths.GetGitDirectoryPath(gitProvider)

	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, f := range files {
		filePath := fmt.Sprintf("%s/%s", path, f.Name())
		fileDir := utils.FilePaths.GetGitDirectoryOutputPath(gitProvider)
		fileName := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
		templ, err := template.ParseFiles(filePath)
		if err != nil {
			return err
		}
		if err := createFileUsingTemplate(templ, fileDir, fileName, templateConfig); err != nil {
			return err
		}
	}

	return nil
}

func createFileUsingTemplate(t *template.Template, dir, filename string, data interface{}) error {
	// create directory if it doesnt exist
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	filePath := fmt.Sprintf("%s/%s", dir, filename)
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

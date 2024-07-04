package v1

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	constantsV1 "github.com/deployix/deployed/pkg/constants/v1"
	templateV1 "github.com/deployix/deployed/pkg/template/v1"
)

func TemplateFuncGithubActionPromotion() error {
	// get the dir path for github actions
	githubDir := constantsV1.DEFAULT_GITHUB_ACTIONS_DIRECTORY_PATH
	templateFilePath := constantsV1.TEMPLATE_GITHUB_ACTION_PROMOTION_FILEPATH

	templateConfig, err := templateV1.NewTemplateConfig()
	if err != nil {
		return err
	}

	// loop through all promotions and create an actions job for each
	for _, promotion := range templateConfig.Promotions.Promotions {
		fileName := fmt.Sprintf("promote-%s.yml", promotion.Name)

		// create directory if it doesnt exist
		if _, err := os.Stat(githubDir); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(githubDir, os.ModePerm)
			if err != nil {
				return err
			}
		}

		filePath := filepath.Join(githubDir, fileName)
		f, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer f.Close()

		template, err := template.ParseFiles(templateFilePath)
		if err != nil {
			return err
		}

		err = template.Execute(f, templateConfig)
		if err != nil {
			return err
		}
	}

	return nil
}

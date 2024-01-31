package deployed

import (
	"fmt"
	"os"
	"text/template"
)

type TemplateConfig struct {
	Config           *Config
	Git              *GitConfig
	PromotionsConfig *PromotionsConfig
}

func NewTemplateConfig() (*TemplateConfig, error) {
	config, err := GetConfig()
	if err != nil {
		return nil, err
	}

	promotions, err := GetPromotions()
	if err != nil {
		return nil, err
	}

	return &TemplateConfig{
		Config:           config,
		PromotionsConfig: promotions,
	}, nil
}

func GenerateGitPromotionTemplate() error {
	templateConfig, err := NewTemplateConfig()
	if err != nil {
		return err
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	path := dir + "/internal/templates/github/post-promotion-notification.yml.tpl"

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return err
	}

	file, err := os.Create(fmt.Sprintf("%v.txt", "git"))
	if err != nil {
		return err
	}

	return tmpl.Execute(file, templateConfig)
}

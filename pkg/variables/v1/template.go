package v1

import (
	"fmt"

	constantsV1 "github.com/deployix/deployed/pkg/constants/v1"
)

type TemplateNameGenerateFunc func() error

// templateNamesToFunc converts a template name into a function that renders the template file
var templateNamesToFunc map[string]TemplateNameGenerateFunc = map[string]TemplateNameGenerateFunc{
	constantsV1.TEMPLATE_NAME_GITHUB_ACTION_PROMOTION: TemplateFuncGithubActionPromotion,
}

func GetTemplateFunc(templateName string) (TemplateNameGenerateFunc, error) {
	if _, found := templateNamesToFunc[templateName]; !found {
		return nil, fmt.Errorf("template name '%s' does not exist", templateName)
	}

	return templateNamesToFunc[templateName], nil
}

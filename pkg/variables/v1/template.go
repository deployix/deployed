package v1

import (
 constantsV1 "github.com/deployix/deployed/pkg/constants/v1"
)

type TemplateNameGenerateFunc func() error

// TemplateNamesToFunc converts a template name into a function that renders the template file
var TemplateNamesToFunc map[string]TemplateNameGenerateFunc = map[string]TemplateNameGenerateFunc{
	constantsV1.TEMPLATE_NAME_GITHUB_ACTION_PROMOTION: 
}

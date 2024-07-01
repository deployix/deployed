package v1

import (
	"context"
	"fmt"
	"os"

	variablesV1 "github.com/deployix/deployed/pkg/variables/v1"
	"github.com/spf13/cobra"
)

var templateGenerateNames []string

func init() {
	templateGenerate.Flags().StringArrayVar(&templateGenerateNames, "names", []string{}, "(required) template names to generate")
	if err := promotionUpdate.MarkFlagRequired("names"); err != nil {
		os.Exit(1)
	}

	template.AddCommand(templateGenerate)
}

var templateGenerate = &cobra.Command{
	Use:          "generate",
	RunE:         TemplateGenerateRun,
	SilenceUsage: true,
}

func TemplateGenerateRun(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	// get config data
	templateConfig, err := NewTemplateConfig()
	if err != nil {
		return err
	}

	// Loop through list of template names to generate
	for _, templateName := range templateGenerateNames {
		if err := TemplateGenerateExecuteGenerateName(ctx, templateName, templateConfig); err != nil {
			return err
		}
	}
	return nil
}

func TemplateGenerateExecuteGenerateName(ctx context.Context, templateName string, data *TemplateConfig) error {
	// Get function that maps to template name
	if _, found := variablesV1.TemplateNamesToFunc[templateName]; !found {
		return fmt.Errorf("template name '%s' does not exist", templateName)
	}

	generateFunc := variablesV1.TemplateNamesToFunc[templateName]

	return generateFunc()
}

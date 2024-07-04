package v1

import (
	"context"
	"fmt"
	"os"

	templateV1 "github.com/deployix/deployed/pkg/template/v1"
	variablesV1 "github.com/deployix/deployed/pkg/variables/v1"
	"github.com/spf13/cobra"
)

var templateGenerateNames []string

func init() {
	templateGenerate.Flags().StringArrayVarP(&templateGenerateNames, "names", "n", []string{}, "(required) template names to generate")
	if err := templateGenerate.MarkFlagRequired("names"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	templates.AddCommand(templateGenerate)
}

var templateGenerate = &cobra.Command{
	Use:          "generate",
	RunE:         TemplateGenerateRun,
	SilenceUsage: true,
}

func TemplateGenerateRun(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	// get config data
	templateConfig, err := templateV1.NewTemplateConfig()
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

func TemplateGenerateExecuteGenerateName(ctx context.Context, templateName string, data *templateV1.TemplateFilesConfig) error {
	// Get function that maps to template name
	generateFunction, err := variablesV1.GetTemplateFunc(templateName)
	if err != nil {
		return err
	}

	return generateFunction()
}

package v1

import (
	"fmt"

	variablesV1 "github.com/deployix/deployed/pkg/variables/v1"
	"github.com/spf13/cobra"
)

func init() {
	template.AddCommand(templateListNames)
}

var templateListNames = &cobra.Command{
	Use:          "list-names",
	Run:          TemplateListNamesRun,
	SilenceUsage: true,
}

func TemplateListNamesRun(cmd *cobra.Command, args []string) {
	for _, name := range variablesV1.GetTemplateNameList() {
		fmt.Println(name)
	}
}

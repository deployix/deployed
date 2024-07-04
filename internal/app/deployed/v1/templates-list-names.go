package v1

import (
	"fmt"

	variablesV1 "github.com/deployix/deployed/pkg/variables/v1"
	"github.com/spf13/cobra"
)

func init() {
	templates.AddCommand(templateListNames)
}

var templateListNames = &cobra.Command{
	Use:          "list",
	Run:          TemplateListRun,
	SilenceUsage: true,
}

func TemplateListRun(cmd *cobra.Command, args []string) {
	for _, name := range variablesV1.GetTemplateNameList() {
		fmt.Println(name)
	}
}

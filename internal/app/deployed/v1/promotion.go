package v1

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(promotion)
}

var promotion = &cobra.Command{
	Use: "promotions",
	Run: promotionsRun,
}

func promotionsRun(cmd *cobra.Command, args []string) {
}

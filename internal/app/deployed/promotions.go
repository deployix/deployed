package deployed

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(promotion)
}

var promotion = &cobra.Command{
	Use:     "promotion",
	Aliases: []string{"promo"},
	Short:   "",
	Long:    "",
	Run:     promotionRun,
}

func promotionRun(cmd *cobra.Command, args []string) {
}

func CreatePromotionsFile() error {
	return nil
}

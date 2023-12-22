package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	promotion.AddCommand(promotionList)
}

var promotionList = &cobra.Command{
	Use:          "list",
	RunE:         PromotionListRun,
	SilenceUsage: true,
}

func PromotionListRun(cmd *cobra.Command, args []string) error {
	// get promotions from file if it exists
	if err := getPromotions(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// list promotions
	for PromotionName, _ := range promos.Promotions {
		fmt.Println(PromotionName)
	}
	return nil
}

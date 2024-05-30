package v1

import (
	"fmt"
	"os"

	promotionsV1 "github.com/deployix/deployed/pkg/promotions/v1"
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
	if err := cmd.ValidateRequiredFlags(); err != nil {
		return err
	}

	// get promotions from file if it exists
	promotions, err := promotionsV1.GetPromotions()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// list promotions
	for PromotionName, _ := range promotions.Promotions {
		fmt.Println(PromotionName)
	}
	return nil
}

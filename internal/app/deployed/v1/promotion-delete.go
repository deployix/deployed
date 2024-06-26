package v1

import (
	"fmt"
	"os"

	promotionsV1 "github.com/deployix/deployed/pkg/promotions/v1"
	"github.com/spf13/cobra"
)

var promotionDeletePromotionName string
var promotionDeleteVerifyPromotion bool

func init() {
	promotionDelete.Flags().StringVarP(&promotionDeletePromotionName, "name", "n", "", "(required) promotion name to delete")
	if err := promotionDelete.MarkFlagRequired("name"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	promotionDelete.Flags().BoolVarP(&promotionDeleteVerifyPromotion, "verify-promotion", "v", false, "verify the promotion exists before deleting, otherwise throw error")

	promotion.AddCommand(promotionDelete)
}

var promotionDelete = &cobra.Command{
	Use:          "delete",
	RunE:         PromotionDeleteRun,
	SilenceUsage: true,
}

func PromotionDeleteRun(cmd *cobra.Command, args []string) error {
	if err := cmd.ValidateRequiredFlags(); err != nil {
		return err
	}

	// get promotions from file if it exists
	promotions, err := promotionsV1.GetPromotions()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if promotionDeleteVerifyPromotion {
		if _, found := promotions.Promotions[promotionDeletePromotionName]; !found {
			return fmt.Errorf("promotion '%s' does not exist", promotionDeletePromotionName)
		}
	}

	// delete promotion
	delete(promotions.Promotions, promotionDeletePromotionName)

	return promotions.WriteToFile()
}

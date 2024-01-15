package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var promotionDeletePromotionName string
var promotionDeleteVerifyPromotion bool

func init() {
	promotionDelete.Flags().StringVarP(&promotionDeletePromotionName, "name", "n", "", "(required) promotion name to delete")
	if err := promotionDelete.MarkFlagRequired("name"); err != nil {
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
	if err := getPromotions(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if promotionDeleteVerifyPromotion {
		if _, found := Promos.Promotions[promotionDeletePromotionName]; !found {
			return fmt.Errorf("promotion '%s' does not exist", promotionDeletePromotionName)
		}
	}

	// delete promotion
	delete(Promos.Promotions, promotionDeletePromotionName)

	// update promotions file
	if err := CreatePromotionsFile(); err != nil {
		return fmt.Errorf("Unable to update promotions config.")
	}

	return nil
}

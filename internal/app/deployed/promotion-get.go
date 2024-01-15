package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var promotionGetPromotionName string

func init() {
	promotionGet.Flags().StringVarP(&promotionGetPromotionName, "name", "n", "", "(required) promotion name to get")
	if err := promotionGet.MarkFlagRequired("name"); err != nil {
		os.Exit(1)
	}

	promotion.AddCommand(promotionGet)
}

var promotionGet = &cobra.Command{
	Use:          "get",
	RunE:         PromotionGetRun,
	SilenceUsage: true,
}

func PromotionGetRun(cmd *cobra.Command, args []string) error {
	if err := cmd.ValidateRequiredFlags(); err != nil {
		return err
	}

	// get promotions from file if it exists
	if err := getPromotions(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if _, found := Promos.Promotions[promotionDeletePromotionName]; !found {
		return fmt.Errorf("promotion '%s' does not exist", promotionGetPromotionName)
	}

	// delete promotion
	fmt.Println(Promos.Promotions[promotionDeletePromotionName])

	// update promotions file
	if err := CreatePromotionsFile(); err != nil {
		return fmt.Errorf("Unable to update promotions config.")
	}

	return nil
}
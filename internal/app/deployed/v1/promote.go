package v1

import (
	"fmt"
	"os"

	promotionsV1 "github.com/deployix/deployed/pkg/promotions/v1"
	"github.com/spf13/cobra"
)

var promotionName string

func init() {
	// create required promotion name flag
	promote.Flags().StringVarP(&promotionName, "name", "n", "", "(required) name of the promotion you want to promote")
	if err := promote.MarkFlagRequired("name"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	rootCmd.AddCommand(promote)
}

var promote = &cobra.Command{
	Use:          "promote",
	RunE:         promoteRun,
	SilenceUsage: true,
}

func promoteRun(cmd *cobra.Command, args []string) error {
	promotions, err := promotionsV1.GetPromotions()
	if err != nil {
		return err
	}

	if !promotions.PromotionExists(promotionName) {
		return fmt.Errorf("promotion with the name `%s` does not exist", promotionName)
	}

	targetedPromotion := promotions.Promotions[promotionName]
	return targetedPromotion.Promote()
}

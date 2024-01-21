package deployed

import (
	"fmt"
	"os"

	"github.com/adhocore/gronx"
	"github.com/spf13/cobra"
)

var promotionUpdatePromotionName string
var promotionUpdateDescription string
var promotionUpdateFromChannel string
var promotionUpdateToChannel string
var promotionUpdateCrontime string

func init() {
	promotionUpdate.Flags().StringVarP(&promotionUpdatePromotionName, "name", "n", "", "(required) promotion name to update")
	if err := promotionUpdate.MarkFlagRequired("name"); err != nil {
		os.Exit(1)
	}

	// Update required promotion fromChannel flag
	promotionUpdate.Flags().StringVarP(&promotionUpdateFromChannel, "from", "f", "", "(required) channel name you want to promote from")

	// Update required promotion toChannel flag
	promotionUpdate.Flags().StringVarP(&promotionUpdateToChannel, "to", "t", "", "(required) channel name you want to promote into")

	// Update required promotion crontime flag
	promotionUpdate.Flags().StringVarP(&promotionUpdateCrontime, "crontime", "c", "", "(required) promotion schedule represented as a crontime string")

	// Update optional promotion description flag
	promotionUpdate.Flags().StringVarP(&promotionUpdateDescription, "desc", "d", "", "promotion description")

	promotion.AddCommand(promotionUpdate)
}

var promotionUpdate = &cobra.Command{
	Use:          "update",
	RunE:         PromotionUpdateRun,
	SilenceUsage: true,
}

func PromotionUpdateRun(cmd *cobra.Command, args []string) error {
	if err := cmd.ValidateRequiredFlags(); err != nil {
		return err
	}

	// get promotions from file if it exists
	promotions, err := GetPromotions()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// populate channels var from channels.yml file if it exists
	channels, err := GetChannels()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if _, found := promotions.Promotions[promotionUpdatePromotionName]; !found {
		return fmt.Errorf("promotion '%s' does not exist", promotionGetPromotionName)
	}

	// update promotion
	updatedPromotion := promotions.Promotions[promotionUpdatePromotionName]

	if promotionUpdateDescription != "" {
		updatedPromotion.Description = promotionUpdateDescription
	}
	if promotionUpdateFromChannel != "" {
		// check from channel exists
		if channels.ChannelExists(promotionUpdateToChannel) {
			return fmt.Errorf("channel %s does not exist", promotionUpdateFromChannel)
		}
		updatedPromotion.FromChannel = promotionUpdateFromChannel
	}
	if promotionUpdateToChannel != "" {
		// check to channel exists
		if channels.ChannelExists(promotionUpdateToChannel) {
			return fmt.Errorf("channel %s does not exist", promotionUpdateToChannel)
		}
		updatedPromotion.ToChannel = promotionUpdateToChannel
	}
	if promotionUpdateCrontime != "" {
		gron := gronx.New()
		if !gron.IsValid(promotionUpdateCrontime) {
			return fmt.Errorf("crontime '%s' is not valid. Valid example '0 */5 * * * *'", promotionUpdateCrontime)
		}
		updatedPromotion.Crontime = promotionUpdateCrontime
	}

	// update promotion
	promotions.Promotions[promotionUpdatePromotionName] = updatedPromotion

	return promotions.WriteToFile()
}

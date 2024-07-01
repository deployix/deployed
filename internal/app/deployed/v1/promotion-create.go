package v1

import (
	"fmt"
	"os"

	"github.com/adhocore/gronx"
	channelsV1 "github.com/deployix/deployed/pkg/channels/v1"
	promotionV1 "github.com/deployix/deployed/pkg/promotion/v1"
	promotionsV1 "github.com/deployix/deployed/pkg/promotions/v1"
	"github.com/spf13/cobra"
)

var promotionCreateName string
var promotionCreateDescription string
var promotionCreateFromChannel string
var promotionCreateToChannel string
var promotionCreateCrontime string

func init() {
	// create required promotion name flag
	promoteCreate.Flags().StringVarP(&promotionCreateName, "name", "n", "", "(required) channel name")
	if err := promoteCreate.MarkFlagRequired("name"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// create required promotion fromChannel flag
	promoteCreate.Flags().StringVarP(&promotionCreateFromChannel, "from", "f", "", "(required) channel name you want to promote from")
	if err := promoteCreate.MarkFlagRequired("from"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// create required promotion toChannel flag
	promoteCreate.Flags().StringVarP(&promotionCreateToChannel, "to", "t", "", "(required) channel name you want to promote into")
	if err := promoteCreate.MarkFlagRequired("to"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// create required promotion crontime flag
	promoteCreate.Flags().StringVarP(&promotionCreateCrontime, "crontime", "c", "", "(required) promotion schedule represented as a crontime string")
	if err := promoteCreate.MarkFlagRequired("crontime"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// create optional promotion description flag
	promoteCreate.Flags().StringVarP(&promotionCreateDescription, "description", "d", "", "promotion description")

	promotion.AddCommand(promoteCreate)
}

var promoteCreate = &cobra.Command{
	Use:          "create",
	RunE:         promotionCreateRun,
	SilenceUsage: true,
}

func promotionCreateRun(cmd *cobra.Command, args []string) error {
	if err := cmd.ValidateRequiredFlags(); err != nil {
		return err
	}

	// get promotions from file if it exists
	promotions, err := promotionsV1.GetPromotions()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// populate channels var from channels.yml file if it exists
	channels, err := channelsV1.GetChannels()
	fmt.Println(channels)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// verify from channels exist
	if !channels.ChannelExists(promotionCreateFromChannel) {
		return fmt.Errorf("channel %s does not exist", promotionCreateFromChannel)
	}

	// check to channel exists
	if !channels.ChannelExists(promotionCreateToChannel) {
		return fmt.Errorf("channel %s does not exist", promotionCreateToChannel)
	}

	// validate crontime
	gron := gronx.New()
	if !gron.IsValid(promotionCreateCrontime) {
		return fmt.Errorf("crontime '%s' is not valid. Valid example '0 */5 * * * *'", promotionCreateCrontime)
	}

	// check if promotions with the same name already exists
	if _, found := promotions.Promotions[promotionCreateName]; found {
		return fmt.Errorf(fmt.Sprintf("promotion with the name '%s' already exists", promotionCreateName))
	}

	// add promotion
	promotions.Promotions[promotionCreateName] = promotionV1.Promotion{
		Name:        promotionCreateName,
		Description: promotionCreateDescription,
		FromChannel: promotionCreateFromChannel,
		ToChannel:   promotionCreateToChannel,
		Crontime:    promotionCreateCrontime,
	}

	return promotions.WriteToFile()
}

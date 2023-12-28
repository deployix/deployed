package deployed

import (
	"fmt"
	"os"

	"github.com/adhocore/gronx"
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
		os.Exit(1)
	}

	// create required promotion fromChannel flag
	promoteCreate.Flags().StringVarP(&promotionCreateFromChannel, "from", "f", "", "(required) channel name you want to promote from")
	if err := promoteCreate.MarkFlagRequired("from"); err != nil {
		os.Exit(1)
	}

	// create required promotion toChannel flag
	promoteCreate.Flags().StringVarP(&promotionCreateToChannel, "to", "t", "", "(required) channel name you want to promote into")
	if err := promoteCreate.MarkFlagRequired("to"); err != nil {
		os.Exit(1)
	}

	// create required promotion crontime flag
	promoteCreate.Flags().StringVarP(&promotionCreateCrontime, "crontime", "c", "", "(required) promotion schedule represented as a crontime string")
	if err := promoteCreate.MarkFlagRequired("crontime"); err != nil {
		os.Exit(1)
	}

	// create optional promotion description flag
	promoteCreate.Flags().StringVarP(&promotionCreateDescription, "desc", "d", "", "promotion description")

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

	// get channels from file if it exists
	if err := getPromotions(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// verify from channels exist
	if !ChannelExists(promotionCreateFromChannel) {
		return fmt.Errorf("channel %s does not exist", promotionCreateFromChannel)
	}

	// check to channel exists
	if !ChannelExists(promotionCreateToChannel) {
		return fmt.Errorf("channel %s does not exist", promotionCreateToChannel)
	}

	// validate crontime
	gron := gronx.New()
	if !gron.IsValid(promotionCreateCrontime) {
		return fmt.Errorf("crontime '%s' is not valid. Valid example '0 */5 * * * *'", promotionCreateCrontime)
	}

	// get promotions from file if it exists
	if err := getPromotions(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// check if promotions with the same name already exists
	if _, found := promos.Promotions[promotionCreateName]; found {
		return fmt.Errorf(fmt.Sprintf("promotion with the name '%s' already exists", promotionCreateName))
	}

	// add promotion
	promos.Promotions[promotionCreateName] = Promotion{
		Name:        promotionCreateName,
		Description: promotionCreateDescription,
		FromChannel: promotionCreateFromChannel,
		ToChannel:   promotionCreateToChannel,
		Crontime:    promotionCreateCrontime,
	}

	// update promotions file
	if err := CreatePromotionsFile(); err != nil {
		return fmt.Errorf("unable to create promotions file. Try running `deployed init` to initialize")
	}

	return nil
}

package deployed

import (
	"os"

	"github.com/deployix/deployed/configs"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func init() {
	rootCmd.AddCommand(promotions)
}

var Promos Promotions

type Promotions struct {
	Promotions map[string]Promotion
}

type Promotion struct {
	Name string `short:"n" long:"name" `
}

var promotions = &cobra.Command{
	Use:     "promotions",
	Aliases: []string{"promos"},
	Short:   "",
	Long:    "",
	Run:     promotionsRun,
}

func promotionsRun(cmd *cobra.Command, args []string) {
}

func CreatePromotionsFile() error {
	promoYmlData, err := yaml.Marshal(&Promos)
	if err != nil {
		return err
	}

	f, err := os.Create(configs.Cfg.GetPromotionsPath())
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(promoYmlData)
	if err != nil {
		return err
	}

	if err = f.Sync(); err != nil {
		return err
	}
	return nil
}

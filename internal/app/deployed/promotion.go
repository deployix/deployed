package deployed

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func init() {
	rootCmd.AddCommand(promotion)

	// try to get initialize promotions if config file exists.
	// if not don't error out
	_ = getPromotions()
}

var promos Promotions

type Promotions struct {
	Promotions map[string]Promotion
}

type Promotion struct {
	Name        string `short:"n" long:"name" yaml:"name"`
	Description string `short:"d" long:"desc" yaml:"description,omitempty"`
	FromChannel string `short:"n" long:"name" yaml:"from_channel"`
	ToChannel   string `short:"n" long:"name" yaml:"to_channel"`
	Crontime    string `short:"c" long:"crontime" yaml:"crontime"`
}

var promotion = &cobra.Command{
	Use: "promotions",
	Run: promotionsRun,
}

func promotionsRun(cmd *cobra.Command, args []string) {
}

func CreatePromotionsFile() error {
	promoYmlData, err := yaml.Marshal(&promos)
	if err != nil {
		return err
	}

	f, err := os.Create(cfg.GetPromotionsPath())
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

func getPromotions() error {
	if _, err := os.Stat(cfg.GetPromotionsPath()); err == nil {
		yamlFile, err := os.ReadFile(cfg.GetPromotionsPath())
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(yamlFile, &promos)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Promotions config file does not exists. Make sure the file %s exists", cfg.GetPromotionsPath())
	}
	return nil
}

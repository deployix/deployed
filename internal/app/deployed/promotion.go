package deployed

import (
	"fmt"
	"os"

	"github.com/deployix/deployed/internal/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func init() {
	rootCmd.AddCommand(promotion)
}

type Promotions struct {
	Promotions map[string]Promotion
}

type Promotion struct {
	Name        string `short:"n" long:"name" yaml:"name"`
	Description string `short:"d" long:"desc" yaml:"description,omitempty"`
	FromChannel string `short:"f" long:"from_channel" yaml:"from_channel"`
	ToChannel   string `short:"t" long:"to_channel" yaml:"to_channel"`
	Crontime    string `short:"c" long:"crontime" yaml:"crontime"`
}

var promotion = &cobra.Command{
	Use: "promotions",
	Run: promotionsRun,
}

func (p *Promotions) WriteToFile() error {
	promotionsYmlData, err := yaml.Marshal(&p)
	if err != nil {
		return err
	}

	f, err := os.Create(utils.FilePaths.GetPromotionsFilePath())
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(promotionsYmlData)
	if err != nil {
		return err
	}

	if err = f.Sync(); err != nil {
		return err
	}
	return nil
}

func promotionsRun(cmd *cobra.Command, args []string) {
}

func GetPromotions() (*Promotions, error) {
	if _, err := os.Stat(utils.FilePaths.GetPromotionsFilePath()); err == nil {
		promotionsConfigFile := &Promotions{}
		yamlFile, err := os.ReadFile(utils.FilePaths.GetPromotionsFilePath())
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(yamlFile, promotionsConfigFile)
		if err != nil {
			return nil, err
		}
		return promotionsConfigFile, nil
	}
	return nil, fmt.Errorf("Channels config file does not exists. Make sure the file %s exists", utils.FilePaths.GetChannelsFilePath())
}

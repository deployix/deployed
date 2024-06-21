package v1

import (
	"fmt"
	"os"

	promotionV1 "github.com/deployix/deployed/pkg/promotion/v1"
	utilsV1 "github.com/deployix/deployed/pkg/utils/v1"
	"gopkg.in/yaml.v3"
)

type Promotions struct {
	Promotions map[string]promotionV1.Promotion
}

// PromotionExists validates a promotion exists and returns true is found otherwise returns false
func (p *Promotions) PromotionExists(name string) bool {
	if _, found := p.Promotions[name]; found {
		return true
	}
	return false
}

func (p *Promotions) WriteToFile() error {
	promotionsYmlData, err := yaml.Marshal(&p)
	if err != nil {
		return err
	}

	f, err := os.Create(utilsV1.FilePaths().GetPromotionsFilePath())
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

func GetPromotions() (*Promotions, error) {
	if _, err := os.Stat(utilsV1.FilePaths().GetPromotionsFilePath()); err == nil {
		promotionsConfigFile := &Promotions{}
		yamlFile, err := os.ReadFile(utilsV1.FilePaths().GetPromotionsFilePath())
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(yamlFile, promotionsConfigFile)
		if err != nil {
			return nil, err
		}
		return promotionsConfigFile, nil
	}
	return nil, fmt.Errorf("Promotions config file does not exists. Make sure the file %s exists", utilsV1.FilePaths().GetPromotionsFilePath())
}

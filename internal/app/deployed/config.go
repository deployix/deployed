package deployed

import (
	"log"
	"os"

	"github.com/deployix/deployed/configs"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func init() {
	rootCmd.AddCommand(config)

	// populate cfg var with whats in .deployed/config.yml file
	if _, err := os.Stat(configs.Cfg.GetConfigPath()); err == nil {
		cfgYmlFile, err := os.ReadFile(configs.Cfg.GetConfigPath())
		if err != nil {
			log.Printf("yamlFile.Get err   #%v ", err)
		}
		err = yaml.Unmarshal(cfgYmlFile, configs.Cfg)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	}
}

var config = &cobra.Command{
	Use:     "config",
	Aliases: []string{"cfg"},
	Short:   "",
	Long:    "",
	Run:     configRun,
}

func configRun(cmd *cobra.Command, args []string) {
}

func CreateConfigFile() error {
	configYmlData, err := yaml.Marshal(&configs.Cfg)
	if err != nil {
		return err
	}

	f, err := os.Create(configs.Cfg.GetConfigPath())
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(configYmlData)
	if err != nil {
		return err
	}

	if err = f.Sync(); err != nil {
		return err
	}
	return nil
}

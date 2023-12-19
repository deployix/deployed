package deployed

import (
	"os"

	"github.com/deployix/deployed/configs"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var Ver Versions

type Versions struct {
	Versions map[string]Version
}

type Version struct {
}

func init() {
	rootCmd.AddCommand(versions)
}

var versions = &cobra.Command{
	Use:     "versions",
	Aliases: []string{"v"},
	Short:   "",
	Long:    "",
	Run:     versionsRun,
}

func versionsRun(cmd *cobra.Command, args []string) {
}

func CreateVersionsFile() error {
	versionsYmlData, err := yaml.Marshal(&Ver)
	if err != nil {
		return err
	}

	f, err := os.Create(configs.Cfg.GetVersionsPath())
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(versionsYmlData)
	if err != nil {
		return err
	}

	if err = f.Sync(); err != nil {
		return err
	}
	return nil
}

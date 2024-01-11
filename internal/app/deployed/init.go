package deployed

import (
	"fmt"
	"os"

	"github.com/deployix/deployed/internal/constants"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().BoolP("force", "f", false, "Force initialization")
}

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{""},
	Short:   "",
	Long:    "",
	RunE:    initRun,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	},
}

func initRun(cmd *cobra.Command, args []string) error {
	// check if .deployed directory already exists [Done]
	// if it does check if "force" flag is set other wise throw error indicating dir exists. [Done]
	// create {config,channels,promotions,versions}.yml files
	// create flags to pass in possible values users may want to include later (i.e. channels etc)

	force, _ := cmd.Flags().GetBool("force")

	if err := generateWorkingDir(force); err != nil {
		return err
	}

	return nil

}

func generateWorkingDir(force bool) error {
	if _, err := os.Stat(cfg.GetDirectoryPath()); err == nil && !force {
		// Dir exists and we are not forcing the creation
		return fmt.Errorf("dir %s already exists. Use --force to overwrite", cfg.GetDirectoryPath())
	} else {
		err := os.RemoveAll(cfg.GetDirectoryPath())
		if err != nil {
			return err
		}
	}

	if err := os.Mkdir(cfg.GetDirectoryPath(), constants.DEFAULT_DIR_FILEMODE); err != nil {
		return err
	}

	// TODO: create a struct that had default values to use when init vs var config
	// use func createInitConfig()
	if err := CreateConfigFile(); err != nil {
		return err
	}

	if err := CreateChannelsFile(); err != nil {
		return err
	}

	if err := CreatePromotionsFile(); err != nil {
		return err
	}

	return nil
}

func createInitConfig() Config {
	return Config{}
}

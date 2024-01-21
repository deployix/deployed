package deployed

import (
	"fmt"
	"os"

	"github.com/deployix/deployed/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configSetDatetimeFormat string

func init() {
	// configSet.Flags().StringVarP(&configSetDatetimeFormat, "name", "n", "", "(required) config name to update")
	// if err := configSet.MarkFlagRequired("name"); err != nil {
	// 	os.Exit(1)
	// }

	config.AddCommand(configSet)
}

var configSet = &cobra.Command{
	Use:               "set",
	RunE:              configSetRun,
	Args:              cobra.ExactArgs(2),
	ValidArgsFunction: validateConfigSetArgs,
	SilenceUsage:      true,
}

func configSetRun(cmd *cobra.Command, args []string) error {
	if err := cmd.ValidateRequiredFlags(); err != nil {
		return err
	}

	config, err := GetConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	key := args[0]
	value := utils.ConvertStringToType(args[1])

	// check if key exists in config
	if !viper.InConfig(key) {
		fmt.Printf(fmt.Sprintf("key '%s' does not exist in config", key))
		os.Exit(1)
	}

	// set key
	viper.Set(key, value)

	// write to file
	return config.WriteToFile()
}

// validateConfigSetArgs validates the args sent
func validateConfigSetArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var comps []string

	return comps, cobra.ShellCompDirectiveNoFileComp
}

package deployed

import (
	"fmt"
	"os"

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
	// get configs from file if it exists
	if err := getConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := cmd.ValidateRequiredFlags(); err != nil {
		return err
	}

	key := args[0]
	value := args[1] // TODO: convert string to proper type (i.e. int,bool etc)

	// check if key exists in config
	if !viper.InConfig(key) {
		fmt.Printf(fmt.Sprintf("key '%s' does not exist in config", key))
		os.Exit(1)
	}

	// set key
	if err := setConfigKey(key, value); err != nil {
		return err
	}

	// update struct before updating file
	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Printf("error Unmarshal config %v", err)
	}

	// update config file
	return CreateConfigFile()
}

// set config key
func setConfigKey(key string, value interface{}) error {
	switch value.(type) {
	case string:
		fmt.Println("string")
		viper.Set(key, value.(string))
		break
	case bool:
		fmt.Println("bool")
		viper.Set(key, value.(bool))
		break
	case int:
		fmt.Println("int")
		viper.Set(key, value.(int))
		break
	case int16:
		fmt.Println("int16")
		viper.Set(key, value.(int16))
		break
	case int32:
		fmt.Println("int32")
		viper.Set(key, value.(int32))
		break
	case int64:
		fmt.Println("int64")
		viper.Set(key, value.(int64))
		break
	case float32:
		fmt.Println("float32")
		viper.Set(key, value.(float32))
		break
	case float64:
		fmt.Println("float64")
		viper.Set(key, value.(float64))
		break
	default:
		return fmt.Errorf(fmt.Sprintf("Unsupported type: '%v", value))
	}
	return nil
}

// validateConfigSetArgs validates the args sent
func validateConfigSetArgs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var comps []string

	return comps, cobra.ShellCompDirectiveNoFileComp
}

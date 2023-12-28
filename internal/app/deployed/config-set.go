package deployed

import (
	"fmt"
	"os"

	"github.com/adhocore/gronx"
	"github.com/spf13/cobra"
)

var configSetDatetimeFormat string

func init() {
	configSet.Flags().StringVarP(&configSetDatetimeFormat, "name", "n", "", "(required) config name to update")
	if err := configSet.MarkFlagRequired("name"); err != nil {
		os.Exit(1)
	}

	config.AddCommand(configSet)
}

var configSet = &cobra.Command{
	Use:          "set",
	RunE:         configSetRun,
	SilenceUsage: true,
}

func configSetRun(cmd *cobra.Command, args []string) error {
	if err := cmd.ValidateRequiredFlags(); err != nil {
		return err
	}

	// get configs from file if it exists
	if err := getConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if _, found := promos.configs[configUpdateconfigName]; !found {
		return fmt.Errorf("config '%s' does not exist", configGetconfigName)
	}

	// update config
	updatedconfig := promos.configs[configUpdateconfigName]

	if configUpdateDescription != "" {
		updatedconfig.Description = configUpdateDescription
	}
	if configUpdateFromChannel != "" {
		// check from channel exists
		if !ChannelExists(configUpdateToChannel) {
			return fmt.Errorf("channel %s does not exist", configUpdateFromChannel)
		}
		updatedconfig.FromChannel = configUpdateFromChannel
	}
	if configUpdateToChannel != "" {
		// check to channel exists
		if !ChannelExists(configUpdateToChannel) {
			return fmt.Errorf("channel %s does not exist", configUpdateToChannel)
		}
		updatedconfig.ToChannel = configUpdateToChannel
	}
	if configUpdateCrontime != "" {
		gron := gronx.New()
		if !gron.IsValid(configUpdateCrontime) {
			return fmt.Errorf("crontime '%s' is not valid. Valid example '0 */5 * * * *'", configUpdateCrontime)
		}
		updatedconfig.Crontime = configUpdateCrontime
	}

	// update config
	promos.configs[configUpdateconfigName] = updatedconfig

	// update configs file
	if err := CreateconfigsFile(); err != nil {
		return fmt.Errorf("Unable to update configs config.")
	}

	return nil
}

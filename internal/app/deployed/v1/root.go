package v1

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	viper.WatchConfig()

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.deployed.yml)")
}

var rootCmd = &cobra.Command{
	Use:   "deployed",
	Short: "",
	Long:  "",
}

func initConfig() {

}

func Execute() error {
	return rootCmd.Execute()
}

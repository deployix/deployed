package deployed

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
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
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// set default config file path
		viper.SetConfigFile(cfg.GetConfigPath())
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("using config file:", viper.ConfigFileUsed())
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Printf("error Unmarshal config %v", err)
	}
}

func OnConfigChanged(e fsnotify.Event) {
	fmt.Println("Config file changed:", e.Name)

}

func Execute() error {
	return rootCmd.Execute()
}

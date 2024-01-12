package configs

import "github.com/spf13/viper"

// ConfigFiles manages all the viper config files that makeup deployed
type ConfigFiles struct {
	ChannelsFile   viper.Viper
	ConfigFile     viper.Viper
	PromotionsFile viper.Viper
}

func init() {

}

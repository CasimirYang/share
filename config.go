package share

import (
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("config")
	viper.SetConfigName("application") // name of config file (without extension)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		SugarLogger.Errorf("read config failed: %v", err)
		//os.Exit(1)
	}
}

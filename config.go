package share

import (
	"github.com/spf13/viper"
	"os"
)

func init() {
	viper.AddConfigPath("config")
	viper.SetConfigName("application") // name of config file (without extension)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		SugarLogger.Error("read config failed: %v", err)
		os.Exit(1)
	}
}
package config

import "github.com/spf13/viper"

// 바이퍼
func InitConfig() {
	viper.SetConfigName("./config/config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

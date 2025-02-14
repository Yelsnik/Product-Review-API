package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBSource          string `mapstructure:"DB_SOURCE"`
	GRPCServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`
	RapidAPIHost      string `mapstructure:"RAPID_API_HOST"`
	RapidAPIKey       string `mapstructure:"RAPID_API_KEY"`
	RapidAPISearchUrl string `mapstructure:"RAPID_API_SEARCH_URL"`
	RapidAPIDetailsUrl string `mapstructure:"RAPID_API_DETAILS_URL"`
}


func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

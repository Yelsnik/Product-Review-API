package util

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DBSource           string `mapstructure:"DB_SOURCE"`
	GRPCServerAddress  string `mapstructure:"GRPC_SERVER_ADDRESS"`
	RapidAPIHost       string `mapstructure:"RAPID_API_HOST"`
	RapidAPIKey        string `mapstructure:"RAPID_API_KEY"`
	RapidAPISearchUrl  string `mapstructure:"RAPID_API_SEARCH_URL"`
	RapidAPIDetailsUrl string `mapstructure:"RAPID_API_DETAILS_URL"`
	Redis              string `mapstructure:"REDIS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	viper.BindEnv("DBSource", "DB_SOURCE")
	viper.BindEnv("GRPCServerAddress", "GRPC_SERVER_ADDRESS")
	viper.BindEnv("RapidAPIHost", "RAPID_API_HOST")
	viper.BindEnv("RapidAPIKey", "RAPID_API_KEY")
	viper.BindEnv("RapidAPISearchUrl", "RAPID_API_SEARCH_URL")
	viper.BindEnv("RapidAPIDetailsUrl", "RAPID_API_DETAILS_URL")
	viper.BindEnv("Redis", "REDIS")

	err = viper.ReadInConfig()
	if err != nil && !os.IsNotExist(err) {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

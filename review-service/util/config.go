package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBSource           string `mapstructure:"DB_SOURCE"`
	MigrationURL       string `mapstructure:"MIGRATION_URL"`
	GRPCServerAddress  string `mapstructure:"GRPC_SERVER_ADDRESS"`
	RapidAPIHost       string `mapstructure:"RAPID_API_HOST"`
	RapidAPIKey        string `mapstructure:"RAPID_API_KEY"`
	RapidAPISearchUrl  string `mapstructure:"RAPID_API_SEARCH_URL"`
	RapidAPIDetailsUrl string `mapstructure:"RAPID_API_DETAILS_URL"`
	Redis              string `mapstructure:"REDIS"`
	RedisAddress       string `mapstructure:"REDIS_ADDRESS"`
	RedisPassword      string `mapstructure:"REDIS_PASSWORD"`
	RedisDB            string `mapstructure:"REDIS_DB"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return
		}
		// return
	}

	viper.AutomaticEnv()
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// dbsource := viper.GetString("DB_SOURCE")
	redis := viper.GetString("REDIS_ADDRESS")
	// fmt.Println(dbsource)
	fmt.Println(redis)

	err = viper.Unmarshal(&config)
	return
}

package config

import (
	"os"

	"github.com/spf13/viper"
)

type Conf struct {
	WeatherToken string `mapstructure:"WEATHER_TOKEN"`
}

func LoadConfig(path string) (*Conf, error) {
	if os.Getenv("ENVIRONMENT") == "PROD" {
		return &Conf{
			WeatherToken: os.Getenv("WEATHER_TOKEN"),
		}, nil
	}
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var cfg *Conf
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, nil
}

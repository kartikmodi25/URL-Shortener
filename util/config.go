package util

import "github.com/spf13/viper"

type Config struct {
	DB_USERNAME   string `mapstructure:"DB_USERNAME"`
	DB_PASSWORD   string `mapstructure:"DB_PASSWORD"`
	DB_HOST       string `mapstructure:"DB_HOST"`
	DB_PORT       string `mapstructure:"DB_PORT"`
	DB_NAME       string `mapstructure:"DB_NAME"`
	PORT          string `mapstructure:"PORT"`
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

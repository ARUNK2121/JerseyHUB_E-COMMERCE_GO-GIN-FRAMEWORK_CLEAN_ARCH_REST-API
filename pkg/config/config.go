package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	BASE_URL              string `mapstructure:"BASE_URL"`
	DBHost                string `mapstructure:"DB_HOST"`
	DBName                string `mapstructure:"DB_NAME"`
	DBUser                string `mapstructure:"DB_USER"`
	DBPort                string `mapstructure:"DB_PORT"`
	DBPassword            string `mapstructure:"DB_PASSWORD"`
	AUTHTOKEN             string `mapstructure:"DB_AUTHTOKEN"`
	ACCOUNTSID            string `mapstructure:"DB_ACCOUNTSID"`
	SERVICESID            string `mapstructure:"DB_SERVICESID"`
	AWS_REGION            string `mapstructure:"AWS_REGION"`
	AWS_ACCESS_KEY_ID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWS_SECRET_ACCESS_KEY string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
}

var envs = []string{
	"BASE_URL", "DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "DB_AUTHTOKEN", "DB_ACCOUNTSID", "DB_SERVICESID", "AWS_REGION", "AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}

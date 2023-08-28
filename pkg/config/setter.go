package config

import (
	"blog/config"
	"log"

	"github.com/spf13/viper"
)

func Set() config.Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf(":::::Error Reading Configs::::: %s", err)
	}

	err := viper.Unmarshal(&configurations)

	if err != nil {
		log.Fatalf("unable to decode into struct %s", err)
	}

	return configurations
}

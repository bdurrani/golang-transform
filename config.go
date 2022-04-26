package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port string
}

func GetConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		_ = fmt.Errorf("error reading config file, %s", err)
		return nil, err
	}
	port := viper.GetString("server.port")
	return &Config{
		Port: port,
	}, nil
}

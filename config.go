package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Broker   string
	Version  uint
	Username string
	Password string
	Topic    Topic
	Ca       Ca
}

type Topic struct {
	Sub string
	Pub string
}

type Ca struct {
	Enable   bool
	Cafile   string
	Certfile string
	keyfile  string
}

func LoadConfig() Config {
	var c Config
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("read config failed", err)
	}
	if err := v.Unmarshal(&c); err != nil {
		fmt.Println("unable to decode into struct, %v", err)
	}
	return c
}

package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	DB_Username  string
	DB_Password  string
	DB_Port      string
	DB_Host      string
	DB_Name      string
	API_PORT     string
	TOKEN_SECRET string
}

var Cfg *Config

func InitConfig() {
	cfg := &Config{}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
	}

	err := viper.Unmarshal(cfg)
	if err != nil {
		fmt.Println(err.Error())
	}

	cfg.TOKEN_SECRET = "AbCd3F9H1"

	Cfg = cfg
}

func Env(key, value string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		os.Setenv(key, value)
	} else {
		value = val
	}

	return value
}

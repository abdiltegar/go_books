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

	viper.Unmarshal(cfg)
	//// baca env
	//cfg.DB_Username = Env("DB_USERNAME", "root")
	//cfg.DB_Password = Env("DB_PASSWORD", "secret123")
	//cfg.DB_Port = Env("DB_PORT", "3306")
	//cfg.DB_Host = Env("DB_HOST", "localhost")
	//cfg.DB_Name = Env("DB_NAME", "learn-go")
	//cfg.API_PORT = Env("API_PORT", ":8080")
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

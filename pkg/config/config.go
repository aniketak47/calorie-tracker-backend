package config

import (
	"errors"
	"fmt"
	"log"
	"path"
	"runtime"
	"strings"

	"github.com/iamolegga/enviper"
	"github.com/spf13/viper"
)

func getPwd() string {
	_, b, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(path.Dir(path.Dir(b))))
}

func GetConfig() (Config, error) {
	var cfg Config

	e := enviper.New(viper.New())

	e.AddConfigPath(getPwd())
	e.SetConfigName(".env")
	e.SetConfigType("env")

	e.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	e.AutomaticEnv()

	if err := e.ReadInConfig(); err != nil {
		log.Fatal(err)
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal("Error reading config file: ", err)
		}
	}

	err := e.Unmarshal(&cfg)
	if err != nil {
		fmt.Printf("error to decode, %s", err)
		return cfg, nil
	}

	if cfg.Server.Port == "" {
		return cfg, errors.New("error reading config")
	}

	return cfg, nil
}

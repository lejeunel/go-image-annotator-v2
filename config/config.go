package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	DBPath string `required:"true"`
}

func Parse() Config {
	var cfg Config
	err := envconfig.Process("GOIA", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return cfg
}

package config

import (
	"github.com/andresxlp/gosuite/config"
	"log"
	"sync"
)

type Config struct {
	AppPort int `default:"8082" mapstructure:"app_port"`
}

var (
	Cfg  Config
	Once sync.Once
)

func Environments() *Config {
	Once.Do(func() {
		if err := config.GetConfigFromEnv(&Cfg); err != nil {
			log.Fatal(err)
		}
	})

	return &Cfg
}

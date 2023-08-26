package configFromFile

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config3 struct {
	Host string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port int    `yaml:"port" env:"PORT" env-default:"8080"`
}

func GetFileConfig() *Config3 {
	var once sync.Once
	cfg := &Config3{}
	once.Do(func() {
		if err := cleanenv.ReadConfig("configFromFile/config.yml", cfg); err != nil {
			help, _ := cleanenv.GetDescription(cfg, nil)
			log.Println(help)
			log.Println(err)
		}
	})
	return cfg
}

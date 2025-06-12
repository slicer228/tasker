package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http"`
	Tasker     `yaml:"tasker"`
}

type HTTPServer struct {
	Address string        `yaml:"address" env-default:"localhost:8080"`
	Timeout time.Duration `yaml:"timeout" env-default:"5s"`
}

type Tasker struct {
	MaxTasks uint64 `yaml:"maxTasks" env-default:"10"`
}

func MustLoad() *Config {
	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatal("CONFIG_PATH file not exists")
	}

	cfg := &Config{}

	if err := cleanenv.ReadConfig(cfgPath, cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}

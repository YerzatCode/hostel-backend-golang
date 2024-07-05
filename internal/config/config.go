package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Listen   `yaml:"listen"`
	Database `yaml:"database"`
}

type Listen struct {
	Port string `yaml:"port" env-default:":8080"`
	Host string `yaml:"host" env-default:"localhost"`
	Mode string `yaml:"mod" env-default:"dev"`
}

type Database struct {
	Path string `yaml:"path" env-required:"true"`
	DB   string `yaml:"db"`
}

func ConfigPathInit() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current directory:", err)
	}
	configPath := filepath.Join(currentDir, "/config.yaml")
	os.Setenv("CONFIG_PATH", configPath)
}

func InitConfig() *Config {
	ConfigPathInit()
	configPath := os.Getenv("CONFIG_PATH")

	var cfg Config

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file not found ", configPath)
	}
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("cannat read config file ", err)
	}
	return &cfg
}

package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	env        string     `yaml:"env"`
	HttpServer HttpServer `yaml:"http-server"`
}

type HttpServer struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

func MustLoadConfig() *Config {
	configPath := "./config/config.yaml"

	if configPath == "" {
		panic("config path is empty")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic(err)
	}

	return &cfg
}

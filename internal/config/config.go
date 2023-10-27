package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	GatewayConfig struct {
		App   `yaml:"app"`
		HTTP  `yaml:"http"`
		REDIS `yaml:"redis"`
	}

	StorageConfig struct {
		App   `yaml:"app"`
		HTTP  `yaml:"http"`
		PG    `yaml:"postgres"`
		REDIS `yaml:"redis"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name" env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	REDIS struct {
		Channel  string `env-required:"true" yaml:"channel" env:"REDIS_CHANNEL"`
		Addr     string `env-required:"true" yaml:"address" env:"REDIS_ADDRESS"`
		Password string `yaml:"password" env:"REDIS_PASSWORD"`
		DB       int    `yaml:"db" env:"REDIS_DB"`
	}

	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env-required:"true" yaml:"connect_url" env:"PG_CONNECT_URL"`
	}
)

func NewGateway() (*GatewayConfig, error) {
	cfg := &GatewayConfig{}

	err := cleanenv.ReadConfig("./config/gateway.cfg.yml", cfg)

	log.Print(fmt.Errorf("config error: %w", err))

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func NewStorage() (*StorageConfig, error) {
	cfg := &StorageConfig{}

	err := cleanenv.ReadConfig("./config/gateway.cfg.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

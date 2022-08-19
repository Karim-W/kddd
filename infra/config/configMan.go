package config

import (
	"os"
)

type ConfigManger interface {
	GetConfig() AppConfig
}

type configMangerImpl struct {
	c AppConfig
}

var _ ConfigManger = (*configMangerImpl)(nil)

func ConfigMangerProvider() ConfigManger {
	c := AppConfig{
		Database: DatabaseConfig{
			DSN:    os.Getenv("DATABASE_DSN"),
			Driver: os.Getenv("DATABASE_DRIVER"),
		},
	}
	return &configMangerImpl{
		c: c,
	}
}

func (c *configMangerImpl) GetConfig() AppConfig {
	return c.c
}

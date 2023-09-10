package config

import (
	"os"
	"www.github.com/Wolves-from-MISIS/internal/models"
)

type Config struct {
	DB          *DBConfig
	ServerAddr  string `json:"server_addr"`
	SwaggerAddr string `json:"swagger_addr"`
}

type DBConfig struct {
	Dsn string
}

func ReadConfig() (*Config, error) {
	var cfg Config

	db := &DBConfig{}
	cfg.DB = db
	v, ok := os.LookupEnv("DB_DSN")
	if !ok {
		return nil, models.NotAllEnvErr
	}
	cfg.DB.Dsn = v
	v, ok = os.LookupEnv("SERVER_ADDR")
	if !ok {
		return nil, models.NotAllEnvErr
	}
	cfg.ServerAddr = v
	v, ok = os.LookupEnv("SWAGGER_ADDR")
	if !ok {
		return nil, models.NotAllEnvErr
	}
	cfg.SwaggerAddr = v

	return &cfg, nil
}

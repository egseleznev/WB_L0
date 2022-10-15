package service

import "l0/service/store"

type Config struct {
	ClusterID string
	ClientID  string
	LogLevel  string
	Database  *store.Config
}

func NewConfig() *Config {
	return &Config{
		ClusterID: "prod",
		ClientID:  "subscriber",
		LogLevel:  "debug",
		Database:  store.NewConfig(),
	}
}

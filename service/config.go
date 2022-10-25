package service

import (
	"l0/service/store"
)

// Config represents the configuration of nats streaming server client
// ClusterID is cluster unique identifier to connect
// ClientID is client unique identifier to connect
// LogLevel is a logging level
// Database is a handle to database configuration
type Config struct {
	ClusterID string
	ClientID  string
	LogLevel  string
	Database  *store.Config
}

// NewConfig returns the default config used in this application to create a stream client
func NewConfig() *Config {

	return &Config{
		ClusterID: "prod",
		ClientID:  "subscriber",
		LogLevel:  "debug",
		Database:  store.NewConfig(),
	}
}

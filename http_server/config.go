package http_server

// Config represents the configuration of http_server
// URL is server connection address
type Config struct {
	Addr string
}

// NewConfig returns the default config used in this application to connect to http_server
func NewConfig() *Config {

	return &Config{
		Addr: ":5000",
	}
}

package store

type Config struct {
	URL string
}

func NewConfig() *Config {
	return &Config{}
}

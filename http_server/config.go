package http_server

type Config struct {
	Addr string
}

func NewConfig() *Config {
	return &Config{
		Addr: ":5000",
	}
}

package store

type Config struct {
	URL string
}

func NewConfig() *Config {
	return &Config{
		URL: "host=localhost port=5432 user=WB password=Qwe123!!! dbname=WB_L0 sslmode=disable",
	}
}

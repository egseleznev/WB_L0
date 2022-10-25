package store

// Config represents the configuration of database
// URL is database connection address(host, port, user, password, dbname is required)
type Config struct {
	URL string
}

// NewConfig returns the default config used in this application to connect to database
func NewConfig() *Config {

	return &Config{
		URL: "host=127.0.0.1 port=5432 user=WB password=Qwe123!!! dbname=WB_L0 sslmode=disable",
	}
}

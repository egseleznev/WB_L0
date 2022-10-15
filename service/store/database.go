package store

type Database struct {
	config *Config
}

func New(config *Config) *Database {
	return &Config{
		config: config,
	}
}

func (d *Database) Connect() error {
	return nil
}

func (d *Database) Disconnect() {
}

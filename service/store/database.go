package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Database struct {
	config     *Config
	db         *sql.DB
	repository *Repository
}

func New(config *Config) *Database {
	return &Database{
		config: config,
	}
}

func (d *Database) Connect() error {
	db, err := sql.Open("postgres", d.config.URL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	d.db = db
	return nil
}

func (d *Database) Disconnect() {
	d.db.Close()
}

func (d *Database) Order() *Repository {
	if d.repository != nil {
		return d.repository
	}
	d.repository = &Repository{
		database: d,
	}
	return d.repository
}

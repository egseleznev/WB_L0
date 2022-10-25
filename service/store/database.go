package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// Database represents entity for interacting with the database
// config is a handle to the specified database config
// db is a database handle representing a connection
// repository is a handle to interacting with stored data
type Database struct {
	config     *Config
	db         *sql.DB
	repository *Repository
}

// New creates and return a database with the specified config
func New(config *Config) *Database {

	return &Database{
		config: config,
	}
}

// Connect establishes a connection to the database and sends a ping request to test the connection
// return nil if the connection is successful
// otherwise return error
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

// Disconnect close an established database connection
func (d *Database) Disconnect() {

	d.db.Close()
}

// Order return current repository handle if already exists
// otherwise create and return it instead
func (d *Database) Order() *Repository {

	if d.repository != nil {
		return d.repository
	}

	d.repository = &Repository{
		database: d,
	}

	return d.repository
}

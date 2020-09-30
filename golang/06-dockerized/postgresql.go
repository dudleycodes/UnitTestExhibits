// Package postgresql is a toy package for use in demonstrating how Docker containers can be used to automate some Go
// integration tests.
package postgresql

import (
	"database/sql"
	"fmt"

	// Add in the "pq" PostgreSQL driver
	_ "github.com/lib/pq"
)

// Config determines the configuration for the PostgreSQL Broker.
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// Broker encapsulates
type Broker struct {
	database *sql.DB
}

// New creates a new connection to a remote PostgreSQL database.
func New(cfg Config) (*Broker, error) {
	conn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName)

	fmt.Printf("\n%s\n", conn)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, fmt.Errorf("Couldn't create PostgreSQL connection: %w", err)
	}

	return &Broker{database: db}, nil
}

// Close the connection to the remote PostgreSQL database.
func (bkr *Broker) Close() error {
	return bkr.database.Close()
}

// RowCount counts the number of rows for a specified PostgreSQL table.
func (bkr *Broker) RowCount(tableName string) (int, error) {
	row := bkr.database.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName))

	var count int

	if err := row.Scan(&count); err != nil {
		return 0, fmt.Errorf("Failed to count rows for table %s: %w", tableName, err)
	}

	return count, nil
}

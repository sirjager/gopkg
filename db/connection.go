package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Config struct {
	Migrations  string `mapstructure:"MIGRATIONS_DIR"` // The path to the database migration files
	PostgresURL string `mapstructure:"POSTGRES_URL"`   // The URL for the database connection

	RedisURL string `mapstructure:"REDIS_URL"` // The URL for the redis connction
}

// Database represents a database connection.
type Database struct {
	logr   zerolog.Logger
	pool   *pgxpool.Pool
	config Config
}

// NewConnection creates a new database connection based on the provided configuration.
// It returns a Database instance and any error encountered during the connection process.
func NewDatabae(
	ctx context.Context,
	config Config,
	logr zerolog.Logger,
) (*Database, *pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, config.PostgresURL)
	if err != nil {
		return nil, nil, err
	}
	// Return a new Database instance with the connection.
	return &Database{logr, pool, config}, pool, err
}

// Close closes the database connection.
func (database *Database) Close() {
	database.pool.Close()
}

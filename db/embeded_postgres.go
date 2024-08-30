package db

import (
	"context"
	"time"

	embPostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmbededPostgresOpts struct {
	DBName     string
	DBUser     string
	DBPass     string
	RunTimeDir string
	DBPort     uint32
	Timeout    time.Duration
}

type EmbededPostgres struct {
	Pool          *pgxpool.Pool
	pg            *embPostgres.EmbeddedPostgres
	ConnectionURL string
}

// NewEmbededPostgres can be used for running embedded postgres instance for testing and demo purposes.
func NewEmbededPostgres(opts *EmbededPostgresOpts) *EmbededPostgres {
	config := embPostgres.DefaultConfig().
		Username(opts.DBUser).Database(opts.DBName).Password(opts.DBPass).
		Port(opts.DBPort).
		RuntimePath(opts.RunTimeDir).
		StartTimeout(opts.Timeout)
	connectionURL := config.GetConnectionURL() + "?sslmode=disable"
	pg := embPostgres.NewDatabase(config)
	return &EmbededPostgres{pg: pg, ConnectionURL: connectionURL}
}

func (e *EmbededPostgres) Start(ctx context.Context) error {
	pool, err := pgxpool.New(ctx, e.ConnectionURL)
	if err != nil {
		return err
	}
	e.Pool = pool
	return e.pg.Start()
}

func (e *EmbededPostgres) Stop() error {
	e.Pool.Close()
	return e.pg.Stop()
}

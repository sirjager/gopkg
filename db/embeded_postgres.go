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
	Timeout    time.Duration
}

type EmbededPostgres struct {
	Pool          *pgxpool.Pool
	pg            *embPostgres.EmbeddedPostgres
	ConnectionURL string
}

// NewEmbededPostgres can be used for running embedded postgres instance for testing and demo purposes.
func NewEmbededPostgres(ctx context.Context, opts *EmbededPostgresOpts) (*EmbededPostgres, error) {
	pgConfig := embPostgres.DefaultConfig().
		Username(opts.DBUser).Database(opts.DBName).Password(opts.DBPass).
		RuntimePath(opts.RunTimeDir).
		StartTimeout(opts.Timeout)

	connectionURL := pgConfig.GetConnectionURL() + "?sslmode=disable"

	pool, err := pgxpool.New(ctx, connectionURL)
	if err != nil {
		return nil, err
	}

	pg := embPostgres.NewDatabase(pgConfig)
	return &EmbededPostgres{
		pg:            pg,
		Pool:          pool,
		ConnectionURL: connectionURL,
	}, nil
}

func (e *EmbededPostgres) Start() error {
	return e.pg.Start()
}

func (e *EmbededPostgres) Stop() error {
	e.Pool.Close()
	return e.pg.Stop()
}

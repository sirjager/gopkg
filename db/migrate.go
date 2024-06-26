package db

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	_ "github.com/lib/pq"
)

// MigrateUsingBindata applies database migrations using the `go-bindata` package as the source driver.
// It creates a migration source using the binary data generated by the `migrations` package, and then
// creates a new migrate instance with the `go-bindata` source driver and the provided database URL. It
// then applies any pending migrations to the database and logs the result.
// 
// Usage:
// import "github.com/username/repo/migrations"  // path to generated go-bindata
//
//	if err = database.MigrateUsingBindata(migrations.AssetNames(), migrations.Asset); err != nil {
//		logr.Fatal().Err(err).Msg("failed to migrate database")
//	}
//
// For Reference: https://github.com/sirjager/trueauth/blob/master/main.go
func (d *Database) MigrateUsingBindata(
	migrationsAssetNames []string,
	migrationAsset func(name string) ([]byte, error),
) (err error) {
	// migration binary data
	migrationSource := bindata.Resource(migrationsAssetNames,
		func(name string) ([]byte, error) {
			return migrationAsset(name)
		})

	sourceDriver, err := bindata.WithInstance(migrationSource)
	if err != nil {
		d.logr.Fatal().Err(err).Msg("failed to create gobindata source driver instance")
	}

	dbmigrate, err := migrate.NewWithSourceInstance("go-bindata", sourceDriver, d.config.PostgresURL)
	if err != nil {
		return err
	}
	err = dbmigrate.Up()
	if err != nil {
		// Check if the error is "no change" which indicates that there are no pending migrations.
		// Log an info message in this case.
		if err != migrate.ErrNoChange {
			return err
		}
		d.logr.Info().Msg("database migration is up to date")
	} else {
		d.logr.Info().Msg("database migration complete")
	}

	return nil
}

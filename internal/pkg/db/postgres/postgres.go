package postgres

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/johejo/golang-migrate-extra/source/iofs"
	_ "github.com/lib/pq" // need for init PostgreSQL interface
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/db/options"
)

var (
	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar) // nolint unused

	//go:embed migrations/*.sql
	migrations embed.FS
)

// Init ...
func (p *Store) Init(ctx context.Context) error {
	var err error

	// Set configuration
	p.setConfig()

	// Apply migration
	err = p.migrate()
	if err != nil {
		return err
	}

	// Connect to Postgres
	if p.client, err = pgxpool.Connect(ctx, p.config.URI); err != nil {
		return err
	}

	return nil
}

// GetConn ...
func (s *Store) GetConn() interface{} {
	return s.client
}

// Close ...
func (p *Store) Close() error { // nolint unparam
	p.client.Close()
	return nil
}

// Migrate ...
func (p *Store) migrate() error { // nolint unused
	// Create connect
	db, err := sql.Open("postgres", p.config.URI)
	if err != nil {
		return err
	}

	driver, err := iofs.New(migrations, "migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithSourceInstance("iofs", driver, p.config.URI)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}

// setConfig - set configuration
func (p *Store) setConfig() {
	dbinfo := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", "postgres", "shortlink", "shortlink")

	viper.AutomaticEnv()
	viper.SetDefault("STORE_POSTGRES_URI", dbinfo)                  // Postgres URI
	viper.SetDefault("STORE_MODE_WRITE", options.MODE_SINGLE_WRITE) // mode write to db

	p.config = Config{
		URI:  viper.GetString("STORE_POSTGRES_URI"),
		mode: viper.GetInt("STORE_MODE_WRITE"),
	}
}

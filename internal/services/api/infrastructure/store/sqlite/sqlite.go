package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3" // Init SQLite-driver

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/services/api/domain/link"
	"github.com/batazor/shortlink/internal/services/api/infrastructure/store/query"
)

// Store implementation of db interface
type Store struct { // nolint unused
	client *sql.DB
}

// Init ...
func (_ *Store) Init(_ context.Context, _ *db.Store) error {
	return nil
}

// Get ...
func (lite *Store) Get(ctx context.Context, id string) (*link.Link, error) {
	// query builder
	links := squirrel.Select("url, hash, describe").
		From("links").
		Where(squirrel.Eq{"hash": id})
	query, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	stmt, err := lite.client.Prepare(query)
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}
	defer stmt.Close() // nolint errcheck

	var response link.Link
	err = stmt.QueryRow(args...).Scan(&response.Url, &response.Hash, &response.Describe)
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (lite *Store) List(_ context.Context, _ *query.Filter) ([]*link.Link, error) {
	// query builder
	links := squirrel.Select("url, hash, describe").
		From("links")
	query, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := lite.client.Query(query, args...)
	if err != nil || rows.Err() != nil {
		return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
	}
	defer rows.Close() // nolint errcheck

	var response []*link.Link

	for rows.Next() {
		var result link.Link
		err = rows.Scan(&result.Url, &result.Hash, &result.Describe)
		if err != nil {
			return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
		}

		response = append(response, &result)
	}

	return response, nil
}

// Add ...
func (lite *Store) Add(_ context.Context, source *link.Link) (*link.Link, error) {
	err := link.NewURL(source)
	if err != nil {
		return nil, err
	}

	// query builder
	links := squirrel.Insert("links").
		Columns("url", "hash", "describe").
		Values(source.Url, source.Hash, source.Describe)

	query, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = lite.client.Exec(query, args...)
	if err != nil {
		return nil, &link.NotFoundError{Link: source, Err: fmt.Errorf("Failed save link: %s", source.Url)}
	}

	return source, nil
}

// Update ...
func (lite *Store) Update(_ context.Context, _ *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (lite *Store) Delete(ctx context.Context, id string) error {
	// query builder
	links := squirrel.Delete("links").
		Where(squirrel.Eq{"hash": id})
	query, args, err := links.ToSql()
	if err != nil {
		return err
	}

	_, err = lite.client.Exec(query, args...)
	if err != nil {
		return &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Failed delete link: %s", id)}
	}

	return nil
}

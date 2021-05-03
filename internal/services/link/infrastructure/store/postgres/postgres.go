//go:generate protoc -I../../../../../services/link/domain/link --gotemplate_out=all=true,template_dir=template:. link.proto
package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq" // need for init PostgreSQL interface
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/batch"
	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/db/options"
	"github.com/batazor/shortlink/internal/services/link/domain/link"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/query"
)

var (
	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar) // nolint unused
)

// Init ...
func (s *Store) Init(ctx context.Context, db *db.Store) error {
	// Set configuration
	s.setConfig()
	s.client = db.Store.GetConn().(*pgxpool.Pool)

	// Create batch job
	if s.config.mode == options.MODE_BATCH_WRITE {
		cb := func(args []*batch.Item) interface{} {
			sources := make([]*link.Link, len(args))

			for key := range args {
				sources[key] = args[key].Item.(*link.Link)
			}

			dataList, errBatchWrite := s.batchWrite(ctx, sources)
			if errBatchWrite != nil {
				for index := range args {
					// TODO: add logs for error
					args[index].CB <- errors.New("Error write to PostgreSQL")
				}
				return errBatchWrite
			}

			for key, item := range dataList.Link {
				args[key].CB <- item
			}

			return nil
		}

		var err error
		s.config.job, err = batch.New(ctx, cb)
		if err != nil {
			return err
		}

		go s.config.job.Run(ctx)
	}

	return nil
}

// Get ...
func (p *Store) Get(ctx context.Context, id string) (*link.Link, error) {
	// query builder
	links := psql.Select("url, hash, describe").
		From("links").
		Where(squirrel.Eq{"hash": id})
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.client.Query(ctx, q, args...)

	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}
	if rows.Err() != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response link.Link
	for rows.Next() {
		err = rows.Scan(&response.Url, &response.Hash, &response.Describe)
		if err != nil {
			return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
		}
	}

	if response.Hash == "" {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (p *Store) List(ctx context.Context, filter *query.Filter) (*link.Links, error) {
	// query builder
	links := psql.Select("url, hash, describe, created_at, updated_at").
		From("links")
	links = p.buildFilter(links, filter)
	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.client.Query(ctx, q, args...)
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
	}

	response := &link.Links{
		Link: []*link.Link{},
	}

	for rows.Next() {
		var result link.Link
		var (
			created_ad sql.NullTime
			updated_at sql.NullTime
		)
		err = rows.Scan(&result.Url, &result.Hash, &result.Describe, &created_ad, &updated_at)
		if err != nil {
			return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
		}
		result.CreatedAt = &timestamp.Timestamp{Seconds: int64(created_ad.Time.Second()), Nanos: int32(created_ad.Time.Nanosecond())}
		result.UpdatedAt = &timestamp.Timestamp{Seconds: int64(updated_at.Time.Second()), Nanos: int32(updated_at.Time.Nanosecond())}

		response.Link = append(response.Link, &result)
	}

	return response, nil
}

// Add ...
func (p *Store) Add(ctx context.Context, source *link.Link) (*link.Link, error) {
	switch p.config.mode {
	case options.MODE_BATCH_WRITE:
		cb, err := p.config.job.Push(source)
		if err != nil {
			return nil, err
		}

		res := <-cb
		switch data := res.(type) {
		case error:
			return nil, data
		case *link.Link:
			return data, nil
		default:
			return nil, nil
		}
	case options.MODE_SINGLE_WRITE:
		data, err := p.singleWrite(ctx, source)
		return data, err
	}

	return nil, nil
}

// Update ...
func (p *Store) Update(_ context.Context, _ *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (p *Store) Delete(ctx context.Context, id string) error {
	// query builder
	links := psql.Delete("links").
		Where(squirrel.Eq{"hash": id})
	q, args, err := links.ToSql()
	if err != nil {
		return err
	}

	_, err = p.client.Exec(ctx, q, args...)
	if err != nil {
		return &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Failed delete link: %s", id)}
	}

	return nil
}

func (p *Store) singleWrite(ctx context.Context, source *link.Link) (*link.Link, error) {
	err := link.NewURL(source)
	if err != nil {
		return nil, err
	}

	// save as JSON. it doesn't make sense
	dataJson, err := json.Marshal(source)
	if err != nil {
		return nil, err
	}

	// query builder
	links := psql.Insert("links").
		Columns("url", "hash", "describe", "json").
		Values(source.Url, source.Hash, source.Describe, dataJson)

	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	row := p.client.QueryRow(ctx, q, args...)

	errScan := row.Scan(&source.Url, &source.Hash, &source.Describe).Error()
	if errScan == "no rows in result set" {
		return source, nil
	}
	if errScan != "" {
		return nil, &link.NotFoundError{Link: source, Err: fmt.Errorf("Failed save link: %s", source.Url)}
	}

	return source, nil
}

func (p *Store) batchWrite(ctx context.Context, sources []*link.Link) (*link.Links, error) {
	// Create a new link
	for key := range sources {
		err := link.NewURL(sources[key])
		if err != nil {
			return nil, err
		}
	}

	links := psql.Insert("links").Columns("url", "hash", "describe", "json")

	// query builder
	for _, source := range sources {
		// save as JSON. it doesn't make sense
		dataJson, err := json.Marshal(source)
		if err != nil {
			return nil, err
		}

		links = links.Values(source.Url, source.Hash, source.Describe, dataJson)
	}

	q, args, err := links.ToSql()
	if err != nil {
		return nil, err
	}

	row := p.client.QueryRow(ctx, q, args...)
	errScan := row.Scan(&sources)
	if errScan.Error() == "no rows in result set" {
		return &link.Links{
			Link: []*link.Link{},
		}, nil
	}
	if errScan != nil {
		return nil, fmt.Errorf("Error save link")
	}

	response := &link.Links{
		Link: []*link.Link{},
	}

	for item := range sources {
		response.Link = append(response.Link, sources[item])
	}

	return response, nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MODE_WRITE", options.MODE_SINGLE_WRITE) // mode write to db. Select: 0 (MODE_SINGLE_WRITE), 1 (MODE_BATCH_WRITE)

	s.config = Config{
		mode: viper.GetInt("STORE_MODE_WRITE"),
	}
}

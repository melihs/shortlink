package leveldb

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/services/link/domain/link"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/query"
)

// Store implementation of db interface
type Store struct { // nolint unused
	client *leveldb.DB
}

// Init ...
func (s *Store) Init(_ context.Context, db *db.Store) error {
	s.client = db.Store.GetConn().(*leveldb.DB)
	return nil
}

// Add ...
func (l *Store) Add(_ context.Context, source *link.Link) (*link.Link, error) {
	err := link.NewURL(source)
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(source)
	if err != nil {
		return nil, err
	}

	err = l.client.Put([]byte(source.Hash), payload, nil)
	if err != nil {
		return nil, err
	}

	return source, nil
}

// Get ...
func (l *Store) Get(ctx context.Context, id string) (*link.Link, error) {
	value, err := l.client.Get([]byte(id), nil)
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response link.Link

	err = json.Unmarshal(value, &response)
	if err != nil {
		return nil, err
	}

	if response.Url == "" {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (l *Store) List(_ context.Context, _ *query.Filter) (*link.Links, error) {
	links := &link.Links{
		Link: []*link.Link{},
	}
	iterator := l.client.NewIterator(nil, nil)

	for iterator.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		value := iterator.Value()

		var response link.Link

		err := json.Unmarshal(value, &response)
		if err != nil {
			return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
		}

		links.Link = append(links.Link, &response)
	}

	iterator.Release()
	err := iterator.Error()
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
	}

	return links, nil
}

// Update ...
func (l *Store) Update(_ context.Context, _ *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (l *Store) Delete(ctx context.Context, id string) error {
	err := l.client.Delete([]byte(id), nil)
	return err
}

package redis

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/query"
)

// Store implementation of db interface
type Store struct {
	client redis.UniversalClient
}

// New store
func New(ctx context.Context, db *db.Store) (*Store, error) {
	s := &Store{
		client: db.Store.GetConn().(redis.UniversalClient),
	}

	return s, nil
}

// Get ...
func (s *Store) Get(ctx context.Context, id string) (*v1.Link, error) {
	val, err := s.client.Get(ctx, id).Result()
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response v1.Link
	if err = protojson.Unmarshal([]byte(val), &response); err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Failed parse link: %s", id)}
	}

	return &response, nil
}

// List ...
func (s *Store) List(ctx context.Context, filter *query.Filter) (*v1.Links, error) {
	keys := s.client.Keys(ctx, "*")
	links := &v1.Links{
		Link: []*v1.Link{},
	}

	for _, key := range keys.Val() {
		var response v1.Link
		val, err := s.client.Get(ctx, key).Result()
		if err != nil {
			return nil, &v1.NotFoundError{Link: &v1.Link{}, Err: fmt.Errorf("Not found links")}
		}

		if err = protojson.Unmarshal([]byte(val), &response); err != nil {
			return nil, &v1.NotFoundError{Link: &v1.Link{}, Err: fmt.Errorf("Not found links")}
		}

		links.Link = append(links.Link, &response)
	}

	return links, nil
}

// Add ...
func (s *Store) Add(ctx context.Context, source *v1.Link) (*v1.Link, error) {
	err := v1.NewURL(source)
	if err != nil {
		return nil, err
	}

	val, err := protojson.Marshal(source)
	if err != nil {
		return nil, &v1.NotFoundError{Link: source, Err: fmt.Errorf("Failed marsharing link: %s", source.Url)}
	}

	if err = s.client.Set(ctx, source.Hash, val, 0).Err(); err != nil {
		return nil, &v1.NotFoundError{Link: source, Err: fmt.Errorf("Failed save link: %s", source.Url)}
	}

	return source, nil
}

// Update ...
func (s *Store) Update(_ context.Context, _ *v1.Link) (*v1.Link, error) {
	return nil, nil
}

// Delete ...
func (s *Store) Delete(ctx context.Context, id string) error {
	if err := s.client.Del(ctx, id).Err(); err != nil {
		return &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("Failed save link: %s", id)}
	}

	return nil
}

package badger

import (
	"context"
	"fmt"

	"github.com/dgraph-io/badger/v4"
	"google.golang.org/protobuf/encoding/protojson"

	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/query"
)

// Store implementation of db interface
type Store struct {
	client *badger.DB
}

// New store
func New(_ context.Context) (*Store, error) {
	return &Store{}, nil
}

// Get ...
func (b *Store) Get(ctx context.Context, id string) (*v1.Link, error) {
	var valCopy []byte

	err := b.client.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			// Copying or parsing val is valid.
			valCopy = append([]byte{}, val...)

			return nil
		})
		if err != nil {
			return err
		}

		// Alternatively, you could also use item.ValueCopy().
		valCopy, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("not found id: %s", id)}
	}

	var response v1.Link

	err = protojson.Unmarshal(valCopy, &response)
	if err != nil {
		return nil, err
	}

	if response.Url == "" {
		return nil, &v1.NotFoundError{Link: &v1.Link{Hash: id}, Err: fmt.Errorf("not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (b *Store) List(_ context.Context, _ *query.Filter) (*v1.Links, error) {
	var list [][]byte

	err := b.client.View(func(txn *badger.Txn) error {
		iterator := txn.NewIterator(badger.DefaultIteratorOptions)
		defer iterator.Close()

		for iterator.Rewind(); iterator.Valid(); iterator.Next() {
			var valCopy []byte
			item := iterator.Item()

			err := item.Value(func(val []byte) error {
				// Copying or parsing val is valid.
				valCopy = append([]byte{}, val...)

				return nil
			})
			if err != nil {
				return err
			}

			// Alternatively, you could also use item.ValueCopy().
			valCopy, err = item.ValueCopy(nil)
			if err != nil {
				return err
			}

			list = append(list, valCopy)
		}

		return nil
	})
	if err != nil {
		return nil, &v1.NotFoundError{Link: &v1.Link{}, Err: fmt.Errorf("not found links: %w", err)}
	}

	response := &v1.Links{
		Link: []*v1.Link{},
	}

	for _, item := range list {
		l := &v1.Link{}
		err = protojson.Unmarshal(item, l)
		if err != nil {
			return nil, err
		}

		response.Link = append(response.Link, l)
	}

	return response, nil
}

// Add ...
func (b *Store) Add(ctx context.Context, source *v1.Link) (*v1.Link, error) {
	err := v1.NewURL(source)
	if err != nil {
		return nil, err
	}

	payload, err := protojson.Marshal(source)
	if err != nil {
		return nil, err
	}

	err = b.client.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(source.Hash), payload)
	})
	if err != nil {
		return nil, err
	}

	return source, nil
}

// Update ...
func (b *Store) Update(_ context.Context, _ *v1.Link) (*v1.Link, error) {
	return nil, nil
}

// Delete ...
func (b *Store) Delete(ctx context.Context, id string) error {
	err := b.client.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(id))
		return err
	})

	return err
}

package badger

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/store/query"
	"github.com/batazor/shortlink/pkg/link"
	"github.com/dgraph-io/badger"
)

// BadgerConfig ...
type BadgerConfig struct { // nolint unused
	Path string
}

// BadgerLinkList implementation of store interface
type BadgerLinkList struct { // nolint unused
	client *badger.DB
	config BadgerConfig
}

// Init ...
func (b *BadgerLinkList) Init() error {
	var err error

	// Set configuration
	b.setConfig()

	b.client, err = badger.Open(badger.DefaultOptions(b.config.Path))
	if err != nil {
		return err
	}

	return nil
}

// Close ...
func (b *BadgerLinkList) Close() error {
	return b.client.Close()
}

// Migrate ...
func (b *BadgerLinkList) migrate() error { // nolint unused
	return nil
}

// Get ...
func (b *BadgerLinkList) Get(id string) (*link.Link, error) {
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
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: fmt.Errorf("not found id: %s", id)}
	}

	var response link.Link

	err = json.Unmarshal(valCopy, &response)
	if err != nil {
		return nil, err
	}

	if response.Url == "" {
		return nil, &link.NotFoundError{Link: link.Link{Url: id}, Err: fmt.Errorf("not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (b *BadgerLinkList) List(filter *query.Filter) ([]*link.Link, error) { // nolint unused
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
		return nil, &link.NotFoundError{Link: link.Link{}, Err: fmt.Errorf("not found links: %s", err)}
	}

	response := make([]*link.Link, len(list))

	for index, link := range list {
		err = json.Unmarshal(link, &response[index])
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

// Add ...
func (b *BadgerLinkList) Add(data link.Link) (*link.Link, error) {
	hash := data.CreateHash([]byte(data.Url), []byte("secret"))
	data.Hash = hash[:7]

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = b.client.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(data.Hash), payload)
	})
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// Update ...
func (b *BadgerLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (b *BadgerLinkList) Delete(id string) error {
	err := b.client.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(id))
		return err
	})
	return err
}

// setConfig - set configuration
func (b *BadgerLinkList) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_BADGER_PATH", "/tmp/links.badger")

	b.config = BadgerConfig{
		Path: viper.GetString("STORE_BADGER_PATH"),
	}
}

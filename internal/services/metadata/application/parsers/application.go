/*
Metadata Service. Application layer
*/
package parsers

import (
	"context"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"

	http_client "github.com/shortlink-org/shortlink/internal/pkg/http/client"
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
	v1 "github.com/shortlink-org/shortlink/internal/services/metadata/domain/metadata/v1"
	meta_store "github.com/shortlink-org/shortlink/internal/services/metadata/infrastructure/repository"
)

type Service struct {
	MetaStore *meta_store.MetaStore
}

func New(store *meta_store.MetaStore) (*Service, error) {
	return &Service{
		MetaStore: store,
	}, nil
}

func (r *Service) Get(ctx context.Context, hash string) (*v1.Meta, error) {
	meta, err := r.MetaStore.Store.Get(ctx, hash)
	if err != nil {
		return nil, err
	}

	return meta, nil
}

func (r *Service) Set(ctx context.Context, url string) (*v1.Meta, error) {
	meta := &v1.Meta{
		Id: url,
	}

	newCtx, cancel := context.WithTimeout(ctx, time.Minute*1)
	defer cancel()

	// Request the HTML page.
	req, err := http.NewRequestWithContext(newCtx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	client := http_client.New()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); name == "description" {
			meta.Description, _ = s.Attr("content")
		}

		if name, _ := s.Attr("name"); name == "keyworlds" {
			meta.Keywords, _ = s.Attr("content")
		}
	})

	// Write to DB
	err = r.MetaStore.Store.Add(newCtx, meta)
	if err != nil {
		return nil, err
	}

	// publish event by this service
	notify.Publish(newCtx, v1.METHOD_ADD, meta, nil)

	return meta, nil
}

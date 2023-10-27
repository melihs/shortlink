//go:generate go run entgo.io/ent/cmd/ent generate --feature sql/upsert ./ent/schema

package mysql

import (
	"context"
	"database/sql"
	"embed"
	"errors"

	"github.com/google/uuid"
	"github.com/segmentio/encoding/json"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	"github.com/shortlink-org/shortlink/internal/pkg/db/mysql/migrate"
	domain "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/mysql/schema/crud"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/query"
)

var (
	//go:embed migrations/*.sql
	migrations embed.FS
)

// New store
func New(ctx context.Context, store db.DB) (*Store, error) {
	client, ok := store.GetConn().(*sql.DB)
	if !ok {
		return nil, errors.New("error get connection to MySQL")
	}

	s := &Store{
		client: crud.New(client),
	}

	// run migration
	err := migrate.Migration(ctx, store, migrations, "link")
	if err != nil {
		return nil, err
	}

	return s, nil
}

// TODO: use uuid.FromString(id) for all id
func (s Store) Get(ctx context.Context, hash string) (*domain.Link, error) {
	link, err := s.client.GetLinkByHash(ctx, hash)
	if err != nil {
		return nil, err
	}

	var payload domain.Link
	err = json.Unmarshal(link.Json, &payload)
	if err != nil {
		return nil, err
	}

	return &domain.Link{
		Url:       link.Url,
		Hash:      link.Hash,
		Describe:  link.Describe.String,
		CreatedAt: timestamppb.New(link.CreatedAt),
		UpdatedAt: timestamppb.New(link.UpdatedAt),
	}, nil
}

func (s Store) List(ctx context.Context, filter *query.Filter) (*domain.Links, error) {
	links, err := s.client.GetLinks(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*domain.Link, 0, len(links))
	for _, link := range links {
		resp = append(resp, &domain.Link{
			Url:       link.Url,
			Hash:      link.Hash,
			Describe:  link.Describe.String,
			CreatedAt: timestamppb.New(link.CreatedAt),
			UpdatedAt: timestamppb.New(link.UpdatedAt),
		})
	}

	return &domain.Links{
		Link: resp,
	}, nil
}

func (s Store) Add(ctx context.Context, in *domain.Link) (*domain.Link, error) {
	// create uuid
	id := uuid.New()

	payload, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	_, err = s.client.CreateLink(ctx, crud.CreateLinkParams{
		ID:       id.String(),
		Url:      in.Url,
		Hash:     in.Hash,
		Describe: sql.NullString{String: in.Describe, Valid: true},
		Json:     payload,
	})
	if err != nil {
		return nil, err
	}

	return in, nil
}

func (s Store) Update(ctx context.Context, in *domain.Link) (*domain.Link, error) {
	payload, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	_, err = s.client.UpdateLink(ctx, crud.UpdateLinkParams{
		Url:      in.Url,
		Hash:     in.Hash,
		Describe: sql.NullString{String: in.Describe, Valid: true},
		Json:     payload,
	})
	if err != nil {
		return nil, err
	}

	return in, nil
}

func (s Store) Delete(ctx context.Context, id string) error {
	err := s.client.DeleteLink(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
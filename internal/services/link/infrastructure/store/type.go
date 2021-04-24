package store

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/link/domain/link"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/query"
)

type Repository interface {
	Init(ctx context.Context, db *db.Store) error
	Get(ctx context.Context, id string) (*link.Link, error)
	List(ctx context.Context, filter *query.Filter) ([]*link.Link, error)
	Add(ctx context.Context, data *link.Link) (*link.Link, error)
	Update(ctx context.Context, data *link.Link) (*link.Link, error)
	Delete(ctx context.Context, id string) error
}

// Store abstract type
type LinkStore struct { // nolint unused
	typeStore string
	Store     Repository

	// system event
	notify.Subscriber // Observer interface for subscribe on system event
}

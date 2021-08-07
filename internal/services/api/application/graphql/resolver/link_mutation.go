package resolver

import (
	"context"
	"fmt"

	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/api/domain"
	"github.com/batazor/shortlink/internal/services/link/domain/link/v1"
)

// CreateLink ...
func (r *Resolver) CreateLink(ctx context.Context, args *struct { //nolint unused
	URL      *string
	Hash     *string
	Describe *string
}) (*LinkResolver, error) {
	newLink := &v1.Link{
		Url:      *args.URL,
		Hash:     *args.Hash,
		Describe: *args.Describe,
	}

	responseCh := make(chan interface{})

	go notify.Publish(ctx, api_domain.METHOD_ADD, newLink, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_ADD"})

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_ADD")
		return &LinkResolver{
			Link: nil,
		}, err
	case notify.Response:
		err := r.Error
		if err != nil {
			return nil, err
		}
		response := r.Payload.(*v1.Link) // nolint errcheck
		return &LinkResolver{
			Link: response,
		}, err
	default:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_ADD")
		return &LinkResolver{
			Link: nil,
		}, err
	}
}

// UpdateLink ...
func (*Resolver) UpdateLink(ctx context.Context, args *struct { //nolint unused
	URL      *string
	Hash     *string
	Describe *string
}) (*bool, error) {
	return nil, nil
}

// DeleteLink ...
func (r *Resolver) DeleteLink(ctx context.Context, args *struct { //nolint unused
	Hash *string
}) (bool, error) {
	responseCh := make(chan interface{})

	go notify.Publish(ctx, api_domain.METHOD_DELETE, *args.Hash, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_DELETE"})

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_DELETE")
		return false, err
	case notify.Response:
		err := r.Error
		if err != nil {
			return false, err
		}
		return true, nil
	default:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_DELETE")
		return false, err
	}
}

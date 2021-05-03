/*
Metadata Service. Application layer
*/
package link_application

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/link/domain/link"
	link_store "github.com/batazor/shortlink/internal/services/link/infrastructure/store"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/query"
	metadata_rpc "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc"
	"github.com/batazor/shortlink/pkg/saga"
)

type Service struct {
	// Delivery
	MetadataClient metadata_rpc.MetadataClient

	// Repository
	*link_store.LinkStore

	logger logger.Logger
}

func New(logger logger.Logger, metadataService metadata_rpc.MetadataClient, linkStore *link_store.LinkStore) (*Service, error) {
	return &Service{
		MetadataClient: metadataService,
		LinkStore:      linkStore,
		logger:         logger,
	}, nil
}

func errorHelper(ctx context.Context, logger logger.Logger, errs []error) error {
	if len(errs) > 0 {
		errList := field.Fields{}
		for index := range errs {
			errList[fmt.Sprintf("stack error: %d", index)] = errs[index]
		}

		logger.ErrorWithContext(ctx, "Error create a new link", errList)
		return fmt.Errorf("Error create a new link")
	}

	return nil
}

func (s *Service) AddLink(ctx context.Context, in *link.Link) (*link.Link, error) {
	const (
		SAGA_NAME                        = "ADD_LINK"
		SAGA_STEP_STORE_SAVE             = "SAGA_STEP_STORE_SAVE"
		SAGA_STEP_METADATA_GET           = "SAGA_STEP_METADATA_GET"
		SAGA_STEP_PUBLISH_EVENT_NEW_LINK = "SAGA_STEP_PUBLISH_EVENT_NEW_LINK"
	)

	// create a new saga for create a new link
	sagaAddLink, errs := saga.New(SAGA_NAME, saga.Logger(s.logger)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// add step: save to Store
	_, errs = sagaAddLink.AddStep(SAGA_STEP_STORE_SAVE).
		Then(func(ctx context.Context) error {
			var err error
			in, err = s.Store.Add(ctx, in)
			return err
		}).
		Reject(func(ctx context.Context) error {
			return s.Store.Delete(ctx, in.Hash)
		}).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// add step: request to metadata
	_, errs = sagaAddLink.AddStep(SAGA_STEP_METADATA_GET).
		Then(func(ctx context.Context) error {
			_, err := s.MetadataClient.Set(ctx, &metadata_rpc.SetMetaRequest{
				Id: in.Url,
			})
			return err
		}).
		Reject(func(ctx context.Context) error {
			return s.Store.Delete(ctx, in.Hash)
		}).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// add step: publish event by this service
	_, errs = sagaAddLink.AddStep(SAGA_STEP_PUBLISH_EVENT_NEW_LINK).
		Then(func(ctx context.Context) error {
			notify.Publish(ctx, uint32(link.LinkEvent_ADD), in, nil)
			return nil
		}).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaAddLink.Play(nil)
	if err != nil {
		return nil, err
	}

	return in, nil
}

func (s *Service) GetLink(ctx context.Context, hash string) (*link.Link, error) {
	const (
		SAGA_NAME           = "GET_LINK"
		SAGA_STEP_STORE_GET = "SAGA_STEP_STORE_GET"
	)

	resp := &link.Link{}

	// create a new saga for get link by hash
	sagaGetLink, errs := saga.New(SAGA_NAME, saga.Logger(s.logger)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// add step: get link from store
	_, errs = sagaGetLink.AddStep(SAGA_STEP_STORE_GET).
		Then(func(ctx context.Context) error {
			var err error
			resp, err = s.Store.Get(ctx, hash)
			return err
		}).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaGetLink.Play(nil)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, &link.NotFoundError{Link: &link.Link{Hash: hash}, Err: fmt.Errorf("Not found links")}
	}

	return resp, nil
}

func (s *Service) ListLink(ctx context.Context, in string) (*link.Links, error) {
	// Parse args
	filter := &query.Filter{}

	if in != "" {
		errJsonUnmarshal := json.Unmarshal([]byte(in), &filter)
		if errJsonUnmarshal != nil {
			return nil, errors.New("error parse payload as string")
		}
	}

	const (
		SAGA_NAME            = "LIST_LINK"
		SAGA_STEP_STORE_LIST = "SAGA_STEP_STORE_LIST"
	)

	resp := &link.Links{}

	// create a new saga for get list of link
	sagaListLink, errs := saga.New(SAGA_NAME, saga.Logger(s.logger)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// add step: get link from store
	_, errs = sagaListLink.AddStep(SAGA_STEP_STORE_LIST).
		Then(func(ctx context.Context) error {
			var err error
			resp, err = s.Store.List(ctx, filter)
			return err
		}).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaListLink.Play(nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Service) UpdateLink(ctx context.Context, in *link.Link) (*link.Link, error) {
	return nil, nil
}

func (s *Service) DeleteLink(ctx context.Context, hash string) (*link.Link, error) {
	const (
		SAGA_NAME              = "DELETE_LINK"
		SAGA_STEP_STORE_DELETE = "SAGA_STEP_STORE_DELETE"
	)

	// create a new saga for delete link by hash
	sagaDeleteLink, errs := saga.New(SAGA_NAME, saga.Logger(s.logger)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// add step: get link from store
	_, errs = sagaDeleteLink.AddStep(SAGA_STEP_STORE_DELETE).
		Then(func(ctx context.Context) error {
			return s.Store.Delete(ctx, hash)
		}).
		Build()
	if err := errorHelper(ctx, s.logger, errs); err != nil {
		return nil, err
	}

	// Run saga
	err := sagaDeleteLink.Play(nil)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

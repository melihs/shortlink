// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package di

import (
	"context"
	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq/v1"
	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/link/application/link"
	"github.com/batazor/shortlink/internal/services/link/application/link_cqrs"
	"github.com/batazor/shortlink/internal/services/link/application/sitemap"
	v1_5 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/mq"
	v1_3 "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	v1_2 "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/run"
	v1_4 "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/sitemap/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/cqrs/cqs"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/cqrs/query"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud"
	v1_6 "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc/metadata/v1"
	"github.com/batazor/shortlink/pkg/rpc"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

// Injectors from di.go:

func InitializeLinkService(ctx context.Context, runRPCClient *grpc.ClientConn, runRPCServer *rpc.RPCServer, log logger.Logger, db2 *db.Store, mq v1.MQ) (*LinkService, func(), error) {
	metadataServiceClient, err := NewMetadataRPCClient(runRPCClient)
	if err != nil {
		return nil, nil, err
	}
	store, err := NewLinkStore(ctx, log, db2)
	if err != nil {
		return nil, nil, err
	}
	service, err := NewLinkApplication(log, metadataServiceClient, store)
	if err != nil {
		return nil, nil, err
	}
	cqsStore, err := NewCQSLinkStore(ctx, log, db2)
	if err != nil {
		return nil, nil, err
	}
	queryStore, err := NewQueryLinkStore(ctx, log, db2)
	if err != nil {
		return nil, nil, err
	}
	link_cqrsService, err := NewLinkCQRSApplication(log, cqsStore, queryStore)
	if err != nil {
		return nil, nil, err
	}
	sitemapService, err := NewSitemapApplication(log, mq)
	if err != nil {
		return nil, nil, err
	}
	event, err := InitLinkMQ(ctx, log, mq)
	if err != nil {
		return nil, nil, err
	}
	link, err := NewLinkCQRSRPCServer(runRPCServer, link_cqrsService, log)
	if err != nil {
		return nil, nil, err
	}
	v1Link, err := NewLinkRPCServer(runRPCServer, service, log)
	if err != nil {
		return nil, nil, err
	}
	response, err := NewRunRPCServer(runRPCServer, link, v1Link)
	if err != nil {
		return nil, nil, err
	}
	sitemap, err := NewSitemapRPCServer(runRPCServer, sitemapService, log)
	if err != nil {
		return nil, nil, err
	}
	linkService, err := NewLinkService(log, service, link_cqrsService, sitemapService, event, response, v1Link, link, sitemap, store, cqsStore, queryStore)
	if err != nil {
		return nil, nil, err
	}
	return linkService, func() {
	}, nil
}

// di.go:

type LinkService struct {
	Logger logger.Logger

	// Delivery
	linkMQ            *api_mq.Event
	run               *run.Response
	linkRPCServer     *v1_2.Link
	linkCQRSRPCServer *v1_3.Link
	sitemapRPCServer  *v1_4.Sitemap

	// Application
	linkService     *link.Service
	linkCQRSService *link_cqrs.Service
	sitemapService  *sitemap.Service

	// Repository
	linkStore *crud.Store

	// CQRS
	cqsStore   *cqs.Store
	queryStore *query.Store
}

// LinkService =========================================================================================================
var LinkSet = wire.NewSet(

	InitLinkMQ,

	NewLinkRPCServer,
	NewLinkCQRSRPCServer,
	NewSitemapRPCServer,
	NewRunRPCServer,

	NewLinkRPCClient,
	NewMetadataRPCClient,

	NewLinkApplication,
	NewLinkCQRSApplication,
	NewSitemapApplication,

	NewLinkStore,
	NewCQSLinkStore,
	NewQueryLinkStore,

	NewLinkService,
)

func InitLinkMQ(ctx context.Context, log logger.Logger, mq v1.MQ) (*api_mq.Event, error) {
	linkMQ, err := api_mq.New(mq, log)
	if err != nil {
		return nil, err
	}
	notify.Subscribe(v1_5.METHOD_ADD, linkMQ)

	return linkMQ, nil
}

func NewLinkStore(ctx context.Context, logger2 logger.Logger, db2 *db.Store) (*crud.Store, error) {
	store := &crud.Store{}
	linkStore, err := store.Use(ctx, logger2, db2)
	if err != nil {
		return nil, err
	}

	return linkStore, nil
}

func NewCQSLinkStore(ctx context.Context, logger2 logger.Logger, db2 *db.Store) (*cqs.Store, error) {
	store := &cqs.Store{}
	cqsStore, err := store.Use(ctx, logger2, db2)
	if err != nil {
		return nil, err
	}

	return cqsStore, nil
}

func NewQueryLinkStore(ctx context.Context, logger2 logger.Logger, db2 *db.Store) (*query.Store, error) {
	store := &query.Store{}
	queryStore, err := store.Use(ctx, logger2, db2)
	if err != nil {
		return nil, err
	}

	return queryStore, nil
}

func NewLinkApplication(logger2 logger.Logger, metadataService v1_6.MetadataServiceClient, store *crud.Store) (*link.Service, error) {
	linkService, err := link.New(logger2, metadataService, store)
	if err != nil {
		return nil, err
	}

	return linkService, nil
}

func NewLinkCQRSApplication(logger2 logger.Logger, cqsStore *cqs.Store, queryStore *query.Store) (*link_cqrs.Service, error) {
	linkCQRSService, err := link_cqrs.New(logger2, cqsStore, queryStore)
	if err != nil {
		return nil, err
	}

	return linkCQRSService, nil
}

func NewLinkRPCClient(runRPCClient *grpc.ClientConn) (v1_2.LinkServiceClient, error) {
	LinkServiceClient := v1_2.NewLinkServiceClient(runRPCClient)
	return LinkServiceClient, nil
}

func NewSitemapApplication(logger2 logger.Logger, mq v1.MQ) (*sitemap.Service, error) {
	sitemapService, err := sitemap.New(logger2, mq)
	if err != nil {
		return nil, err
	}

	return sitemapService, nil
}

func NewLinkCQRSRPCServer(runRPCServer *rpc.RPCServer, application *link_cqrs.Service, log logger.Logger) (*v1_3.Link, error) {
	linkRPCServer, err := v1_3.New(runRPCServer, application, log)
	if err != nil {
		return nil, err
	}

	return linkRPCServer, nil
}

func NewLinkRPCServer(runRPCServer *rpc.RPCServer, application *link.Service, log logger.Logger) (*v1_2.Link, error) {
	linkRPCServer, err := v1_2.New(runRPCServer, application, log)
	if err != nil {
		return nil, err
	}

	return linkRPCServer, nil
}

func NewSitemapRPCServer(runRPCServer *rpc.RPCServer, application *sitemap.Service, log logger.Logger) (*v1_4.Sitemap, error) {
	sitemapRPCServer, err := v1_4.New(runRPCServer, application, log)
	if err != nil {
		return nil, err
	}

	return sitemapRPCServer, nil
}

func NewRunRPCServer(runRPCServer *rpc.RPCServer, cqrsLinkRPC *v1_3.Link, linkRPC *v1_2.Link) (*run.Response, error) {
	return run.Run(runRPCServer)
}

func NewMetadataRPCClient(runRPCClient *grpc.ClientConn) (v1_6.MetadataServiceClient, error) {
	metadataRPCClient := v1_6.NewMetadataServiceClient(runRPCClient)
	return metadataRPCClient, nil
}

func NewLinkService(
	log logger.Logger,

	linkService *link.Service,
	linkCQRSService *link_cqrs.Service,
	sitemapService *sitemap.Service,

	linkMQ *api_mq.Event, run2 *run.Response,
	linkRPCServer *v1_2.Link,
	linkCQRSRPCServer *v1_3.Link,
	sitemapRPCServer *v1_4.Sitemap,

	linkStore *crud.Store,

	cqsStore *cqs.Store,
	queryStore *query.Store,
) (*LinkService, error) {
	return &LinkService{
		Logger: log,

		linkService:     linkService,
		linkCQRSService: linkCQRSService,
		sitemapService:  sitemapService,

		run:               run2,
		linkRPCServer:     linkRPCServer,
		linkCQRSRPCServer: linkCQRSRPCServer,
		sitemapRPCServer:  sitemapRPCServer,
		linkMQ:            linkMQ,

		linkStore: linkStore,

		cqsStore:   cqsStore,
		queryStore: queryStore,
	}, nil
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package link_di

import (
	"context"
	"github.com/authzed/authzed-go/v1"
	cache2 "github.com/go-redis/cache/v9"
	"github.com/google/wire"
	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	"github.com/shortlink-org/shortlink/internal/di/pkg/context"
	"github.com/shortlink-org/shortlink/internal/di/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/di/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/di/pkg/permission"
	"github.com/shortlink-org/shortlink/internal/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/internal/di/pkg/store"
	"github.com/shortlink-org/shortlink/internal/di/pkg/traicing"
	"github.com/shortlink-org/shortlink/internal/pkg/cache"
	"github.com/shortlink-org/shortlink/internal/pkg/db"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/internal/pkg/rpc"
	"github.com/shortlink-org/shortlink/internal/services/link/application/link"
	"github.com/shortlink-org/shortlink/internal/services/link/application/link_cqrs"
	"github.com/shortlink-org/shortlink/internal/services/link/application/sitemap"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/mq"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/cqrs/cqs"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/cqrs/query"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud"
	v1_2 "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/run"
	v1_3 "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/sitemap/v1"
	v1_4 "github.com/shortlink-org/shortlink/internal/services/metadata/infrastructure/rpc/metadata/v1"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

// Injectors from wire.go:

func InitializeLinkService() (*LinkService, func(), error) {
	context, cleanup, err := ctx.New()
	if err != nil {
		return nil, nil, err
	}
	logger, cleanup2, err := logger_di.New(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	configConfig, err := config.New()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	monitoringMonitoring, cleanup3, err := monitoring.New(context, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tracerProvider, cleanup4, err := traicing_di.New(context, logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	pprofEndpoint, err := profiling.New(context, logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	autoMaxProAutoMaxPro, cleanup5, err := autoMaxPro.New(logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	client, err := permission.New(context, logger, tracerProvider, monitoringMonitoring)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	mq, cleanup6, err := mq_di.New(context, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	clientConn, cleanup7, err := rpc.InitClient(logger, tracerProvider, monitoringMonitoring)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	metadataServiceClient, err := NewMetadataRPCClient(clientConn)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	db, cleanup8, err := store.New(context, logger, tracerProvider, monitoringMonitoring)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	cacheCache, err := cache.New(context)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	repository, err := NewLinkStore(context, logger, db, cacheCache)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewLinkApplication(logger, mq, metadataServiceClient, repository, client)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	cqsStore, err := NewCQSLinkStore(context, logger, db, cacheCache)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	queryStore, err := NewQueryLinkStore(context, logger, db, cacheCache)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	link_cqrsService, err := NewLinkCQRSApplication(logger, cqsStore, queryStore)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	sitemapService, err := NewSitemapApplication(logger, mq)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	event, err := InitLinkMQ(context, logger, mq, service)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	server, cleanup9, err := rpc.InitServer(logger, tracerProvider, monitoringMonitoring)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	link, err := NewLinkCQRSRPCServer(server, link_cqrsService, logger)
	if err != nil {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	v1Link, err := NewLinkRPCServer(server, service, logger)
	if err != nil {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	response, err := NewRunRPCServer(server, link, v1Link)
	if err != nil {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	sitemap, err := NewSitemapRPCServer(server, sitemapService, logger)
	if err != nil {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	linkService, err := NewLinkService(logger, configConfig, monitoringMonitoring, tracerProvider, pprofEndpoint, autoMaxProAutoMaxPro, client, service, link_cqrsService, sitemapService, event, response, v1Link, link, sitemap, repository, cqsStore, queryStore)
	if err != nil {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return linkService, func() {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

type LinkService struct {
	// Common
	Log    logger.Logger
	Config *config.Config

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro

	// Security
	authPermission *authzed.Client

	// Delivery
	linkMQ            *api_mq.Event
	run               *run.Response
	linkRPCServer     *v1.Link
	linkCQRSRPCServer *v1_2.Link
	sitemapRPCServer  *v1_3.Sitemap

	// Application
	linkService     *link.Service
	linkCQRSService *link_cqrs.Service
	sitemapService  *sitemap.Service

	// Repository
	linkStore crud.Repository

	// CQRS
	cqsStore   *cqs.Store
	queryStore *query.Store
}

// LinkService =========================================================================================================
var LinkSet = wire.NewSet(di.DefaultSet, mq_di.New, rpc.InitServer, rpc.InitClient, store.New, InitLinkMQ,

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

func InitLinkMQ(ctx2 context.Context, log logger.Logger, mq2 mq.MQ, service *link.Service) (*api_mq.Event, error) {
	linkMQ, err := api_mq.New(mq2, log, service)
	if err != nil {
		return nil, err
	}

	return linkMQ, nil
}

func NewLinkStore(ctx2 context.Context, log logger.Logger, db2 db.DB, cache3 *cache2.Cache) (crud.Repository, error) {
	linkStore, err := crud.New(ctx2, log, db2, cache3)
	if err != nil {
		return nil, err
	}

	return linkStore, nil
}

func NewCQSLinkStore(ctx2 context.Context, log logger.Logger, db2 db.DB, cache3 *cache2.Cache) (*cqs.Store, error) {
	store2, err := cqs.New(ctx2, log, db2, cache3)
	if err != nil {
		return nil, err
	}

	return store2, nil
}

func NewQueryLinkStore(ctx2 context.Context, log logger.Logger, db2 db.DB, cache3 *cache2.Cache) (*query.Store, error) {
	store2, err := query.New(ctx2, log, db2, cache3)
	if err != nil {
		return nil, err
	}

	return store2, nil
}

func NewLinkApplication(log logger.Logger, mq2 mq.MQ, metadataService v1_4.MetadataServiceClient, store2 crud.Repository, authPermission *authzed.Client) (*link.Service, error) {
	linkService, err := link.New(log, mq2, metadataService, store2, authPermission)
	if err != nil {
		return nil, err
	}

	return linkService, nil
}

func NewLinkCQRSApplication(log logger.Logger, cqsStore *cqs.Store, queryStore *query.Store) (*link_cqrs.Service, error) {
	linkCQRSService, err := link_cqrs.New(log, cqsStore, queryStore)
	if err != nil {
		return nil, err
	}

	return linkCQRSService, nil
}

func NewLinkRPCClient(runRPCClient *grpc.ClientConn) (v1.LinkServiceClient, error) {
	LinkServiceClient := v1.NewLinkServiceClient(runRPCClient)
	return LinkServiceClient, nil
}

func NewSitemapApplication(log logger.Logger, dataBus mq.MQ) (*sitemap.Service, error) {
	sitemapService, err := sitemap.New(log, dataBus)
	if err != nil {
		return nil, err
	}

	return sitemapService, nil
}

func NewLinkCQRSRPCServer(runRPCServer *rpc.Server, application *link_cqrs.Service, log logger.Logger) (*v1_2.Link, error) {
	linkRPCServer, err := v1_2.New(runRPCServer, application, log)
	if err != nil {
		return nil, err
	}

	return linkRPCServer, nil
}

func NewLinkRPCServer(runRPCServer *rpc.Server, application *link.Service, log logger.Logger) (*v1.Link, error) {
	linkRPCServer, err := v1.New(runRPCServer, application, log)
	if err != nil {
		return nil, err
	}

	return linkRPCServer, nil
}

func NewSitemapRPCServer(runRPCServer *rpc.Server, application *sitemap.Service, log logger.Logger) (*v1_3.Sitemap, error) {
	sitemapRPCServer, err := v1_3.New(runRPCServer, application, log)
	if err != nil {
		return nil, err
	}

	return sitemapRPCServer, nil
}

func NewRunRPCServer(runRPCServer *rpc.Server, _ *v1_2.Link, _ *v1.Link) (*run.Response, error) {
	return run.Run(runRPCServer)
}

func NewMetadataRPCClient(runRPCClient *grpc.ClientConn) (v1_4.MetadataServiceClient, error) {
	metadataRPCClient := v1_4.NewMetadataServiceClient(runRPCClient)
	return metadataRPCClient, nil
}

func NewLinkService(

	log logger.Logger, config2 *config.Config, monitoring2 *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,

	authPermission *authzed.Client,

	linkService *link.Service,
	linkCQRSService *link_cqrs.Service,
	sitemapService *sitemap.Service,

	linkMQ *api_mq.Event, run2 *run.Response,
	linkRPCServer *v1.Link,
	linkCQRSRPCServer *v1_2.Link,
	sitemapRPCServer *v1_3.Sitemap,

	linkStore crud.Repository,

	cqsStore *cqs.Store,
	queryStore *query.Store,
) (*LinkService, error) {
	return &LinkService{

		Log:    log,
		Config: config2,

		Tracer:        tracer,
		Monitoring:    monitoring2,
		PprofEndpoint: pprofHTTP,
		AutoMaxPro:    autoMaxProcsOption,

		authPermission: authPermission,

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

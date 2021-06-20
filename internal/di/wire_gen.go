// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package di

import (
	"context"
	"github.com/batazor/shortlink/internal/di/internal/autoMaxPro"
	"github.com/batazor/shortlink/internal/di/internal/config"
	"github.com/batazor/shortlink/internal/di/internal/context"
	"github.com/batazor/shortlink/internal/di/internal/flags"
	"github.com/batazor/shortlink/internal/di/internal/logger"
	"github.com/batazor/shortlink/internal/di/internal/monitoring"
	"github.com/batazor/shortlink/internal/di/internal/mq"
	"github.com/batazor/shortlink/internal/di/internal/profiling"
	"github.com/batazor/shortlink/internal/di/internal/sentry"
	"github.com/batazor/shortlink/internal/di/internal/store"
	"github.com/batazor/shortlink/internal/di/internal/traicing"
	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq"
	"github.com/batazor/shortlink/internal/services/api/di"
	di2 "github.com/batazor/shortlink/internal/services/link/di"
	di3 "github.com/batazor/shortlink/internal/services/metadata/di"
	"github.com/batazor/shortlink/internal/services/metadata/infrastructure/store"
	"github.com/batazor/shortlink/pkg/rpc"
	"github.com/getsentry/sentry-go/http"
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"net/http"
)

// Injectors from default.go:

func InitializeFullService() (*Service, func(), error) {
	context, cleanup, err := ctx.New()
	if err != nil {
		return nil, nil, err
	}
	configConfig, err := config.New()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	logger, cleanup2, err := logger_di.New(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mq, cleanup3, err := mq_di.New(context, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	handler, cleanup4, err := sentry.New()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serveMux := monitoring.New(handler, logger)
	tracer, cleanup5, err := traicing_di.New(context, logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dbStore, cleanup6, err := store.New(context, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	pprofEndpoint := profiling.New(logger)
	autoMaxProAutoMaxPro, cleanup7, err := autoMaxPro.New(logger)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	rpcServer, cleanup8, err := rpc.InitServer(logger, tracer)
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
	clientConn, cleanup9, err := rpc.InitClient(logger, tracer)
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
	service, err := NewFullService(context, configConfig, logger, mq, handler, serveMux, tracer, dbStore, pprofEndpoint, autoMaxProAutoMaxPro, rpcServer, clientConn)
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
	return service, func() {
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

// Injectors from service_api.go:

func InitializeAPIService() (*ServiceAPI, func(), error) {
	context, cleanup, err := ctx.New()
	if err != nil {
		return nil, nil, err
	}
	configConfig, err := config.New()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	logger, cleanup2, err := logger_di.New(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	handler, cleanup3, err := sentry.New()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serveMux := monitoring.New(handler, logger)
	tracer, cleanup4, err := traicing_di.New(context, logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	pprofEndpoint := profiling.New(logger)
	autoMaxProAutoMaxPro, cleanup5, err := autoMaxPro.New(logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	clientConn, cleanup6, err := rpc.InitClient(logger, tracer)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	rpcServer, cleanup7, err := rpc.InitServer(logger, tracer)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	apiService, cleanup8, err := InitAPIService(context, clientConn, rpcServer, logger, tracer)
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
	serviceAPI, err := NewAPIService(context, configConfig, logger, handler, serveMux, tracer, pprofEndpoint, autoMaxProAutoMaxPro, clientConn, apiService)
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
	return serviceAPI, func() {
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

// Injectors from service_billing.go:

func InitializeBillingService() (*Service, func(), error) {
	context, cleanup, err := ctx.New()
	if err != nil {
		return nil, nil, err
	}
	configConfig, err := config.New()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	logger, cleanup2, err := logger_di.New(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	handler, cleanup3, err := sentry.New()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serveMux := monitoring.New(handler, logger)
	tracer, cleanup4, err := traicing_di.New(context, logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	mq, cleanup5, err := mq_di.New(context, logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	autoMaxProAutoMaxPro, cleanup6, err := autoMaxPro.New(logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewBillingService(context, configConfig, logger, serveMux, tracer, mq, autoMaxProAutoMaxPro)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// Injectors from service_link.go:

func InitializeLinkService() (*ServiceLink, func(), error) {
	context, cleanup, err := ctx.New()
	if err != nil {
		return nil, nil, err
	}
	configConfig, err := config.New()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	logger, cleanup2, err := logger_di.New(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mq, cleanup3, err := mq_di.New(context, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	handler, cleanup4, err := sentry.New()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serveMux := monitoring.New(handler, logger)
	tracer, cleanup5, err := traicing_di.New(context, logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dbStore, cleanup6, err := store.New(context, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	pprofEndpoint := profiling.New(logger)
	autoMaxProAutoMaxPro, cleanup7, err := autoMaxPro.New(logger)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	rpcServer, cleanup8, err := rpc.InitServer(logger, tracer)
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
	clientConn, cleanup9, err := rpc.InitClient(logger, tracer)
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
	linkService, cleanup10, err := InitLinkService(context, clientConn, rpcServer, logger, dbStore, mq)
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
	serviceLink, err := NewLinkService(context, configConfig, logger, mq, handler, serveMux, tracer, dbStore, pprofEndpoint, autoMaxProAutoMaxPro, rpcServer, clientConn, linkService)
	if err != nil {
		cleanup10()
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
	return serviceLink, func() {
		cleanup10()
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

// Injectors from service_logger.go:

func InitializeLoggerService() (*Service, func(), error) {
	context, cleanup, err := ctx.New()
	if err != nil {
		return nil, nil, err
	}
	configConfig, err := config.New()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	logger, cleanup2, err := logger_di.New(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	handler, cleanup3, err := sentry.New()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serveMux := monitoring.New(handler, logger)
	tracer, cleanup4, err := traicing_di.New(context, logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	mq, cleanup5, err := mq_di.New(context, logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	autoMaxProAutoMaxPro, cleanup6, err := autoMaxPro.New(logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewLoggerService(context, configConfig, logger, serveMux, tracer, mq, autoMaxProAutoMaxPro)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// Injectors from service_metadata.go:

func InitializeMetadataService() (*ServiceMetadata, func(), error) {
	context, cleanup, err := ctx.New()
	if err != nil {
		return nil, nil, err
	}
	logger, cleanup2, err := logger_di.New(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	autoMaxProAutoMaxPro, cleanup3, err := autoMaxPro.New(logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dbStore, cleanup4, err := store.New(context, logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tracer, cleanup5, err := traicing_di.New(context, logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	rpcServer, cleanup6, err := rpc.InitServer(logger, tracer)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	handler, cleanup7, err := sentry.New()
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serveMux := monitoring.New(handler, logger)
	metaDataService, cleanup8, err := InitMetaService(context, rpcServer, logger, dbStore)
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
	serviceMetadata, err := NewMetadataService(logger, autoMaxProAutoMaxPro, dbStore, rpcServer, serveMux, handler, metaDataService)
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
	return serviceMetadata, func() {
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

// Injectors from service_notify.go:

func InitializeNotifyService() (*Service, func(), error) {
	context, cleanup, err := ctx.New()
	if err != nil {
		return nil, nil, err
	}
	configConfig, err := config.New()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	logger, cleanup2, err := logger_di.New(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mq, cleanup3, err := mq_di.New(context, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	handler, cleanup4, err := sentry.New()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serveMux := monitoring.New(handler, logger)
	tracer, cleanup5, err := traicing_di.New(context, logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	autoMaxProAutoMaxPro, cleanup6, err := autoMaxPro.New(logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewNotifyService(context, configConfig, logger, mq, serveMux, tracer, autoMaxProAutoMaxPro)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// default.go:

// Service - heplers
type Service struct {
	Ctx           context.Context
	Cfg           *config.Config
	Log           logger.Logger
	Tracer        *opentracing.Tracer
	Sentry        *sentryhttp.Handler
	DB            *db.Store
	MetaStore     *meta_store.MetaStore
	MQ            mq.MQ
	ServerRPC     *rpc.RPCServer
	ClientRPC     *grpc.ClientConn
	Monitoring    *http.ServeMux
	PprofEndpoint profiling.PprofEndpoint
}

// Default =============================================================================================================
var DefaultSet = wire.NewSet(ctx.New, autoMaxPro.New, flags.New, config.New, logger_di.New, traicing_di.New)

// FullService =========================================================================================================
var FullSet = wire.NewSet(
	DefaultSet,
	NewFullService, store.New, sentry.New, monitoring.New, profiling.New, mq_di.New, rpc.InitServer, rpc.InitClient,
)

func NewFullService(ctx2 context.Context,

	cfg *config.Config,
	log logger.Logger, mq2 mq.MQ,

	sentryHandler *sentryhttp.Handler, monitoring2 *http.ServeMux,
	tracer *opentracing.Tracer, db2 *db.Store,

	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
	serverRPC *rpc.RPCServer,
	clientRPC *grpc.ClientConn,
) (*Service, error) {
	return &Service{
		Ctx:           ctx2,
		Cfg:           cfg,
		Log:           log,
		MQ:            mq2,
		Tracer:        tracer,
		Monitoring:    monitoring2,
		Sentry:        sentryHandler,
		DB:            db2,
		PprofEndpoint: pprofHTTP,
		ClientRPC:     clientRPC,
		ServerRPC:     serverRPC,
	}, nil
}

// service_api.go:

type ServiceAPI struct {
	Service

	APIService *di.APIService
}

// InitAPIService =====================================================================================================
func InitAPIService(ctx2 context.Context, runRPCClient *grpc.ClientConn, runRPCServer *rpc.RPCServer, log logger.Logger, tracer *opentracing.Tracer) (*di.APIService, func(), error) {
	return di.InitializeAPIService(ctx2, runRPCClient, runRPCServer, log, tracer)
}

// APIService ==========================================================================================================
var APISet = wire.NewSet(
	DefaultSet, sentry.New, monitoring.New, profiling.New, rpc.InitServer, rpc.InitClient, InitAPIService,
	NewAPIService,
)

func NewAPIService(ctx2 context.Context,

	cfg *config.Config,
	log logger.Logger,
	sentryHandler *sentryhttp.Handler, monitoring2 *http.ServeMux,
	tracer *opentracing.Tracer,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
	clientRPC *grpc.ClientConn,
	apiService *di.APIService,
) (*ServiceAPI, error) {
	return &ServiceAPI{
		Service: Service{
			Ctx:           ctx2,
			Log:           log,
			Tracer:        tracer,
			Monitoring:    monitoring2,
			Sentry:        sentryHandler,
			PprofEndpoint: pprofHTTP,
			ClientRPC:     clientRPC,
		},
		APIService: apiService,
	}, nil
}

// service_billing.go:

// BillingService =======================================================================================================
var BillingSet = wire.NewSet(
	DefaultSet, mq_di.New, sentry.New, monitoring.New, NewBillingService,
)

func NewBillingService(ctx2 context.Context,

	cfg *config.Config,
	log logger.Logger, monitoring2 *http.ServeMux,
	tracer *opentracing.Tracer, mq2 mq.MQ,

	autoMaxProcsOption autoMaxPro.AutoMaxPro,
) (*Service, error) {
	return &Service{
		Ctx:        ctx2,
		Log:        log,
		MQ:         mq2,
		Tracer:     tracer,
		Monitoring: monitoring2,
	}, nil
}

// service_link.go:

type ServiceLink struct {
	Service

	LinkService *di2.LinkService
}

// InitLinkService =====================================================================================================
func InitLinkService(ctx2 context.Context, runRPCClient *grpc.ClientConn, runRPCServer *rpc.RPCServer, log logger.Logger, db2 *db.Store, mq2 mq.MQ) (*di2.LinkService, func(), error) {
	return di2.InitializeLinkService(ctx2, runRPCClient, runRPCServer, log, db2, mq2)
}

// APIService ==========================================================================================================
var LinkSet = wire.NewSet(
	DefaultSet, store.New, sentry.New, monitoring.New, profiling.New, mq_di.New, rpc.InitServer, rpc.InitClient, InitLinkService,
	NewLinkService,
)

func NewLinkService(ctx2 context.Context,

	cfg *config.Config,
	log logger.Logger, mq2 mq.MQ,

	sentryHandler *sentryhttp.Handler, monitoring2 *http.ServeMux,
	tracer *opentracing.Tracer, db2 *db.Store,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
	serverRPC *rpc.RPCServer,
	clientRPC *grpc.ClientConn,
	linkService *di2.LinkService,
) (*ServiceLink, error) {
	return &ServiceLink{
		Service: Service{
			Ctx:           ctx2,
			Log:           log,
			MQ:            mq2,
			Tracer:        tracer,
			Monitoring:    monitoring2,
			Sentry:        sentryHandler,
			DB:            db2,
			PprofEndpoint: pprofHTTP,
			ClientRPC:     clientRPC,
			ServerRPC:     serverRPC,
		},
		LinkService: linkService,
	}, nil
}

// service_logger.go:

// LoggerService =======================================================================================================
var LoggerSet = wire.NewSet(
	DefaultSet, mq_di.New, sentry.New, monitoring.New, NewLoggerService,
)

func NewLoggerService(ctx2 context.Context,

	cfg *config.Config,
	log logger.Logger, monitoring2 *http.ServeMux,
	tracer *opentracing.Tracer, mq2 mq.MQ,

	autoMaxProcsOption autoMaxPro.AutoMaxPro,
) (*Service, error) {
	return &Service{
		Ctx:        ctx2,
		Log:        log,
		MQ:         mq2,
		Tracer:     tracer,
		Monitoring: monitoring2,
	}, nil
}

// service_metadata.go:

type ServiceMetadata struct {
	Service

	MetaService *di3.MetaDataService
}

// InitMetaService =====================================================================================================
func InitMetaService(ctx2 context.Context, runRPCServer *rpc.RPCServer, log logger.Logger, db2 *db.Store) (*di3.MetaDataService, func(), error) {
	return di3.InitializeMetaDataService(ctx2, runRPCServer, log, db2)
}

// MetadataService =====================================================================================================
var MetadataSet = wire.NewSet(
	DefaultSet, store.New, rpc.InitServer, sentry.New, monitoring.New, InitMetaService,
	NewMetadataService,
)

func NewMetadataService(
	log logger.Logger,
	autoMaxProcsOption autoMaxPro.AutoMaxPro, db2 *db.Store,
	serverRPC *rpc.RPCServer, monitoring2 *http.ServeMux,
	sentryHandler *sentryhttp.Handler,
	metadataService *di3.MetaDataService,
) (*ServiceMetadata, error) {
	return &ServiceMetadata{
		Service: Service{
			Log:        log,
			ServerRPC:  serverRPC,
			DB:         db2,
			Monitoring: monitoring2,
			Sentry:     sentryHandler,
		},
		MetaService: metadataService,
	}, nil
}

// service_notify.go:

// NotifyService ==========================================================================================================
var NotifySet = wire.NewSet(
	DefaultSet, mq_di.New, sentry.New, monitoring.New, NewNotifyService,
)

func NewNotifyService(ctx2 context.Context,

	cfg *config.Config,
	log logger.Logger, mq2 mq.MQ, monitoring2 *http.ServeMux,
	tracer *opentracing.Tracer,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
) (*Service, error) {
	return &Service{
		Ctx:        ctx2,
		Log:        log,
		MQ:         mq2,
		Tracer:     tracer,
		Monitoring: monitoring2,
	}, nil
}

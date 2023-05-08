// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package metadata_di

import (
	"context"
	"github.com/google/wire"
	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	"github.com/shortlink-org/shortlink/internal/di/pkg/context"
	"github.com/shortlink-org/shortlink/internal/di/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/di/pkg/monitoring"
	"github.com/shortlink-org/shortlink/internal/di/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/internal/di/pkg/store"
	"github.com/shortlink-org/shortlink/internal/di/pkg/traicing"
	"github.com/shortlink-org/shortlink/internal/pkg/db"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
	"github.com/shortlink-org/shortlink/internal/services/metadata/application/parsers"
	v1_2 "github.com/shortlink-org/shortlink/internal/services/metadata/domain/metadata/v1"
	"github.com/shortlink-org/shortlink/internal/services/metadata/infrastructure/mq"
	"github.com/shortlink-org/shortlink/internal/services/metadata/infrastructure/rpc/metadata/v1"
	"github.com/shortlink-org/shortlink/internal/services/metadata/infrastructure/store"
	"github.com/shortlink-org/shortlink/pkg/rpc"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

// Injectors from wire.go:

func InitializeMetaDataService() (*MetaDataService, func(), error) {
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
	serveMux, err := monitoring.New(logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tracerProvider, cleanup3, err := traicing_di.New(context, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	pprofEndpoint, err := profiling.New(logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	autoMaxProAutoMaxPro, cleanup4, err := autoMaxPro.New(logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dbStore, cleanup5, err := store.New(context, logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	metaStore, err := NewMetaDataStore(context, logger, dbStore)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewMetaDataApplication(metaStore)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dataBus, cleanup6, err := mq_di.New(context, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	event, err := InitMetadataMQ(context, logger, dataBus)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	rpcServer, cleanup7, err := rpc.InitServer(logger, tracerProvider)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	metadata, err := NewMetaDataRPCServer(rpcServer, service, logger)
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
	metaDataService, err := NewMetaDataService(logger, configConfig, serveMux, tracerProvider, pprofEndpoint, autoMaxProAutoMaxPro, service, event, metadata, metaStore)
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
	return metaDataService, func() {
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

type MetaDataService struct {
	// Common
	Log    logger.Logger
	Config *config.Config

	// Observability
	Tracer        *trace.TracerProvider
	Monitoring    *http.ServeMux
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro

	// Delivery
	metadataMQ        *metadata_mq.Event
	metadataRPCServer *v1.Metadata

	// Application
	service *parsers.Service

	// Repository
	metadataStore *meta_store.MetaStore
}

// MetaDataService =====================================================================================================
var MetaDataSet = wire.NewSet(di.DefaultSet, mq_di.New, store.New, rpc.InitServer, InitMetadataMQ,
	NewMetaDataRPCServer,

	NewMetaDataApplication,

	NewMetaDataStore,

	NewMetaDataService,
)

func InitMetadataMQ(ctx2 context.Context, log logger.Logger, mq2 *mq.DataBus) (*metadata_mq.Event, error) {
	metadataMQ, err := metadata_mq.New(mq2)
	if err != nil {
		return nil, err
	}
	notify.Subscribe(v1_2.METHOD_ADD, metadataMQ)

	return metadataMQ, nil
}

func NewMetaDataStore(ctx2 context.Context, logger2 logger.Logger, db2 *db.Store) (*meta_store.MetaStore, error) {
	store2 := &meta_store.MetaStore{}
	metadataStore, err := store2.Use(ctx2, logger2, db2)
	if err != nil {
		return nil, err
	}

	return metadataStore, nil
}

func NewMetaDataApplication(store2 *meta_store.MetaStore) (*parsers.Service, error) {
	metadataService, err := parsers.New(store2)
	if err != nil {
		return nil, err
	}

	return metadataService, nil
}

func NewMetaDataRPCServer(runRPCServer *rpc.RPCServer, application *parsers.Service, log logger.Logger) (*v1.Metadata, error) {
	metadataRPCServer, err := v1.New(runRPCServer, application, log)
	if err != nil {
		return nil, err
	}

	return metadataRPCServer, nil
}

func NewMetaDataService(

	log logger.Logger, config2 *config.Config, monitoring2 *http.ServeMux,
	tracer *trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,

	service *parsers.Service,

	metadataMQ *metadata_mq.Event,
	metadataRPCServer *v1.Metadata,

	metadataStore *meta_store.MetaStore,
) (*MetaDataService, error) {
	return &MetaDataService{

		Log:    log,
		Config: config2,

		Tracer:        tracer,
		Monitoring:    monitoring2,
		PprofEndpoint: pprofHTTP,
		AutoMaxPro:    autoMaxProcsOption,

		service: service,

		metadataMQ:        metadataMQ,
		metadataRPCServer: metadataRPCServer,

		metadataStore: metadataStore,
	}, nil
}

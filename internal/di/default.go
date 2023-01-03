/*
Main DI-package
*/
package di

import (
	"context"
	"net/http"

	sentryhttp "github.com/getsentry/sentry-go/http"
	redisCache "github.com/go-redis/redis/v9"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/text/message"
	"google.golang.org/grpc"

	"github.com/shortlink-org/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	ctx "github.com/shortlink-org/shortlink/internal/di/pkg/context"
	"github.com/shortlink-org/shortlink/internal/di/pkg/flags"
	"github.com/shortlink-org/shortlink/internal/di/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/di/pkg/monitoring"
	mq_di "github.com/shortlink-org/shortlink/internal/di/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/internal/di/pkg/sentry"
	"github.com/shortlink-org/shortlink/internal/di/pkg/store"
	"github.com/shortlink-org/shortlink/internal/pkg/cache"
	"github.com/shortlink-org/shortlink/internal/pkg/db"
	"github.com/shortlink-org/shortlink/internal/pkg/i18n"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/mq/v1"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

// Service - heplers
type Service struct {
	// Common
	Ctx  context.Context
	Cfg  *config.Config
	Log  logger.Logger
	I18N *message.Printer

	// Delivery
	DB        *db.Store
	Cache     *redisCache.UniversalClient
	MQ        v1.MQ
	ServerRPC *rpc.RPCServer
	ClientRPC *grpc.ClientConn

	// Observability
	Tracer        *trace.TracerProvider
	Sentry        *sentryhttp.Handler
	Monitoring    *http.ServeMux
	PprofEndpoint profiling.PprofEndpoint
}

// Default =============================================================================================================
var DefaultSet = wire.NewSet(ctx.New, autoMaxPro.New, flags.New, config.New, logger_di.New, traicing_di.New, cache.New, i18n.New)

// FullService =========================================================================================================
var FullSet = wire.NewSet(
	DefaultSet,
	NewFullService,
	store.New,
	sentry.New,
	monitoring.New,
	profiling.New,
	mq_di.New,
	rpc.InitServer,
	rpc.InitClient,
)

func NewFullService(
	// Common
	ctx context.Context,
	cfg *config.Config,
	log logger.Logger,
	i18n *message.Printer,

	// Delivery
	serverRPC *rpc.RPCServer,
	clientRPC *grpc.ClientConn,
	mq v1.MQ,
	db *db.Store,
	cache *redisCache.UniversalClient,

	// Observability
	sentryHandler *sentryhttp.Handler,
	monitoring *http.ServeMux,
	tracer *trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
) (*Service, error) {
	return &Service{
		// Common
		Ctx:  ctx,
		Cfg:  cfg,
		Log:  log,
		I18N: i18n,

		// Delivery
		MQ:        mq,
		DB:        db,
		Cache:     cache,
		ClientRPC: clientRPC,
		ServerRPC: serverRPC,

		// Observability
		Tracer:        tracer,
		Monitoring:    monitoring,
		Sentry:        sentryHandler,
		PprofEndpoint: pprofHTTP,
	}, nil
}

func InitializeFullService() (*Service, func(), error) {
	panic(wire.Build(FullSet))
}

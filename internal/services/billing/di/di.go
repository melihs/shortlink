//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package billing_di

import (
	"context"

	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq"
	account_application "github.com/batazor/shortlink/internal/services/billing/application/account"
	balance_application "github.com/batazor/shortlink/internal/services/billing/application/balance"
	order_application "github.com/batazor/shortlink/internal/services/billing/application/order"
	payment_application "github.com/batazor/shortlink/internal/services/billing/application/payment"
	tariff_application "github.com/batazor/shortlink/internal/services/billing/application/tariff"
	api "github.com/batazor/shortlink/internal/services/billing/infrastructure/api/http"
	balance_rpc "github.com/batazor/shortlink/internal/services/billing/infrastructure/api/rpc/balance/v1"
	order_rpc "github.com/batazor/shortlink/internal/services/billing/infrastructure/api/rpc/order/v1"
	payment_rpc "github.com/batazor/shortlink/internal/services/billing/infrastructure/api/rpc/payment/v1"
	tariff_rpc "github.com/batazor/shortlink/internal/services/billing/infrastructure/api/rpc/tariff/v1"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/store"
	"github.com/batazor/shortlink/pkg/rpc"
)

type BillingService struct {
	Logger logger.Logger

	// Delivery
	httpAPIServer    *api.Server
	balanceRPCServer *balance_rpc.Balance
	orderRPCServer   *order_rpc.Order
	paymentRPCServer *payment_rpc.Payment
	tariffRPCServer  *tariff_rpc.Tariff

	// Repository
	accountRepository *billing_store.AccountRepository
	balanceRepository *billing_store.BalanceRepository
	orderRepository   *billing_store.OrderRepository
	paymentRepository *billing_store.PaymentRepository
	tariffRepository  *billing_store.TariffRepository
}

// BillingService ======================================================================================================
var BillingSet = wire.NewSet(
	// infrastructure
	NewBillingAPIServer,

	// repository
	NewBillingStore,

	// application
	NewTariffApplication,
	NewAccountApplication,
	NewBalanceApplication,
	NewOrderApplication,
	NewPaymentApplication,

	NewBillingService,
)

func NewBillingStore(ctx context.Context, logger logger.Logger, db *db.Store) (*billing_store.BillingStore, error) {
	store := &billing_store.BillingStore{}
	billingStore, err := store.Use(ctx, logger, db)
	if err != nil {
		return nil, err
	}

	return billingStore, nil
}

func NewAccountApplication(logger logger.Logger, store *billing_store.BillingStore) (*account_application.AccountService, error) {
	accountService, err := account_application.New(logger, store.Account)
	if err != nil {
		return nil, err
	}

	return accountService, nil
}

func NewBalanceApplication(logger logger.Logger, store *billing_store.BillingStore) (*balance_application.BalanceService, error) {
	balanceService, err := balance_application.New(logger, store.Balance)
	if err != nil {
		return nil, err
	}

	return balanceService, nil
}

func NewOrderApplication(logger logger.Logger, store *billing_store.BillingStore) (*order_application.OrderService, error) {
	orderService, err := order_application.New(logger, store.Order)
	if err != nil {
		return nil, err
	}

	return orderService, nil
}

func NewPaymentApplication(logger logger.Logger, store *billing_store.BillingStore) (*payment_application.PaymentService, error) {
	paymentService, err := payment_application.New(logger, store.Payment)
	if err != nil {
		return nil, err
	}

	return paymentService, nil
}

func NewTariffApplication(logger logger.Logger, store *billing_store.BillingStore) (*tariff_application.TariffService, error) {
	tariffService, err := tariff_application.New(logger, store.Tariff)
	if err != nil {
		return nil, err
	}

	return tariffService, nil
}

func NewBillingAPIServer(
	ctx context.Context,
	logger logger.Logger,
	tracer *opentracing.Tracer,
	rpcServer *rpc.RPCServer,
	db *db.Store,

	// Applications
	accountService *account_application.AccountService,
	balanceService *balance_application.BalanceService,
	orderService *order_application.OrderService,
	paymentService *payment_application.PaymentService,
	tariffService *tariff_application.TariffService,
) (*api.Server, error) {
	// Run API server
	API := api.Server{}

	apiService, err := API.Use(
		ctx,
		db,
		logger,
		tracer,

		// services
		accountService,
		balanceService,
		orderService,
		paymentService,
		tariffService,
	)
	if err != nil {
		return nil, err
	}

	return apiService, nil
}

func NewBillingService(
	log logger.Logger,

	// Delivery
	httpAPIServer *api.Server,
) (*BillingService, error) {
	return &BillingService{
		Logger: log,

		// Delivery
		httpAPIServer: httpAPIServer,
	}, nil
}

func InitializeBillingService(ctx context.Context, runRPCClient *grpc.ClientConn, runRPCServer *rpc.RPCServer, log logger.Logger, db *db.Store, mq mq.MQ, tracer *opentracing.Tracer) (*BillingService, func(), error) {
	panic(wire.Build(BillingSet))
}

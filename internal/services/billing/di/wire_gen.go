// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package billing_di

import (
	"context"
	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq"
	"github.com/batazor/shortlink/internal/services/billing/application/account"
	"github.com/batazor/shortlink/internal/services/billing/application/payment"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/api/http"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/rpc/balance/v1"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/rpc/order/v1"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/rpc/payment/v1"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/rpc/tariff/v1"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/store"
	"github.com/batazor/shortlink/pkg/rpc"
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

// Injectors from di.go:

func InitializeBillingService(ctx context.Context, runRPCClient *grpc.ClientConn, runRPCServer *rpc.RPCServer, log logger.Logger, db2 *db.Store, mq2 mq.MQ, tracer *opentracing.Tracer) (*BillingService, func(), error) {
	billingStore, err := NewBillingStore(ctx, log, db2)
	if err != nil {
		return nil, nil, err
	}
	accountService, err := NewAccountApplication(log, billingStore)
	if err != nil {
		return nil, nil, err
	}
	server, err := NewBillingAPIServer(ctx, log, tracer, runRPCServer, db2, accountService)
	if err != nil {
		return nil, nil, err
	}
	billingService, err := NewBillingService(log, server)
	if err != nil {
		return nil, nil, err
	}
	return billingService, func() {
	}, nil
}

// di.go:

type BillingService struct {
	Logger logger.Logger

	// Delivery
	httpAPIServer    *api.Server
	balanceRPCServer *balance_rpc.Balance
	orderRPCServer   *order_rpc.Order
	paymentRPCServer *payment_rpc.Payment
	tariffRPCServer  *tariff_rpc.Tariff

	// Application
	payment *payment.Payment

	// Repository
	accountRepository *billing_store.AccountRepository
	balanceRepository *billing_store.BalanceRepository
	orderRepository   *billing_store.OrderRepository
	paymentRepository *billing_store.PaymentRepository
	tariffRepository  *billing_store.TariffRepository
}

// BillingService ======================================================================================================
var BillingSet = wire.NewSet(

	NewBillingAPIServer,

	NewBillingStore,

	NewAccountApplication,

	NewBillingService,
)

func NewAccountApplication(logger2 logger.Logger, store *billing_store.BillingStore) (*account_application.AccountService, error) {
	accountService, err := account_application.New(logger2, store.Account)
	if err != nil {
		return nil, err
	}

	return accountService, nil
}

func NewBillingAPIServer(
	ctx context.Context, logger2 logger.Logger,

	tracer *opentracing.Tracer,
	rpcServer *rpc.RPCServer, db2 *db.Store,

	accountService *account_application.AccountService,
) (*api.Server, error) {

	API := api.Server{}

	apiService, err := API.Use(ctx, db2, logger2, tracer, accountService)
	if err != nil {
		return nil, err
	}

	return apiService, nil
}

func NewBillingStore(ctx context.Context, logger2 logger.Logger, db2 *db.Store) (*billing_store.BillingStore, error) {
	store := &billing_store.BillingStore{}
	billingStore, err := store.Use(ctx, logger2, db2)
	if err != nil {
		return nil, err
	}

	return billingStore, nil
}

func NewBillingService(
	log logger.Logger,

	httpAPIServer *api.Server,
) (*BillingService, error) {
	return &BillingService{
		Logger: log,

		httpAPIServer: httpAPIServer,
	}, nil
}

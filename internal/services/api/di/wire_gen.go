// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"context"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/services/api/application"
	v1_2 "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	v1_3 "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/sitemap/v1"
	v1_4 "github.com/batazor/shortlink/internal/services/metadata/infrastructure/rpc/metadata/v1"
	"github.com/batazor/shortlink/pkg/rpc"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/text/message"
	"google.golang.org/grpc"
)

// Injectors from di.go:

func InitializeAPIService(ctx context.Context, i18n *message.Printer, runRPCClient *grpc.ClientConn, runRPCServer *rpc.RPCServer, log logger.Logger, tracer *trace.TracerProvider) (*APIService, func(), error) {
	metadataServiceClient, err := NewMetadataRPCClient(runRPCClient)
	if err != nil {
		return nil, nil, err
	}
	linkServiceClient, err := NewLinkRPCClient(runRPCClient)
	if err != nil {
		return nil, nil, err
	}
	linkCommandServiceClient, err := NewLinkCommandRPCClient(runRPCClient)
	if err != nil {
		return nil, nil, err
	}
	linkQueryServiceClient, err := NewLinkQueryRPCClient(runRPCClient)
	if err != nil {
		return nil, nil, err
	}
	sitemapServiceClient, err := NewSitemapServiceClient(runRPCClient)
	if err != nil {
		return nil, nil, err
	}
	api, err := NewAPIApplication(ctx, i18n, log, runRPCServer, metadataServiceClient, linkServiceClient, linkCommandServiceClient, linkQueryServiceClient, sitemapServiceClient)
	if err != nil {
		return nil, nil, err
	}
	apiService, err := NewAPIService(log, i18n, api)
	if err != nil {
		return nil, nil, err
	}
	return apiService, func() {
	}, nil
}

// di.go:

type APIService struct {
	Logger logger.Logger

	// applications
	service *api_application.API
}

// APIService ==========================================================================================================
var APISet = wire.NewSet(

	NewLinkRPCClient,
	NewLinkCommandRPCClient,
	NewLinkQueryRPCClient,
	NewSitemapServiceClient,
	NewMetadataRPCClient,

	NewAPIApplication,

	NewAPIService,
)

func NewLinkRPCClient(runRPCClient *grpc.ClientConn) (v1.LinkServiceClient, error) {
	LinkServiceClient := v1.NewLinkServiceClient(runRPCClient)
	return LinkServiceClient, nil
}

func NewLinkCommandRPCClient(runRPCClient *grpc.ClientConn) (v1_2.LinkCommandServiceClient, error) {
	LinkCommandRPCClient := v1_2.NewLinkCommandServiceClient(runRPCClient)
	return LinkCommandRPCClient, nil
}

func NewLinkQueryRPCClient(runRPCClient *grpc.ClientConn) (v1_2.LinkQueryServiceClient, error) {
	LinkQueryRPCClient := v1_2.NewLinkQueryServiceClient(runRPCClient)
	return LinkQueryRPCClient, nil
}

func NewSitemapServiceClient(runRPCClient *grpc.ClientConn) (v1_3.SitemapServiceClient, error) {
	sitemapRPCClient := v1_3.NewSitemapServiceClient(runRPCClient)
	return sitemapRPCClient, nil
}

func NewMetadataRPCClient(runRPCClient *grpc.ClientConn) (v1_4.MetadataServiceClient, error) {
	metadataRPCClient := v1_4.NewMetadataServiceClient(runRPCClient)
	return metadataRPCClient, nil
}

func NewAPIApplication(
	ctx context.Context,
	i18n *message.Printer, logger2 logger.Logger,

	rpcServer *rpc.RPCServer,

	metadataClient v1_4.MetadataServiceClient,
	link_rpc v1.LinkServiceClient,
	link_command v1_2.LinkCommandServiceClient,
	link_query v1_2.LinkQueryServiceClient,
	sitemap_rpc v1_3.SitemapServiceClient,
) (*api_application.API, error) {

	apiService, err := api_application.RunAPIServer(
		ctx,
		i18n, logger2, rpcServer,

		link_rpc,
		link_command,
		link_query,
		sitemap_rpc,
	)
	if err != nil {
		return nil, err
	}

	return apiService, nil
}

func NewAPIService(
	log logger.Logger,
	i18n *message.Printer,

	service *api_application.API,
) (*APIService, error) {
	return &APIService{
		Logger: log,

		service: service,
	}, nil
}

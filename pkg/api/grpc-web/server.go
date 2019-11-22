package grpcweb

import (
	"context"
	"fmt"
	"github.com/batazor/shortlink/pkg/api"
	"net"
	"net/http"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/store"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// API ...
type API struct { // nolint unused
	store store.DB
	ctx   context.Context
}

// Run HTTP-server
func (api *API) Run(ctx context.Context, db store.DB, config api.Config) error {
	api.ctx = ctx
	api.store = db

	log := logger.GetLogger(ctx)
	log.Info("Run gRPC-GateWay API")

	// Rug gRPC
	go func() {
		err := api.runGRPC()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	gw := runtime.NewServeMux(
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := RegisterLinkHandlerFromEndpoint(ctx, gw, "localhost:9090", opts)
	if err != nil {
		return err
	}

	srv := http.Server{Addr: fmt.Sprintf(":%d", config.Port), Handler: gw}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	err = srv.ListenAndServe()
	return err
}

// runGRPC ...
func (api *API) runGRPC() error {
	lis, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	RegisterLinkServer(grpcServer, api)
	err = grpcServer.Serve(lis)
	return err
}

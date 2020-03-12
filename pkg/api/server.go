package api

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/pkg/api/cloudevents"
	gokit "github.com/batazor/shortlink/pkg/api/go-kit"
	"github.com/batazor/shortlink/pkg/api/graphql"
	grpcweb "github.com/batazor/shortlink/pkg/api/grpc-web"
	httpchi "github.com/batazor/shortlink/pkg/api/http-chi"
	api_type "github.com/batazor/shortlink/pkg/api/type"
)

// runAPIServer - start HTTP-server
func (*Server) RunAPIServer(ctx context.Context, log logger.Logger, tracer opentracing.Tracer) {
	var server API

	viper.SetDefault("API_TYPE", "http-chi")
	viper.SetDefault("API_PORT", 7070)

	config := api_type.Config{
		Port: viper.GetInt("API_PORT"),
	}

	serverType := viper.GetString("API_TYPE")

	switch serverType {
	case "http-chi":
		server = &httpchi.API{}
	case "go-kit":
		server = &gokit.API{}
	case "gRPC-web":
		server = &grpcweb.API{}
	case "graphql":
		server = &graphql.API{}
	case "cloudevents":
		server = &cloudevents.API{}
	default:
		server = &httpchi.API{}
	}

	if err := server.Run(ctx, config, log, tracer); err != nil {
		log.Fatal(err.Error())
	}
}

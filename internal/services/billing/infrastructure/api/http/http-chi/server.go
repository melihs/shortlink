package http_chi

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/api/http/http-chi/controllers/account"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/api/http/http-chi/controllers/balance"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/api/http/http-chi/controllers/order"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/api/http/http-chi/controllers/payment"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/api/http/http-chi/controllers/tariff"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/api/http/type"
	"github.com/batazor/shortlink/pkg/http/handler"
	additionalMiddleware "github.com/batazor/shortlink/pkg/http/middleware"
)

// Run HTTP-server
func (api *API) Run(ctx context.Context, config api_type.Config, log logger.Logger, tracer *opentracing.Tracer) error { // nolint unparam
	api.ctx = ctx
	api.jsonpb = protojson.MarshalOptions{
		UseProtoNames: true,
	}

	log.Info("Run HTTP-CHI API")

	r := chi.NewRouter()

	// CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
		//Debug:            true,
	})

	r.Use(cors.Handler)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// A good base middleware stack
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/healthz"))
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(config.Timeout))

	// Additional middleware
	r.Use(additionalMiddleware.NewTracing(tracer))
	r.Use(additionalMiddleware.Logger(log))

	r.NotFound(handler.NotFoundHandler)

	r.Mount("/api/billing", r.Group(func(router chi.Router) {
		account.Routes(router)
		balance.Routes(router)
		order.Routes(router)
		payment.Routes(router)
		tariff.Routes(router)
	}))

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: r,

		ReadTimeout:       1 * time.Second,                   // the maximum duration for reading the entire request, including the body
		WriteTimeout:      (config.Timeout + 30*time.Second), // the maximum duration before timing out writes of the response
		IdleTimeout:       30 * time.Second,                  // the maximum amount of time to wait for the next request when keep-alive is enabled
		ReadHeaderTimeout: 2 * time.Second,                   // the amount of time allowed to read request headers
	}

	// start HTTP-server
	log.Info(fmt.Sprintf("API run on port %d", config.Port))
	err := srv.ListenAndServe()
	return err
}

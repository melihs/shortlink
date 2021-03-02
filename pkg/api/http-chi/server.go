package httpchi

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

	"github.com/batazor/shortlink/internal/pkg/logger"
	additionalMiddleware "github.com/batazor/shortlink/pkg/api/middleware"
	api_type "github.com/batazor/shortlink/pkg/api/type"
)

// Run HTTP-server
func (api *API) Run(ctx context.Context, config api_type.Config, log logger.Logger, tracer *opentracing.Tracer) error { // nolint unparam
	api.ctx = ctx

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
	r.Use(middleware.Timeout(config.Timeout * time.Second))

	// Additional middleware
	r.Use(additionalMiddleware.NewTracing(tracer))
	r.Use(additionalMiddleware.Logger(log))

	r.NotFound(NotFoundHandler)

	r.Mount("/api", api.Routes())

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: r,

		ReadTimeout:       1 * time.Second,                     // the maximum duration for reading the entire request, including the body
		WriteTimeout:      (config.Timeout + 30) * time.Second, // the maximum duration before timing out writes of the response
		IdleTimeout:       30 * time.Second,                    // the maximum amount of time to wait for the next request when keep-alive is enabled
		ReadHeaderTimeout: 2 * time.Second,                     // the amount of time allowed to read request headers
	}

	// start HTTP-server
	log.Info(fmt.Sprintf("API run on port %d", config.Port))
	err := srv.ListenAndServe()
	return err
}

// NotFoundHandler - default handler for don't existing routers
func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
}

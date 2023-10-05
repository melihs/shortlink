/*
Tracing wrapping
*/
package traicing

import (
	"context"
	"time"

	otelpyroscope "github.com/pyroscope-io/otel-profiling-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
	"github.com/shortlink-org/shortlink/internal/pkg/observability/common"
)

// Init returns an instance of Tracer Provider that samples 100% of traces and logs all spans to stdout.
func Init(ctx context.Context, cnf Config, log logger.Logger) (*trace.TracerProvider, func(), error) {
	// Setup resource.
	res, err := common.NewResource(cnf.ServiceName, cnf.ServiceVersion)
	if err != nil {
		return nil, nil, err
	}

	// Setup trace provider.
	tp, err := newTraceProvider(res)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		err := tp.Shutdown(ctx)
		if err != nil {
			log.Error(`Tracing disable`, field.Fields{
				"uri": cnf.URI,
				"err": err,
			})
		}
	}

	// Register the global Tracer provider
	otel.SetTracerProvider(otelpyroscope.NewTracerProvider(
		tp,
		otelpyroscope.WithAppName(cnf.ServiceName),
		otelpyroscope.WithPyroscopeURL(cnf.URI),
		otelpyroscope.WithRootSpanOnly(true),
		otelpyroscope.WithAddSpanName(true),
		otelpyroscope.WithProfileURL(true),
		otelpyroscope.WithProfileBaselineURL(true),
	))

	// Register the W3C trace context and baggage propagators so data is propagated across services/processes
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	log.Info(`Tracing enable`, field.Fields{
		"uri": cnf.URI,
	})

	return tp, cleanup, nil
}

func newTraceProvider(res *resource.Resource) (*trace.TracerProvider, error) {
	traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter, trace.WithBatchTimeout(5*time.Second)),
		trace.WithResource(res),
	)

	return traceProvider, nil
}
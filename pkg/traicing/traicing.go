package traicing

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

// initJaeger returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func Init() (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: "ShortLink",
		RPCMetrics:  true,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "localhost:6831",
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		return nil, nil, err
	}
	return tracer, closer, nil
}

// WithLogger set logger
func WithTraicer(ctx context.Context, traicer opentracing.Tracer) context.Context {
	return context.WithValue(ctx, KeyTraicer, traicer)
}

// GetLogger return logger
func GetTraicer(ctx context.Context) opentracing.Tracer {
	return ctx.Value(KeyTraicer).(opentracing.Tracer)
}

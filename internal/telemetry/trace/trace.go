package trace

import (
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

// NewTracer creates a new otl tracer.
func NewTracer(cfg Config) (trace.Tracer, error) {
	// create a new jaeger exporter
	exporter, err := jaeger.New(
		jaeger.WithAgentEndpoint(jaeger.WithAgentHost(cfg.Host), jaeger.WithAgentPort(cfg.Port)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize export pipeline: %w", err)
	}

	// generate resources
	res, err := resource.Merge(
		resource.Default(),
		resource.NewSchemaless(
			semconv.ServiceNamespaceKey.String("amirhnajafiz"),
			semconv.ServiceNameKey.String("stallion-load-test"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to manage resourses: %w", err)
	}

	// create a new batch span
	bsp := sdk.NewBatchSpanProcessor(exporter)
	tp := sdk.NewTracerProvider(
		sdk.WithSampler(sdk.ParentBased(sdk.TraceIDRatioBased(cfg.Ratio))),
		sdk.WithSpanProcessor(bsp),
		sdk.WithResource(res),
	)

	// setting the trace provider
	otel.SetTracerProvider(tp)

	// register the TraceContext propagator globally
	var tc propagation.TraceContext

	otel.SetTextMapPropagator(tc)

	// create a new tracer
	tracer := otel.Tracer("stallion-load-test")

	return tracer, nil
}

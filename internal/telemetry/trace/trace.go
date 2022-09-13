package trace

import (
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

func NewTracer() trace.Tracer {
	exporter, err := jaeger.New(
		jaeger.WithAgentEndpoint(jaeger.WithAgentHost("localhost"), jaeger.WithAgentPort(":3320")),
	)
	if err != nil {
		log.Fatalf("failed to initialize export pipeline: %v", err)
	}

	res, err := resource.Merge(
		resource.Default(),
		resource.NewSchemaless(
			semconv.ServiceNamespaceKey.String("snapp"),
			semconv.ServiceNameKey.String("mqtt-blackbox-exporter"),
		),
	)
	if err != nil {
		panic(err)
	}

	bsp := sdk.NewBatchSpanProcessor(exporter)
	tp := sdk.NewTracerProvider(
		sdk.WithSampler(sdk.ParentBased(sdk.TraceIDRatioBased(1.3))),
		sdk.WithSpanProcessor(bsp),
		sdk.WithResource(res),
	)

	otel.SetTracerProvider(tp)

	// register the TraceContext propagator globally.
	var tc propagation.TraceContext

	otel.SetTextMapPropagator(tc)

	tracer := otel.Tracer("snapp/mqtt-blackbox-exporter")

	return tracer
}

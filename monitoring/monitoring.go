package monitoring

import (
	"context"
	"log"
	"os"

	"github.com/dezzles-apps/go-common/monitoring/middleware"
	"github.com/gin-gonic/gin"
	"github.com/hyperdxio/opentelemetry-go/otelzap"
	"github.com/hyperdxio/opentelemetry-logs-go/exporters/otlp/otlplogs"
	sdk "github.com/hyperdxio/opentelemetry-logs-go/sdk/logs"
	"github.com/hyperdxio/otel-config-go/otelconfig"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.uber.org/zap"
)

type Monitoring struct {
	otelShutdown   func()
	ctx            context.Context
	loggerProvider *sdk.LoggerProvider
	Logger         *zap.Logger
}

func (m *Monitoring) newResource() *resource.Resource {
	hostName, _ := os.Hostname()
	version := os.Getenv("DEPLOYED_VERSION")
	if version == "" {
		version = "1.0.0"
	}
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceVersion(version),
		semconv.HostName(hostName),
	)
}

func (m *Monitoring) Initialise() {
	otelShutdown, err := otelconfig.ConfigureOpenTelemetry()
	if err != nil {
		log.Fatalf("Error configuring otel: %s", err.Error())
	}
	m.otelShutdown = otelShutdown

	m.ctx = context.Background()
	logExporter, err := otlplogs.NewExporter(m.ctx)
	if err != nil {
		log.Fatalf("Error configuring OTLP log exporter: %s", err.Error())
	}
	m.loggerProvider = sdk.NewLoggerProvider(
		sdk.WithResource(m.newResource()),
		sdk.WithBatcher(logExporter),
	)

	m.Logger = zap.New(otelzap.NewOtelCore(m.loggerProvider))
	zap.ReplaceGlobals(m.Logger)
}

func (m *Monitoring) ConfigureGinRouter(router *gin.Engine) {
	router.Use(otelgin.Middleware(os.Getenv("OTEL_SERVICE_NAME")))
	router.Use(middleware.WithTraceMetadata(m.Logger))
}

func (m *Monitoring) Shutdown() {
	m.loggerProvider.Shutdown(m.ctx)
	m.otelShutdown()
}

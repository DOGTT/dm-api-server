package server

import (
	"fmt"
	"time"

	"net/http"
	"os"

	"github.com/DOGTT/dm-api-server/internal/conf"
	"github.com/DOGTT/dm-api-server/internal/service"
	"github.com/slok/go-http-metrics/middleware"
	"go.uber.org/zap"

	api "github.com/DOGTT/dm-api-server/api/gin"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	ginmiddleware "github.com/slok/go-http-metrics/middleware/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func NewGinHandler(c *conf.Server, svc *service.Service) http.Handler {
	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil
	r := gin.Default()

	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))

	if c.HTTP.EnableTrace {
		r.Use(otelgin.Middleware("dog-api-server"))
	}
	if c.HTTP.EnableMetric {
		r.Use(ginmiddleware.Handler("", middleware.New(middleware.Config{
			Recorder: metrics.NewRecorder(metrics.Config{}),
		})))
	}

	api.RegisterHandlers(r, svc)

	return r
}

func NewGinServer(c *conf.Server, svc *service.Service) *http.Server {
	s := &http.Server{
		Handler:     NewGinHandler(c, svc),
		ReadTimeout: time.Second * 30,
	}
	return s
}

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

// AuthMiddleware 鉴权中间件
func AuthMiddleware(whitelist []string, svc *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查请求路径是否在白名单中
		for _, path := range whitelist {
			if c.Request.URL.Path == path {
				c.Next() // 如果在白名单中，继续处理请求
				return
			}
		}

		// 从请求中获取 Authorization 标头
		token := c.GetHeader("Authorization")

		if err := svc.Auth(token); err != nil {
			c.JSON(http.StatusUnauthorized,
				gin.H{"error": "unauthorized", "msg": err.Error()})
			c.Abort() // 中止请求
			return
		}

		// 鉴权通过，继续处理请求
		c.Next()
	}
}

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

	r.Use(AuthMiddleware(c.HTTP.AuthWhitePathlist, svc))
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

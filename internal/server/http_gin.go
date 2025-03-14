package server

import (
	"fmt"
	"time"

	"net/http"
	"os"

	"github.com/DOGTT/dm-api-server/internal/conf"
	"github.com/DOGTT/dm-api-server/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/slok/go-http-metrics/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	base_api "github.com/DOGTT/dm-api-server/api/base"
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

		if tc, err := svc.AuthToken(token); err != nil {
			emAPI := &base_api.ErrorMessage{
				Code: err.Code,
				Desc: err.Desc,
			}
			c.JSON(err.HttpStatus, emAPI)
			c.Abort() // 中止请求
			return
		} else {
			c.Set(string(service.TOKEN_CLAIM_KEY), tc)
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

	// 未来可以使用github.com/bytedance/sonic，增强json处理性能。

	// 配置 CORS
	if c.HTTP.EnableCORS {
		r.Use(cors.New(cors.Config{
			// 允许所有来源访问
			AllowOrigins:     []string{"*"},                                       // 可以是特定的 URL 地址，也可以是 "*" 允许所有来源
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的方法
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允许的请求头
			AllowCredentials: true,                                                // 是否允许客户端携带 Cookie
		}))
	}
	if c.HTTP.EnableSwagger {
		r.GET("/openapi.yaml", func(c *gin.Context) {
			c.File("api/openapi/openapi.yaml") // 确保 openapi.yaml 文件在项目根目录下
		})
		_, port := conf.GetAddrSplit(c.HTTP.Addr)
		var fn = func(gc *ginSwagger.Config) {
			gc.URL = fmt.Sprintf("http://localhost:%s/openapi.yaml", port)
		}
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, fn))
	}

	r.Use(AuthMiddleware(c.HTTP.AuthWhitePathlist, svc))
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))

	if c.HTTP.EnableTrace {
		r.Use(otelgin.Middleware("dm-api-server"))
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

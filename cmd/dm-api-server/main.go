package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/DOGTT/dm-api-server/internal/conf"
	"github.com/DOGTT/dm-api-server/internal/data"
	"github.com/DOGTT/dm-api-server/internal/metrics"
	"github.com/DOGTT/dm-api-server/internal/server"
	"github.com/DOGTT/dm-api-server/internal/service"
	"github.com/google/uuid"
	"github.com/jinzhu/configor"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	_ "go.uber.org/automaxprocs"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.uber.org/zap"
)

// Build Info, go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = ""
	// Version is the version of the compiled software.
	Version   string = "0.0.1"
	OSAndArch        = runtime.GOOS + "/" + runtime.GOARCH
	BuildTS          = "None"
	GitHash          = "None"
	GitBranch        = "None"
)

// Flag List, the flag for cli.
var (
	// flagConfFile is the config flag.
	flagConfFile string
	printVersion bool
)

func init() {
	flag.StringVar(&flagConfFile, "conf", "./configs/config.yaml", "config path, eg: -conf config.yaml")
	flag.BoolVar(&printVersion, "version", false, "print version of this build, eg: -version")
	metrics.Init()
}
func ShortenUUID(u uuid.UUID) string {
	return base64.RawURLEncoding.EncodeToString(u[:])
}

// 生成唯一标识符
func generateUniqueID() string {
	// 生成 UUID
	uuid := uuid.New().String()
	// 获取当前时间戳的最后 4 位
	timestamp := time.Now().UnixMicro() % 10000 // 取时间戳的最后 4 位
	// 从 UUID 中提取前 4 位（去掉 '-'）
	uuidPart := uuid[:4]
	// 组合 UUID 的前 4 位和时间戳
	uniqueID := fmt.Sprintf("%s%04d", uuidPart, timestamp)
	return uniqueID
}

func main() {
	uid := uuid.New()
	fmt.Println(uid.String())
	fmt.Println(generateUniqueID())
	fmt.Println(ShortenUUID(uid))

	flag.Parse()
	if printVersion {
		printFullVersionInfo()
		return
	}
	// load config
	c := conf.Config{}
	configLoader := configor.New(&configor.Config{
		Verbose: true,
		Debug:   false,
	})
	if err := configLoader.Load(&c, flagConfFile); err != nil {
		panic(fmt.Sprintf("config:%v", err))
	}
	// set log
	setLogger(c.Log)
	// set metrics
	go runMetricsServer(c.Metric)
	// init data
	dc, err := data.New(c.Data)
	if err != nil {
		panic(fmt.Sprintf("data new error:%v", err))
	}
	// init service
	svc := service.New(c.Service, dc)
	// init server
	srv, err := server.New(c.Server, svc)
	if err != nil {
		panic(fmt.Sprintf("server new error:%v", err))
	}
	// start server and wait for stop signal
	srv.Start()
}

func printFullVersionInfo() {
	fmt.Println("Name:             ", Name)
	fmt.Println("Version:          ", Version)
	fmt.Println("OSAndArch:        ", OSAndArch)
	fmt.Println("Git Branch:       ", GitBranch)
	fmt.Println("Git Commit:       ", GitHash)
	fmt.Println("Build Time (UTC): ", BuildTS)
}

func runMetricsServer(conf *conf.MetricConfig) {
	if conf != nil && conf.Enable {
		http.Handle("/metrics", promhttp.Handler())
		zap.L().Info("Metrics Listening", zap.Any("addr", conf.Addr))
		_ = http.ListenAndServe(conf.Addr, nil)
	}
}

func setLogger(conf *zap.Config) {
	loggerConfig := zap.NewDevelopmentConfig()
	if conf != nil {
		loggerConfig = *conf
	}
	loggerG, err := loggerConfig.Build()
	if err != nil {
		panic(fmt.Sprintf("log_build:%v", err))
	}

	// set global logger
	zap.ReplaceGlobals(loggerG)
	_ = otelzap.ReplaceGlobals(otelzap.New(loggerG))
}

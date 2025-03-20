package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"

	"github.com/DOGTT/dm-api-server/internal/conf"
	"github.com/DOGTT/dm-api-server/internal/data"
	"github.com/DOGTT/dm-api-server/internal/metrics"
	"github.com/DOGTT/dm-api-server/internal/server"
	"github.com/DOGTT/dm-api-server/internal/service"
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

func main() {
	flag.Parse()
	if printVersion {
		printFullVersionInfo()
		return
	}
	// load config
	c := conf.Config{}
	configLoader := configor.New(&configor.Config{
		Verbose: false,
		Debug:   true,
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
	svc, err := service.New(c.Service, dc)
	if err != nil {
		panic(fmt.Sprintf("service new error:%v", err))
	}
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

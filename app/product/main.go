package main

import (
	"net"
	"time"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/czczcz831/tiktok-mall/app/product/biz/dal"
	"github.com/czczcz831/tiktok-mall/app/product/conf"
	"github.com/czczcz831/tiktok-mall/app/product/kitex_gen/product/productservice"
	_ "github.com/joho/godotenv/autoload"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	consul "github.com/kitex-contrib/registry-consul"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	dal.Init()
	opts := kitexInit()

	svr := productservice.NewServer(new(ProductServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))
	// thrift meta handler
	opts = append(opts, server.WithMetaHandler(transmeta.ServerTTHeaderHandler))

	//server registry
	r, err := consul.NewConsulRegisterWithConfig(conf.GetConsulCfg())
	if err != nil {
		klog.Fatalf("new consul register failed: %v", err)
	}
	opts = append(opts, server.WithRegistry(r))
	//Metrics
	opts = append(opts, server.WithTracer(prometheus.NewServerTracer(conf.GetConf().Metrics, "/metrics")))

	// klog
	logger := kitexlogrus.NewLogger()
	//Logstash
	logstashConn, err := net.Dial("tcp", conf.GetConf().Logstash)
	if err != nil {
		klog.Infof("logstash connect error: %v", err)
	} else {
		logger.Logger().WithField("app", conf.GetConf().Kitex.Service)
		hook := logrustash.New(logstashConn, logger.Logger().Formatter)
		logger.Logger().AddHook(hook)
	}
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	// klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}

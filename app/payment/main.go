package main

import (
	"net"
	"time"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/czczcz831/tiktok-mall/app/payment/biz/dal"
	"github.com/czczcz831/tiktok-mall/app/payment/conf"
	"github.com/czczcz831/tiktok-mall/app/payment/kitex_gen/payment/paymentservice"
	_ "github.com/joho/godotenv/autoload"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	consul "github.com/kitex-contrib/registry-consul"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// 初始化配置
	_ = conf.GetConf()

	// 初始化数据访问层
	dal.Init()

	// 初始化服务
	opts := kitexInit()
	svr := paymentservice.NewServer(new(PaymentServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	config := conf.GetConf()

	// address
	addr, err := net.ResolveTCPAddr("tcp", config.Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: config.Kitex.Service,
	}))
	// thrift meta handler
	opts = append(opts, server.WithMetaHandler(transmeta.ServerTTHeaderHandler))

	// server registry
	r, err := consul.NewConsulRegisterWithConfig(conf.GetConsulCfg())
	if err != nil {
		klog.Fatalf("new consul register failed: %v", err)
	}
	opts = append(opts, server.WithRegistry(r))
	//Metrics
	opts = append(opts, server.WithTracer(prometheus.NewServerTracer(config.Metrics, "/metrics")))

	// klog
	logger := kitexlogrus.NewLogger()
	//Logstash
	logstashConn, err := net.Dial("tcp", config.Logstash)
	if err != nil {
		klog.Infof("logstash connect error: %v", err)
	} else {
		logger.Logger().WithField("app", config.Kitex.Service)
		hook := logrustash.New(logstashConn, logger.Logger().Formatter)
		logger.Logger().AddHook(hook)
	}
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   config.Kitex.LogFileName,
			MaxSize:    config.Kitex.LogMaxSize,
			MaxBackups: config.Kitex.LogMaxBackups,
			MaxAge:     config.Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}

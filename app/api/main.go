// Code generated by hertz generator.

package main

import (
	"context"
	"net"
	"time"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/czczcz831/tiktok-mall/app/api/biz/dal"
	"github.com/czczcz831/tiktok-mall/app/api/biz/router"
	"github.com/czczcz831/tiktok-mall/app/api/conf"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/hertz-contrib/pprof"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	address := conf.GetConf().Hertz.Address
	// server registry

	r := conf.GetRegister()

	h := server.New(
		server.WithHostPorts(address),
		server.WithRegistry(r, &registry.Info{
			ServiceName: conf.GetConf().Hertz.Service,
			Addr:        utils.NewNetAddr("tcp", address),
			Weight:      10,
			Tags:        nil,
		}),
	)

	registerMiddleware(h)
	// init dal
	dal.Init()

	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	router.GeneratedRegister(h)

	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	// log
	logger := hertzlogrus.NewLogger()
	//Logstash
	logstashConn, err := net.Dial("tcp", conf.GetConf().Logstash)
	if err != nil {
		hlog.Fatalf("logstash connect error: %v", err)
	}
	logger.Logger().WithField("app", conf.GetConf().Hertz.Service)

	hook := logrustash.New(logstashConn, logger.Logger().Formatter)

	logger.Logger().AddHook(hook)
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	// hlog.SetOutput(asyncWriter)
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		asyncWriter.Sync()
	})

	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())
}

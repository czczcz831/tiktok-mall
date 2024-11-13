package auth

import (
	"net"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/czczcz831/tiktok-mall/app/auth/conf"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	// todo edit custom config
	defaultClient     RPCClient
	defaultDstService = "auth"
	defaultClientOpts = []client.Option{
		// client.WithHostPorts("127.0.0.1:8888"),
	}
	once sync.Once
)

func init() {
	r, err := consul.NewConsulResolver(net.JoinHostPort(conf.GetConf().OsConf.ConsulConf.ConsulHost, conf.GetConf().OsConf.ConsulConf.ConsulPort))

	if err != nil {
		klog.Fatalf("new consul resolver failed: %v", err)
	}

	defaultClientOpts = []client.Option{
		client.WithResolver(r),
	}

	DefaultClient()
}

func DefaultClient() RPCClient {
	once.Do(func() {
		defaultClient = newClient(defaultDstService, defaultClientOpts...)
	})
	return defaultClient
}

func newClient(dstService string, opts ...client.Option) RPCClient {
	c, err := NewRPCClient(dstService, opts...)
	if err != nil {
		panic("failed to init client: " + err.Error())
	}
	return c
}

func InitClient(dstService string, opts ...client.Option) {
	defaultClient = newClient(dstService, opts...)
}

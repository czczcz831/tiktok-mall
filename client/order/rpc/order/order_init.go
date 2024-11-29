package order

import (
	"net"
	"os"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	// todo edit custom config
	defaultClient     RPCClient
	defaultDstService = "order"
	defaultClientOpts = []client.Option{
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithTransportProtocol(transport.TTHeader),
	}
	once sync.Once
)

func init() {
	consulHost := os.Getenv("CONSUL_HOST")
	consulPort := os.Getenv("CONSUL_PORT")
	r, err := consul.NewConsulResolver(net.JoinHostPort(consulHost, consulPort))

	if err != nil {
		klog.Fatalf("new consul resolver failed: %v", err)
	}

	defaultClientOpts = append(defaultClientOpts, client.WithResolver(r))

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

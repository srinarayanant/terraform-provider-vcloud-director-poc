package grpc

import (
	"net/rpc"

	"google.golang.org/grpc"

	"github.com/hashicorp/go-plugin"
	"github.com/srinarayanant/terraform-provider-vcloud-director/go/src/vcd/proto"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "PyCloud",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"PY_PLUGIN": &PyVcloudProviderPlugin{},
}

// PyVcloudProvider is the interface that we're exposing as a plugin.
type PyVcloudProvider interface {
	Login(username string, password string, org string, ip string) (*proto.LoginResult, error)

	IsPresentCatalog(name string) (*proto.IsPresentCatalogResult, error)

	CreateCatalog(name string, description string, shared bool) (*proto.CreateCatalogResult, error)

	DeleteCatalog(name string) (*proto.DeleteCatalogResult, error)
}

// This is the implementation of plugin.Plugin so we can serve/consume this.
// We also implement GRPCPlugin so that this plugin can be served over
// gRPC.
type PyVcloudProviderPlugin struct {
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl proto.PyVcloudProviderServer
}

func (*PyVcloudProviderPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RPCClient{client: c}, nil
}

func (p *PyVcloudProviderPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RPCServer{Impl: p.Impl}, nil
}

func (p *PyVcloudProviderPlugin) GRPCServer(s *grpc.Server) error {
	proto.RegisterPyVcloudProviderServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *PyVcloudProviderPlugin) GRPCClient(c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: proto.NewPyVcloudProviderClient(c)}, nil
}

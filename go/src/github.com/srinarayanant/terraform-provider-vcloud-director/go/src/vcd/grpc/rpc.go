package grpc

import (
	"github.com/srinarayanant/terraform-provider-vcloud-director/go/src/vcd/proto"
	"net/rpc"
)

// RPCClient is an implementation of KV that talks over RPC.
type RPCClient struct{ client *rpc.Client }

// Here is the RPC server that RPCClient talks to, conforming to
// the requirements of net/rpc
type RPCServer struct {
	// This is the real implementation
	Impl proto.PyVcloudProviderServer
}

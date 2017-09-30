package shared

import (

	"net/rpc"

	"google.golang.org/grpc"

	"github.com/hashicorp/go-plugin"
	"../../proto"
	
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


// KV is the interface that we're exposing as a plugin.
type PyVcloudProvider interface {
	Login(username string, password string,org string) (*pyvcloudprovider.LoginResult, error)
	
}

// This is the implementation of plugin.Plugin so we can serve/consume this.
// We also implement GRPCPlugin so that this plugin can be served over
// gRPC.
type PyVcloudProviderPlugin struct {
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl pyvcloudprovider.PyVcloudProviderServer
}

func (*PyVcloudProviderPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &RPCClient{client: c}, nil
}

func (p *PyVcloudProviderPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &RPCServer{Impl: p.Impl}, nil
}



func (p *PyVcloudProviderPlugin) GRPCServer(s *grpc.Server) error {
	pyvcloudprovider.RegisterPyVcloudProviderServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *PyVcloudProviderPlugin) GRPCClient(c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: pyvcloudprovider.NewPyVcloudProviderClient(c)}, nil
}


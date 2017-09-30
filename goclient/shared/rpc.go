package shared

import (
	"net/rpc"
	"../../proto"
)

// RPCClient is an implementation of KV that talks over RPC.
type RPCClient struct{ client *rpc.Client }

func (m *RPCClient) Put(key string, value []byte) error {
	// We don't expect a response, so we can just use interface{}
	var resp interface{}

	// The args are just going to be a map. A struct could be better.
	return m.client.Call("Plugin.Put", map[string]interface{}{
		"key":   key,
		"value": value,
	}, &resp)
}



// Here is the RPC server that RPCClient talks to, conforming to
// the requirements of net/rpc
type RPCServer struct {
	// This is the real implementation
	Impl pyvcloudprovider.PyVcloudProviderServer
}


package shared

import (
	"../../proto"
	"golang.org/x/net/context"
)

// GRPCClient is an implementation of KV that talks over RPC.
type GRPCClient struct{ client pyvcloudprovider.PyVcloudProviderClient }

func (m *GRPCClient) Login(username string, password string,org string) error {
	_, err := m.client.Login(context.Background(), &pyvcloudprovider.TenantCredentials{
		Username:   username,
		Password: password,
		Org: 	 org,
	})
	return err
}



// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl pyvcloudprovider.PyVcloudProviderServer
}



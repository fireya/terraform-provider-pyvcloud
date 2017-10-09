package shared

import (
	"../../proto"
	"golang.org/x/net/context"
)

// GRPCClient is an implementation of KV that talks over RPC.
type GRPCClient struct {
	client pyvcloudprovider.PyVcloudProviderClient
}

func (m *GRPCClient) Login(username string, password string, org string, ip string) (*pyvcloudprovider.LoginResult, error) {
	result, err := m.client.Login(context.Background(), &pyvcloudprovider.TenantCredentials{
		Username: username,
		Password: password,
		Org:      org,
		Ip:       ip,
	})
	return result, err
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl pyvcloudprovider.PyVcloudProviderServer
}

func (m *GRPCServer) Login(
	ctx context.Context,
	req *pyvcloudprovider.TenantCredentials) (*pyvcloudprovider.LoginResult, error) {
	v, err := m.Impl.Login(ctx, req)
	return &pyvcloudprovider.LoginResult{Token: v.Token}, err
}

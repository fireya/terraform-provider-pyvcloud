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

func (m *GRPCClient) IsPresentCatalog(name string) (*pyvcloudprovider.IsPresentCatalogResult, error) {
	result, err := m.client.IsPresentCatalog(context.Background(), &pyvcloudprovider.Catalog{
		Name: name,
	})
	return result, err
}

func (m *GRPCServer) IsPresentCatalog(
	ctx context.Context,
	req *pyvcloudprovider.Catalog) (*pyvcloudprovider.IsPresentCatalogResult, error) {
	v, err := m.Impl.IsPresentCatalog(ctx, req)
	return &pyvcloudprovider.IsPresentCatalogResult{Present: v.Present}, err
}

func (m *GRPCClient) CreateCatalog(name string, description string, shared bool) (*pyvcloudprovider.CreateCatalogResult, error) {
	result, err := m.client.CreateCatalog(context.Background(), &pyvcloudprovider.Catalog{
		Name: name,
	})
	return result, err
}

func (m *GRPCServer) CreateCatalog(
	ctx context.Context,
	req *pyvcloudprovider.Catalog) (*pyvcloudprovider.CreateCatalogResult, error) {
	v, err := m.Impl.CreateCatalog(ctx, req)
	return &pyvcloudprovider.CreateCatalogResult{Created: v.Created}, err

}

func (m *GRPCClient) DeleteCatalog(name string) (*pyvcloudprovider.DeleteCatalogResult, error) {
	result, err := m.client.DeleteCatalog(context.Background(), &pyvcloudprovider.Catalog{
		Name: name,
	})
	return result, err
}

func (m *GRPCServer) DeleteCatalog(
	ctx context.Context,
	req *pyvcloudprovider.Catalog) (*pyvcloudprovider.DeleteCatalogResult, error) {
	v, err := m.Impl.DeleteCatalog(ctx, req)
	return &pyvcloudprovider.DeleteCatalogResult{Deleted: v.Deleted}, err

}

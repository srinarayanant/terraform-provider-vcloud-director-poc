package grpc

import (
	"github.com/srinarayanant/terraform-provider-vcloud-director/go/src/vcd/proto"
	"golang.org/x/net/context"
)

// GRPCClient is an implementation of KV that talks over RPC.
type GRPCClient struct {
	client proto.PyVcloudProviderClient
}

func (m *GRPCClient) Login(username string, password string, org string, ip string) (*proto.LoginResult, error) {
	result, err := m.client.Login(context.Background(), &proto.TenantCredentials{
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
	Impl proto.PyVcloudProviderServer
}

func (m *GRPCServer) Login(
	ctx context.Context,
	req *proto.TenantCredentials) (*proto.LoginResult, error) {
	v, err := m.Impl.Login(ctx, req)
	return &proto.LoginResult{Token: v.Token}, err
}

func (m *GRPCClient) IsPresentCatalog(name string) (*proto.IsPresentCatalogResult, error) {
	result, err := m.client.IsPresentCatalog(context.Background(), &proto.Catalog{
		Name: name,
	})
	return result, err
}

func (m *GRPCServer) IsPresentCatalog(
	ctx context.Context,
	req *proto.Catalog) (*proto.IsPresentCatalogResult, error) {
	v, err := m.Impl.IsPresentCatalog(ctx, req)
	return &proto.IsPresentCatalogResult{Present: v.Present}, err
}

func (m *GRPCClient) CreateCatalog(name string, description string, shared bool) (*proto.CreateCatalogResult, error) {
	result, err := m.client.CreateCatalog(context.Background(), &proto.Catalog{
		Name: name,
	})
	return result, err
}

func (m *GRPCServer) CreateCatalog(
	ctx context.Context,
	req *proto.Catalog) (*proto.CreateCatalogResult, error) {
	v, err := m.Impl.CreateCatalog(ctx, req)
	return &proto.CreateCatalogResult{Created: v.Created}, err

}

func (m *GRPCClient) DeleteCatalog(name string) (*proto.DeleteCatalogResult, error) {
	result, err := m.client.DeleteCatalog(context.Background(), &proto.Catalog{
		Name: name,
	})
	return result, err
}

func (m *GRPCServer) DeleteCatalog(
	ctx context.Context,
	req *proto.Catalog) (*proto.DeleteCatalogResult, error) {
	v, err := m.Impl.DeleteCatalog(ctx, req)
	return &proto.DeleteCatalogResult{Deleted: v.Deleted}, err

}


// impl for CatalogUploadMedia

func (m *GRPCClient) CatalogUploadMedia(mediaInfo proto.CatalogUploadMediaInfo ) (*proto.CatalogUploadMediaResult, error) {
	result, err := m.client.CatalogUploadMedia(context.Background(), &mediaInfo)
	return result, err
}

func (m *GRPCServer) CatalogUploadMedia(
	ctx context.Context,
	req *proto.CatalogUploadMediaInfo) (*proto.CatalogUploadMediaResult, error) {
	v, err := m.Impl.CatalogUploadMedia(ctx, req)
	return &proto.CatalogUploadMediaResult{Created: v.Created}, err

}


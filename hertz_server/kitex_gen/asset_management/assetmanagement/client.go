// Code generated by Kitex v0.5.2. DO NOT EDIT.

package assetmanagement

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	asset_management "orbital_poc/hertz_server/kitex_gen/asset_management"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	QueryAsset(ctx context.Context, req *asset_management.QueryAssetRequest, callOptions ...callopt.Option) (r *asset_management.QueryAssetResponse, err error)
	InsertAsset(ctx context.Context, req *asset_management.InsertAssetRequest, callOptions ...callopt.Option) (r *asset_management.InsertAssetResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kAssetManagementClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kAssetManagementClient struct {
	*kClient
}

func (p *kAssetManagementClient) QueryAsset(ctx context.Context, req *asset_management.QueryAssetRequest, callOptions ...callopt.Option) (r *asset_management.QueryAssetResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryAsset(ctx, req)
}

func (p *kAssetManagementClient) InsertAsset(ctx context.Context, req *asset_management.InsertAssetRequest, callOptions ...callopt.Option) (r *asset_management.InsertAssetResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InsertAsset(ctx, req)
}

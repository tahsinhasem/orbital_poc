// Code generated by Kitex v0.5.2. DO NOT EDIT.
package assetmanagement

import (
	server "github.com/cloudwego/kitex/server"
	asset_management "orbital_poc/kitex_server/kitex_gen/asset_management"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler asset_management.AssetManagement, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

package main

import (
	"context"
	asset_management "orbital_poc/kitex_server/kitex_gen/asset_management"
)

// AssetManagementImpl implements the last service interface defined in the IDL.
type AssetManagementImpl struct{}

type AssetInfo struct {
	ID     string
	Name   string
	Market string
}

var AssetData = make(map[string]AssetInfo, 5)

// QueryAsset implements the AssetManagementImpl interface.
func (s *AssetManagementImpl) QueryAsset(ctx context.Context, req *asset_management.QueryAssetRequest) (resp *asset_management.QueryAssetResponse, err error) {
	// TODO: Your code here...
	// TODO: Your code here...

	ast, exist := AssetData[req.ID]
	if !exist {
		return &asset_management.QueryAssetResponse{
			Exist: false,
		}, nil
	}

	resp = &asset_management.QueryAssetResponse{
		Exist:  true,
		ID:     ast.ID,
		Name:   ast.Name,
		Market: ast.Market,
	}

	return resp, nil
}

// InsertAsset implements the AssetManagementImpl interface.
func (s *AssetManagementImpl) InsertAsset(ctx context.Context, req *asset_management.InsertAssetRequest) (resp *asset_management.InsertAssetResponse, err error) {
	// TODO: Your code here...
	// TODO: Your code here...
	_, exist := AssetData[req.ID]
	if exist {
		return &asset_management.InsertAssetResponse{
			Ok:  false,
			Msg: "the id already exists",
		}, nil
	}

	AssetData[req.ID] = AssetInfo{
		ID:     req.ID,
		Name:   req.Name,
		Market: req.Market,
	}

	return &asset_management.InsertAssetResponse{
		Ok: true,
	}, nil
}

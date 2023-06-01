package asset_api

import (
	"context"
	"fmt"

	asset_api "orbital_poc/hertz_server/biz/model/asset_api"
	"orbital_poc/hertz_server/kitex_gen/asset_management"
	"orbital_poc/hertz_server/kitex_gen/asset_management/assetmanagement"

	clientK "github.com/cloudwego/kitex/client"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// QueryStudent .
// @router asset/query [GET]
func QueryAsset(ctx context.Context, c *app.RequestContext) {
	var err error
	var req asset_api.QueryAssetRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	//add other code here:
	fmt.Println("Recieved Query Request")

	//Use Kitex lient to make rpc call
	client, err := assetmanagement.NewClient("asset", clientK.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err)
	}

	reqRpc := &asset_management.QueryAssetRequest{
		ID: fmt.Sprintf("%d", req.ID),
	}

	respRpc, err := client.QueryAsset(ctx, reqRpc)
	if err != nil {
		panic(err)
	}

	//RPC call made

	if !respRpc.Exist {
		resp := &asset_api.QueryAssetResponse{
			Msg: fmt.Sprintf("No data for asset ID: %d", req.ID),
		}
		c.JSON(200, resp)
		return
	}

	resp := &asset_api.QueryAssetResponse{
		ID:     respRpc.ID,
		Name:   respRpc.Name,
		Market: respRpc.Market,
	}

	c.JSON(consts.StatusOK, resp)
}

// InsertAsset
// @router asset/insert [POST]
func InsertAsset(ctx context.Context, c *app.RequestContext) {
	var err error
	var req asset_api.InsertAssetRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	fmt.Println("Recieved Insert Request")

	client, err := assetmanagement.NewClient("asset", clientK.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err)
	}

	reqRpc := &asset_management.InsertAssetRequest{
		ID:     req.ID,
		Name:   req.Name,
		Market: req.Market,
	}

	respRpc, err := client.InsertAsset(ctx, reqRpc)
	if err != nil {
		panic(err)
	}

	if !respRpc.Ok {
		resp := asset_api.InsertAssetResponse{
			OK:  false,
			Msg: respRpc.Msg,
		}
		c.JSON(200, resp)
		return
	}

	resp := asset_api.InsertAssetResponse{
		OK:  true,
		Msg: respRpc.Msg,
	}

	c.JSON(200, resp)
}

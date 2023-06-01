namespace Go asset.management

struct QueryAssetRequest {
    1: string ID;
}

struct QueryAssetResponse {
    1: bool   Exist;
    2: string ID;
    3: string Name;
    4: string Market;
}

struct InsertAssetRequest {
    1: string ID;
    2: string Name;
    3: string Market;
}

struct InsertAssetResponse {
    1: bool Ok;
    2: string Msg;
}

service AssetManagement {
    QueryAssetResponse queryAsset(1: QueryAssetRequest req);
    InsertAssetResponse insertAsset(1: InsertAssetRequest req);
}

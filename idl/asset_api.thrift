namespace Go API

struct QueryAssetRequest { 
     1: i32 ID (api.query = "id", api.vd = "$ <100; msg: 'id must be less than 10'"); // 
} 

struct QueryAssetResponse { 
     1: string ID; 
     2: string Name; 
     3: string Market; 
     4: string Msg; // Back Information, if not queried, the reason is returized 
} 

struct InsertAssetRequest { 
     1: string ID (api.form = "id"); 
     2: string Name (api.form = "name"); 
     3: string Market (API.FORM = "market"); 
} 

struct InsertAssetResponse { 
    1: bool OK; // Whether to insert successfully 
    2: string Msg; // Return information, if not query, the reason is returned 
} 

service AssetApi { 
    // Query interface: queryStudent 
    // Function: Inquire the student information according to the school number provided in the Query parameter 
    QueryAssetResponse QueryAsset (1: QueryAssetRequest REQ) (api.get = "asset/query"); 
    // Insert interface: InsertStudent 
    // Function: Use the school number as the key, insert the student's information 
    InsertAssetResponse InsertAsset (1: InsertAssetRequest REQ) (api.post = "asset/insert"); 
}
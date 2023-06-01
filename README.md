# Proof of concept (POC) - Orbital

This is a repository for the Tiktok Orbital 2023 POC.


## About
Our POC consists of communication[^3] between one Hertz server and one RPC server. The Hertz server is generated using the `asset_api.thrift` IDL file, while the Kitex server and client are generated using the `asset_management.thrift` IDL file.

The Hertz server lisens to requests at `port 4200` on two exposed endpoints at "/asset/insert" [POST] and "/asset/query" [GET]. Once it recieves an api request, it then forwards the request to the Kitex server (using the internal kitex client built inside the hertz server). The Kitex server sits on `port 8888` and responds to the RPC calls made to it.


## Endpoints
| Endpoint | Method | Description |
| --- | --- | --- |
| /asset/query | GET | Used to query about an asset, with its `id` speficied in the url query section |
| /asset/insert | POST | Used to insert an new asset into the RPC database[^1] . Usage can be inferred from the tutorial below |


## How to use? [^2]
**Step 1:**

Initialize kitex server using the command:
"go run ."
from the "./kitex_server" directory


**Step 2:**

Initialize hertz server using the command:
"go run ."
from the "./hertz_server" directory

**Step 3:**

Send a POST request to the "asset/insert" endpoint by:

```
curl --location --request POST 'http://127.0.0.1:4200/asset/insert' \
--header 'Content-Type: application/json' \
--data-raw '{
"Name": "APPLE",
	"ID": "1",
	"Market": "US"
}'
```


**Step 4:**

Send a GET request to the "asset/query" endpoint by:
```
curl --location --request GET 'http://127.0.0.1:4200/asset/query?id=1'
```


[^1]: database : We are not current using an actual database in the demo. The data structure used is a go splice which acts as a makeshift database.
[^2]: It is assumed that go is already installed in your system
[^3]: Kitex server is using the port "8888" and Hertz server using the port "4200", so please keep these ports free for the demo servers to run. 

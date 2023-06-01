# Proof of concept - Orbital

This is the orbital poc code.


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


**Please note** that kitex server is using the port "8888" and hertz server using the port "4200", so please keep these ports free for the demo servers to run.

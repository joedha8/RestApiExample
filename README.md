# Simple Rest API template


get all dependencies : 
```

dep ensure -v

```


how to run : 
```

go run main.go

```


request header : 
```
Content-Type: application/json
key : abc
time : 10:00

```


request body example : 
```

signatureBody = {"ping_data":"ping"}10:00 = c378ef9c2b72d46928a4eb53e3f631abbf0634649d764b8829a802519c08daa4 (sha256)
body = {"ping_data":"ping"} = eyJwaW5nX2RhdGEiOiJwaW5nIn0= (base64)

body:signatureBody = eyJwaW5nX2RhdGEiOiJwaW5nIn0=:c378ef9c2b72d46928a4eb53e3f631abbf0634649d764b8829a802519c08daa4

```

curl example : 

```

curl -d "eyJwaW5nX2RhdGEiOiJwaW5nIn0=:c378ef9c2b72d46928a4eb53e3f631abbf0634649d764b8829a802519c08daa4" -H "Content-Type: application/json" -H  "key: abc" -H "time: 10:00" -X POST http://localhost:8000/api/v1/ping

```
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

hashWithHMAC(key, '{"ping_data":"ping"}' + time)

```

curl example : 

```

curl -d "c378ef9c2b72d46928a4eb53e3f631abbf0634649d764b8829a802519c08daa4" -H "Content-Type: application/json" -H  "key: abc" -H "time: 10:00" -X POST http://localhost:8000/api/v1/ping

```
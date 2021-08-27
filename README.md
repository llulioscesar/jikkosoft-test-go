# jikkosoft-test-go

create container docker postgres
```
  docker-compose up -d
```

Create database and data
```
  cd cmd/
  go run main_initdb.go
```

Run app
```
  cd cmd/
  go run main.go
```

URL
```
  POST: http://localhost:8080/array
  body {
    "unsorted": [0,4,2,2,10,5,5,5,2]
  }
  
  GET: http://localhost:8080/user
```

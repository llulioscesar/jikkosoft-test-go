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

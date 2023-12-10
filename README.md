# Duplo Project Assignment

## Setting up application

### Configuring mysql database using docker
```
docker build -t mysql-forecast -f ./mysqlDockerfile . 

docker run -p 3306:3306  mysql-forecast
```

### Prerequisite

Should have go installed

### Building go application 

#### downloading or verifying go packages
```
go mod download 
   or
go mod tidy
```


To get test coverage
```
go test -coverprofile=coverage.out ./...

go tool cover -func=coverage.out
```

#### Running go application
```
go build -o main .

./main
```
Server runs on 3001 port

Accessing Swagger
```
http://localhost:3001/docs/index.html#
```

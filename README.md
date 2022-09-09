# Article API

## Prerequisites

- go lang version 1.19
- docker/docker-compose

## Development

```
go mod download
go run main.go
```

Find the swagger file in http://localhost:3000/swagger/index.html

use 
```
$GOPATH/bin/swag init
or
swag init
```
to generate swagger files

to start the application using docker-compose

run

```
docker-compose up
```

to stop, use

```
docker-compose down
```
# Heytel-api
Heytel-api is api for the Heytel project. It is written in Golang using Gin framework. It is a RESTful API that provides endpoints for the Heytel project. 

## Installation

```bash
git clone https://github.com/heytei/heytel-api.git
cd heytel-api
go get
go build
./heytel-api
```
To run development version use `go run main.go` instead of `go build` and `./heytel-api`.

## Usage


## Other informations
### Swagger
To run swagger install binaries from https://github.com/go-swagger/go-swagger/releases, then open terminal and run `swagger_windows_amd64.exe serve Heytel-api/docs/swagger.json` 

### Postgresql
Run this command in psql/pgadmin sql query tool `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";` it's needed to get uuid to work.


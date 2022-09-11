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

## Running on local domain with ssl

To run the api on a local domain with ssl  change `.env` file, then set it to:
```
SSL_CRT_FILE=".cert/api.heytel.local.pem"
SSL_KEY_FILE=".cert/api.heytel.local-key.pem"
WEBSERVER_SSL=true
```
Install [mkcert](https://github.com/FiloSottile/mkcert) and run the following commands:


Then run the following commands to apply ssl:
```bash
mkcert -install
mkcert -cert-file ./certs/api.heytel.local.pem -key-file ./certs/api.heytel.local-key.pem localhost api.heytel.local
```

Then you need to add the following lines to your hosts file:
```
127.0.0.1 api.heytel.local
```

`hosts` file location depends on your OS:
- Windows: `C:\Windows\System32\drivers\etc\hosts`
- Linux: `/etc/hosts`
- Mac: `/private/etc/hosts`


## Other informations

### Swagger
To run swagger install binaries from https://github.com/go-swagger/go-swagger/releases, then open terminal and run `swagger_windows_amd64.exe serve Heytel-api/docs/swagger.json` 

### Postgresql
Run this command in psql/pgadmin sql query tool `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";` it's needed to get uuid to work.


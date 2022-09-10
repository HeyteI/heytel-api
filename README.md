# Heytel-api
[![Build Status](https://travis-ci.org/heytel/heytel-api.svg?branch=master)](https://travis-ci.org/heytel/heytel-api)
[![Coverage Status](https://coveralls.io/repos/github/heytel/heytel-api/badge.svg?branch=master)](https://coveralls.io/github/heytel/heytel-api?branch=master)
[![Code Climate](https://codeclimate.com/github/heytel/heytel-api/badges/gpa.svg)](https://codeclimate.com/github/heytel/heytel-api)
[![Dependency Status](https://gemnasium.com/badges/github.com/heytel/heytel-api.svg)](https://gemnasium.com/github.com/heytel/heytel-api)
[![Inline docs](http://inch-ci.org/github/heytel/heytel-api.svg?branch=master)](http://inch-ci.org/github/heytel/heytel-api)

## Description
This is the API for the Heytel project. It is written in Gin and uses Gorm for database access.

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


### How API should work


### TODO
while creating new room, available is always false !BUG
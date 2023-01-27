## Description

Cosmart software engineer assignment

### Prerequisite

- Go (Golang)  
  See [Golang Installation](https://golang.org/doc/install)

- MySQL  
  See [MySQL Installation](https://www.mysql.com/downloads/)

### Installation

1. Clone this repository and install the prerequisites above
2. Copy `.env` from `env.sample` and modify the configuration value appropriately
3. install dependencies with `make dep`
4. test with `make test`
4. build and run with `make run`

### Migration

1. Make sure env is configured correctly
2. migrate the schema with `make migrate`
3. seed the database with data from openlibrary with `make seed`

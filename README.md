# 🐘 PREST demo

Application to demonstrate [prest](https://prestd.com/) features.

## Requirements

For the demo we need a postgres instance with a database demo and user postgres.

Using docker you can do:

```
docker run --name=pgdemo --rm -e POSTGRES_USER=postgres -e POSTGRES_DB=demo -e POSTGRES_HOST_AUTH_METHOD=trust -p 5432:5432 -d postgres:15-alpine
```

## Installation

Using go install:

```sh
go install github.com/prest/prest/cmd/prestd@latest
```

with Homebrew:

```sh
brew install prestd
```

with Docker:

```sh
docker run -d -p 3000:3000 \
    -e PREST_PG_URL=postgres://username:password@hostname:port/dbname \
    -e PREST_DEBUG=true \
    prest/prest:v1
```

or download latest version executable:

https://github.com/prest/prest/releases

## Configuration

For this demo, we are using this configuration:

```toml
# prest.toml
migrations = "./migrations"
debug = true

[pg]
user = "postgres"
database = "demo"

[jwt]
default = false

[ssl]
mode = "disable"
```

More details at [Docs](https://docs.prestd.com/prestd/deployment/server-configuration/#toml).

## Environment variables

We can configure using envinroment variables too.

```sh
PREST_DEBUG=true
PREST_MIGRATIONS="./migration"
PREST_PG_USER="postgres"
PREST_PG_DATABASE="demo"
PREST_JWT_DEFAULT=false
PREST_SSL_MODE="disable"
```

More details at [Docs](https://docs.prestd.com/prestd/deployment/server-configuration/#environment-variables).

## Demo

### Migrations

To create a new migration we use the command `prestd migrate create migration_file_xyz`. For this demo we alread create 2 migrations and they can found in [migrations folder](./migrations/).

```sh
# apply next first migration
prestd migrate next +1

# show the current migration version
prestd migrate version

# rollback all migrations
prestd migrate down

# apply all available migrations
prestd migrate up
```

Others commands can be found at the [docs](https://docs.prestd.com/prestd/deployment/migrations/).

### CRUD

To **create** a new person:

```sh
curl -X POST "localhost:3000/demo/public/person" -d '{"name": "cassio", "age": 32}'
curl -X POST "localhost:3000/demo/public/person" -d '{"name": "igor", "age": 36}'
```

**Read** person:

```sh
curl -X GET "localhost:3000/demo/public/person"
```

or filtering data:

```sh
curl -X GET "localhost:3000/demo/public/person?name=cassio"
```

**Update**

Complete update:

```sh
curl -X PUT "localhost:3000/demo/public/person?name=igor" -d '{"name": "igor", "age": 37}'
```

or partial update:

```sh
curl -X PATCH "localhost:3000/demo/public/person?name=igor" -d '{"age": 38}'
```

**Delete**

```sh
curl -X DELETE "localhost:3000/demo/public/person?name=igor"
```

More details at the [docs](https://docs.prestd.com/prestd/api-reference/endpoints/).

## Query statements

Details can be found [here](https://docs.prestd.com/prestd/api-reference/parameters/#operators).

Add four new person to demonstrate these features.

```sh
curl -X POST "localhost:3000/demo/public/person" -d '{"name": "thiago", "age": 31}'
curl -X POST "localhost:3000/demo/public/person" -d '{"name": "luisa", "age": 29}'
curl -X POST "localhost:3000/demo/public/person" -d '{"name": "pedro", "age": 28}'
curl -X POST "localhost:3000/demo/public/person" -d '{"name": "lucas", "age": 18}'
```
### Equal

```sh
curl -X GET "localhost:3000/demo/public/person?name=cassio"
```

or

```sh
curl -X GET "localhost:3000/demo/public/person?name=\$eq.cassio"
```

### Null or not null

```sh
curl -X GET "localhost:3000/demo/public/person?name=\$null"
curl -X GET "localhost:3000/demo/public/person?name=\$notnull"
```

### True or False

```sh
curl -X GET "localhost:3000/demo/public/person?status=\$true"
curl -X GET "localhost:3000/demo/public/person?status=\$false"
```

### Like

```sh
curl -X GET "localhost:3000/demo/public/person?name=\$like.cassio"
```

## Other operators

Other example of queries are:

```sh
# query person with age greater than 30
curl -X GET "localhost:3000/demo/public/person?age=\$gt.30"
# query person with age less than or equal 30
curl -X GET "localhost:3000/demo/public/person?age=\$lte.30"
```

More details at the [docs](https://docs.prestd.com/prestd/api-reference/parameters/#operators).

## Table permissions

For default the pREST will serve in publish mode, making all tables and views visible.
But using prest.toml you can configure read/write/delete permissions of each table.

Add these lines in prest.toml, reset server and see the diference.

```toml
[access]
restrict = true

[[access.tables]]
name = "person"
permissions = ["read", "write", "delete"]
fields = ["name", "age"]
```

Remove write permission and try to add new person. This should not be allowed.

More details at the [docs](https://docs.prestd.com/prestd/deployment/permissions/#table-permissions).

## Batch insert

Let's try to insert 1000 records at once:

```sh
curl -X POST "localhost:3000/batch/demo/public/person" \
   -d @10000records.json -H "Prest-Batch-Method: copy"
```

If everything went well, you can see the records:

```sh
curl -X GET "localhost:3000/demo/public/person"
```

## Plugin

Add new line in `prest.toml`:

```toml
pluginpath = "./plugins"
```

Build `ping.go` as a go plugin:

```sh
go build -o plugins/ping.so -buildmode=plugin ping.go
```

Restart the server and test plugin integration.

```sh
curl -X GET "http://localhost:3000/_PLUGIN/ping/Ping"
```

More details at the [docs](https://docs.prestd.com/prestd/api-reference/plugins/).

## Conclusion

I hope you enjoy this demo.

There a lot of more features like [scripts](), [templates](), [cors](), etc. Try it yourself!

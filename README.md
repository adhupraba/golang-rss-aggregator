1. go get github.com/joho/godotenv
2. go mod tidy
3. go mod vendor

> vendor folder is like node_modules but it is ok to commit it

# migrations

we use goose to handle migrations for our database. to install the goose cli tool,

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

to migrate up or down our migrations,

`cd sql/schema`

to migrate up

```bash
goose postgres postgresql://postgres:postgres@localhost:5432/rss_aggregator up
```

to migrate down

```bash
goose postgres postgresql://postgres:postgres@localhost:5432/rss_aggregator down
```

# queries

we use `sqlc` to handle db queries. it takes raw sql statements and does codegen to generate typesafe go code which can be used in to program to query the database

to install sqlc,

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

be in the project root directory and run,

```bash
sqlc generate
```

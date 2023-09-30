# FLows Microservice - Notification Service

Notification Service for Flows Microservice

---

## Tech Stacks

- Go

---

## Development

- Copy `config.example.yml` to `config.yml`, then adjust your env configuration

- Use [go-migrate](https://github.com/golang-migrate/migrate) to init migration (if not exists):

```bash
migrate create -ext sql -dir pkg/db/migration -seq <migration_name>
```

- To run migration up:

```bash
migrate -path pkg/db/migration -database "postgresql://<user>:<password>@<host>:5432/<db_name>?sslmode=disable" -verbose up
```

- For migration down:

```bash
migrate -path pkg/db/migration -database "postgresql://<user>:<password>@<host>:5432/<db_name>?sslmode=disable" -verbose down
```

- Run server:

```bash
go run cmd/api.go
```

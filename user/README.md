# Spendiy Microservice - User Service

User Service for Spendify Microservice

---

## Tech Stacks

- Go
- Postgres

---

## Development

- Copy `config.example.yml` to `config.yml`, then adjust your env configuration

- Run server:

```bash
go run cmd/api.go
```

- To crete database migration you need to install [migrate](https://github.com/golang-migrate/migrate) on your local machine

```bash
migrate create -ext sql -dir pkg/db/migration -seq <migration_name>
```

- Then run migration:

```bash
migrate -path pkg/db/migration -database "postgresql://postgres:postgres@localhost:5432/flows_users?sslmode=disable" -verbose up
```

- Rollback using:

```bash
migrate -path pkg/db/migration -database "postgresql://postgres:postgres@localhost:5432/flows_users?sslmode=disable" -verbose down
```

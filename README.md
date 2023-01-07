# Expense Tracking API

GoLang REST API. Responsibility for CRUD expense transaction from authorize client.
This is assessment from the [requirements](docs/requirements/index.md) of GO Software Engineering Bootcamp.
this include 4 usecases 
- [Create new expense](docs/requirements/index.md#Story-EXP01)
- [Update exist expense](docs/requirements/index.md#Story-EXP02)
- [Get expense by id](docs/requirements/index.md#Story-EXP03)
- [Get all expense](docs/requirements/index.md#Story-EXP04)


## Build
```bash
go build server.go
```

## Test
Unit Test
```bash 
go test --tags=unit -v ./...
```
Integration Test with docker
```bash 
docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit --exit-code-from expense_integration_test && \
docker-compose -f docker-compose.test.yml down
```

## Run
Run normal
```bash
go mod tidy 

PORT=2565 \
DATABASE_URL=postgres://vdvbjiod:QufYi3doL10wTzl9Cyi3LfYDnHQYBhna@tiny.db.elephantsql.com/vdvbjiod \
go run server.go
```
Run with docker
```bash
docker build -t expenses-api . \ &&
docker run --rm \
-e PORT=2565 \
-e DATABASE_URL=postgres://vdvbjiod:QufYi3doL10wTzl9Cyi3LfYDnHQYBhna@tiny.db.elephantsql.com/vdvbjiod \
-p 2565:2565 \
expenses-api
```

## Component Environment
```
PORT: is port to serve
DATABASE_URL: connection string of PostgreSQL
```


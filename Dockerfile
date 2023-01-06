FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

CMD CGO_ENABLED=0 go test --tags=unit -v ./...

RUN go build -o expenses-api .

FROM alpine:3.16.3

EXPOSE 2527

WORKDIR /app

COPY --from=builder /app/expenses-api /app

CMD [ "/app/expenses-api" ]
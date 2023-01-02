FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o expense-api .

FROM alpine:3.16.3

EXPOSE 2527

WORKDIR /app

COPY --from=builder /app/expense-api /app

CMD [ "/app/expense-api" ]
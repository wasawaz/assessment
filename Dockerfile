# Declare parent image for build
FROM golang:1.19-alpine as builder

# Create app directory
WORKDIR /app

# Prepare source code for build
COPY . .

# Install dependencies
RUN go mod tidy

# Run Unit testing
CMD CGO_ENABLED=0 go test --tags=unit -v ./...

# Build
RUN go build -o expenses-api .


# Declare parent image for run
FROM alpine:3.16.3 as run
# Export required port
EXPOSE 2527

# Create running app directory
WORKDIR /app

# Copy artifact from build stage to running app directory
COPY --from=builder /app/expenses-api /app

# Run app
CMD [ "/app/expenses-api" ]
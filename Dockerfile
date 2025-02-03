# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /DDD-api
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

# Runtime stage
FROM alpine:3.18
WORKDIR /DDD-api
COPY --from=builder /DDD-api/main ./main
EXPOSE 8080
CMD ["./main", "8080"]
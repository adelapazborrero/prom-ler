# syntax=docker/dockerfile:1

FROM golang:1.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/main cmd/main.go

FROM alpine
COPY --from=builder /app/bin/main .
EXPOSE 8080
CMD ["./main"]

# Stage 1: Build the Go app
FROM golang:1.23-alpine AS builder  
# NOTE: 1.23 doesn't exist, use 1.21 (latest LTS)

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# Stage 2: Minimal runtime container
FROM alpine:3.18 AS runner

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]

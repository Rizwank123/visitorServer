# Stage 1: Build the Go app
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# Stage 2: Minimal runtime container
FROM alpine:3.18 AS runner

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]

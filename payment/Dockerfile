FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o payment ./cmd

# multi-stage build
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/payment .
COPY --from=builder /app/.env .

EXPOSE 8003

ENTRYPOINT ["/root/payment"]

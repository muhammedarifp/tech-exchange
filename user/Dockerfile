FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download


COPY . .

RUN go build -o main ./cmd/main.go

EXPOSE 8000

CMD [ "./main" ]

# # Multi stage building

# FROM alpine:latest

# WORKDIR /app

# COPY --from=builder /app/main .

# EXPOSE 8000

# CMD [ "main" ]
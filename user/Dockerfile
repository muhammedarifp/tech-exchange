FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

# Multi stage building

FROM alpine:latest

EXPOSE 8000

CMD [ "./main" ]
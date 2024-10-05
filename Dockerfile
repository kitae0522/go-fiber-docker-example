FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go run github.com/steebchen/prisma-client-go prefetch
COPY . ./

RUN go run github.com/steebchen/prisma-client-go generate

RUN go build -o app cmd/main.go

CMD ["./app"]

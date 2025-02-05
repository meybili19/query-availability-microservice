FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main .

FROM ubuntu:latest

WORKDIR /app

RUN apt-get update && apt-get install -y libc6

COPY --from=builder /app/main .

EXPOSE 6005

CMD ["./main"]

FROM golang:1.23


WORKDIR /app


COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 7004

CMD ["go", "run", "main.go"]


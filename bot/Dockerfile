FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o bot .

EXPOSE 8443

CMD ["./bot"]

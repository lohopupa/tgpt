FROM golang:latest

WORKDIR /app

ARG TG_WH_CERT_FOLDER

COPY . .

COPY ${TG_WH_CERT_FOLDER}/* ./

RUN go mod download

RUN go build -o bot .

EXPOSE 8443

CMD ["./bot"]

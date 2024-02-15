FROM golang:latest

WORKDIR /app

ARG TG_WH_CERT_FOLDER
ENV TG_WH_CERT_FOLDER_E=$TG_WH_CERT_FOLDER

COPY . .

RUN go mod download

RUN go build -o bot .

EXPOSE 8443

CMD ["./bot"]

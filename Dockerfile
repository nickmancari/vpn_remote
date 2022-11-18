FROM ubuntu:lastest
FROM golang:latest

RUN apt-get -y update && apt-get -y install openvpn sudo

WORKDIR /var/tmp/

COPY go.* ./

RUN go mod download

COPY *.go ./
COPY . ./

RUN go build -o /remote

EXPOSE 8080

ENTRYPOINT ["/remote"]


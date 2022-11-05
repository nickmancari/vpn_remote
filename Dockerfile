FROM golang:latest

WORKDIR /var/tmp/

COPY go.* ./

RUN go mod download

COPY *.go ./
COPY . ./

RUN go build -o /remote

EXPOSE 8080

ENTRYPOINT ["/remote"]


FROM golang:1.17-buster

RUN apt-get update && apt-get install default-mysql-client -y
RUN apt-get install default-mysql-client

WORKDIR /go/src

COPY go.mod .
COPY go.sum .

COPY . .

RUN go get

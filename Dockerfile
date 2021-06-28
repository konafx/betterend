FROM golang:1.16-buster
WORKDIR /go/src

RUN apt update && apt install git

COPY go.mod go.sum ./
RUN go mod download

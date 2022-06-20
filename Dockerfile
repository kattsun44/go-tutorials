FROM golang:latest

RUN apt update && apt install -y vim
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app

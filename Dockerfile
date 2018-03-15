FROM golang:onbuild

MAINTAINER Florian Krone

RUN apt-get update -y

RUN apt-get upgrade -y

RUN go build -o main .

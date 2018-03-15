FROM golang:onbuild

MAINTAINER Florian Krone

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]

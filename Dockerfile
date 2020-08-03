FROM golang:1.14.6 as base

WORKDIR /app/

ADD . /app

RUN make build-server

ENTRYPOINT [ "/app/bin/server" ]

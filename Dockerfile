FROM golang:1.17 as base

RUN mkdir -p /go/src/github.com/learnyougo

WORKDIR /go/src/github.com/learnyougo

COPY . .

CMD go build -o learnyougo

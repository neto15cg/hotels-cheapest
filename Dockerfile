FROM golang:1.16-alpine

ENV PATH="$PATH:/bin/bash" \
    GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add --update --upgrade curl unzip bash gcc g++

WORKDIR /go/src

EXPOSE 8081

ENTRYPOINT ["tail", "-f", "/dev/null"]
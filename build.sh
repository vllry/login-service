#!/usr/bin/env bash
set -e

go get -v
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o app *.go

docker build -t $1 .
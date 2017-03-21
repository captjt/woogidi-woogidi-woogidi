#!/bin/bash
# build.sh

glide install

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -tags netgo

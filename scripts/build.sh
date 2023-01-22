#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o random-identity-server ./cmd/main.go
cp random-identity-server docker

docker build -t nightmanager/random-identity-http-server ./docker/

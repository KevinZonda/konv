#!/bin/bash

# restore dependencies
# go mod download

# build
go build -v -o ./out/konv -ldflags "-s -w" ./cmd/konv

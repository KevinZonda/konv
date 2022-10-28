#!/bin/bash

# restore dependencies
go mod download

# build
go build -v -o ./out/apt -ldflags "-s -w" ./cmd/apt-pac

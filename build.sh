#!/bin/bash
# build for linux amd64 environment same as tencent scf platform
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -extldflags '-static'" -o main main.go
#GOOS=linux GOARCH=amd64 go build -o BuildRelease/main main.go
# zip file for uploading
zip main.zip main
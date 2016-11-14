#!/bin/bash

set -e

echo Building...
GOOS=windows GOARCH=386 go build -ldflags="-s -w" -v -o $1

echo Packing...
upx -qq -9 $1

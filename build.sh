#! /bin/sh

GOOS=darwin go build -ldflags="-s -w" -o $1.temp *.go
upx -f --brute -o $1 $1.temp
rm -rf $1.temp

ls -l $1*
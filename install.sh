#!/usr/bin/env bash

rm -rf bin/
rm -rf pkg/

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

gofmt -w src
go install -gcflags "-N" wxspider

#export GOPATH="$OLDGOPATH"
echo 'finished'


#!/bin/bash -eu

PATH=$PATH:$GOPATH/bin
protdir=../pb

protoc --go_out=plugins=grpc:genproto -I $protdir $protdir/demo.proto
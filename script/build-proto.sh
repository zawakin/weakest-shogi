#!/bin/sh

cd "$(dirname "$0")/.."

rm -rf src/usi/*.pb.go

protoc -I=./proto --go_out=./src/usi proto/usi.proto

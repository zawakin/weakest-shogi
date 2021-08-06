#!/bin/sh

cd "$(dirname "$0")/../src"

go build -o ../bin/weakest-shogi ./*.go

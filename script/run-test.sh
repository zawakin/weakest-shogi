#!/bin/sh

cd "$(dirname "$0")/../src"

go test ./... -count=1

.PHONY: clean proto build test

all: clean proto build

SRC_DIR = ${PWD}/proto
DST_DIR = ${PWD}/usi
proto:
	protoc -I=${SRC_DIR} --go_out=${DST_DIR} ${SRC_DIR}/usi.proto

clean:
	rm ${DST_DIR}/*.pb.go

build:
	go build -o weakest-shogi ./*.go

test:
	go test ./... -count=1

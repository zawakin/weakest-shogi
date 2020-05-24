.PHONY: clean proto

all: clean proto

SRC_DIR = ${PWD}/proto
DST_DIR = ${PWD}/usi
proto:
	protoc -I=${SRC_DIR} --go_out=${DST_DIR} ${SRC_DIR}/usi.proto

clean:
	rm ${DST_DIR}/*.pb.go

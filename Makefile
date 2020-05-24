proto:
	protoc -I./proto \
		--go_out=plugins=grpc:usi \
		proto/*.proto
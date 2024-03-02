.PHONY: gen help

PROTO_DIR=pb

protoc:
	protoc \
	--proto_path=$(PROTO_DIR) \
	--go_out=$(PROTO_DIR) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(PROTO_DIR) \
	--go-grpc_opt=paths=source_relative \
	 book/book.proto

	

help:
	@echo "使用make gen生成go代码 "
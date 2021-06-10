.PHONY: gen_proto

gen_proto:
	protoc --proto_path=proto --go_out=./proto --go_opt=paths=source_relative --go-grpc_out=./proto --go-grpc_opt=paths=source_relative ports.proto

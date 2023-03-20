NAME=auth-component

PROTOBUF_MESSAGE=$(shell find app/protobuf/message -type f -name "*.proto")
PROTOBUF_FUNC=$(shell find app/protobuf/rpc -type f -name "*.proto")

grpc: grpc_models grpc_rpc grpc_moves

grpc_moves: grpc_model_move grpc_rpc_move

grpc_models:
	protoc --go_out=. --go_opt=paths=source_relative app/protobuf/message/hello.proto

grpc_model_move:
	mv app/protobuf/message/*.go app/proto-gen/message/

grpc_rpc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative app/protobuf/rpc/hello.proto


grpc_rpc_move:
	mv app/protobuf/rpc/*.go app/proto-gen/rpc/

build:
	go build app
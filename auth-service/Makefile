NAME=auth-component

PROTOBUF_MESSAGE=$(shell find app/protobuf/message -type f -name "*.proto")
PROTOBUF_FUNC=$(shell find app/protobuf/rpc -type f -name "*.proto")

grpc: grpc_models grpc_rpc grpc_moves

grpc_moves: grpc_model_move grpc_rpc_move

grpc_models:
	@echo "Generate proto message: " \
	&& protoc --go_out=. --go_opt=paths=source_relative app/protobuf/message/hello.proto \
	&& echo "  ===> success !"

grpc_model_move:
	@echo "Move proto message: " \
	&& mv app/protobuf/message/*.go app/proto-gen/message/ && echo "===> ok"

grpc_rpc:
	@echo "Generate proto rpc: " \
	&& protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative app/protobuf/rpc/hello.proto \
	&& echo "  ===> success !"


grpc_rpc_move:
	@echo "Move proto rpc: " \
	&& mv app/protobuf/rpc/*.go app/proto-gen/rpc/ && echo "===> ok"

build:
	@echo "AUTH SERVICE BUILD: "
	@echo "Create binary: " && mkdir -p bin && echo "  ===> OK"
	@echo "GO BUILD: " && go build  -o bin/auth_service app/internal/server/*.go && echo " ===> BUILD SUCCESS"
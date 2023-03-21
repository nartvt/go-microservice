NAME=auth-service
BIN_NAME="auth_service"
DATE_NOW=$(shell date +%Y_%m_%d_%m_%S)

deploy: backup build docker_build

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

backup:
	@echo "Backup Current Binary:" \
	&& if [ -e bin/${BIN_NAME} ]; \
	then mv bin/${BIN_NAME} bin/${BIN_NAME}_${DATE_NOW} && echo " ==> OK"; \
	else echo "File does not exist"; fi

build:
	@echo "auth service build: "
	@echo "Create binary: " && mkdir -p bin && echo "  ===> OK"
	@echo "GO BUILD: " && go build  -o bin/${BIN_NAME} app/internal/server/*.go && echo " ===> BUILD SUCCESS"

docker_build:
	@echo "auth service docker image build: " \
	&& docker build -t nartvt/${BIN_NAME}_${DATE_NOW} .  && echo " ==> docker image build success" \
	&& docker scan  --version nartvt/${BIN_NAME}_${DATE_NOW}
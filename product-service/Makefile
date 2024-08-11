NAME=product-service
BIN_NAME="product_service"
DATE_NOW=$(shell date +%Y_%m_%d_%m_%S)

PROTOBUF_MESSAGE=$(shell find app/protobuf/message -type f -name "*.proto")
PROTOBUF_RPC=$(shell find app/protobuf/rpc -type f -name "*.proto")
PROTO_GEN_MESSAGE="app/proto-gen/message/"
PROTO_GEN_RPC="app/proto-gen/rpc/"

deploy: backup build docker_build

grpc: grpc_models grpc_rpc grpc_moves

grpc_moves: grpc_model_move grpc_rpc_move

grpc_models:
	@echo "Generate proto message: "
	@$(foreach dir,$(PROTOBUF_MESSAGE), \
	 	protoc --go_out=. --go_opt=paths=source_relative $(dir) && echo "Generate $(dir) ===> success !"; \
	)

grpc_model_move:
	@echo "Move proto message: " \
	&& mv app/protobuf/message/*.go ${PROTO_GEN_MESSAGE} && echo "===> ok"

grpc_rpc:
	@echo "Generate proto rpc: "
	@$(foreach dir,$(PROTOBUF_RPC), \
		protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative $(dir) && echo "Generate $(dir) ===> success !" ; \
	)

grpc_rpc_move:
	@echo "Move proto rpc: " \
	&& mv app/protobuf/rpc/*.go ${PROTO_GEN_RPC} && echo "===> ok"

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

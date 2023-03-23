NAME=api-gateway
BIN_NAME="api_gateway"
DATE_NOW=$(shell date +%Y_%m_%d_%m_%S)

deploy: backup build docker_build

backup:
	@echo "Backup Current Binary:" \
	&& if [ -e bin/${BIN_NAME} ]; \
	then mv bin/${BIN_NAME} bin/${BIN_NAME}_${DATE_NOW} && echo " ==> OK"; \
	else echo "File does not exist"; fi

build:
	@echo ${DATE_NOW}
	@echo "${NAME} build: "
	@echo "Create binary: " && mkdir -p bin && echo "  ===> OK"
	@echo "GO BUILD: " && go build  -o bin/${BIN_NAME} app/pkg/internal/server/*.go && echo " ===> BUILD SUCCESS"

cleanup:
	@echo "clean up binary" \
	&& rm -rf bin/*

docker_build:
	@echo "api gateway docker image build: " \
	&& docker build -t nartvt/${BIN_NAME}_${DATE_NOW} .  \
	&& echo " ==> docker image build success" \
	&& docker scan  --version nartvt/${BIN_NAME}_${DATE_NOW}
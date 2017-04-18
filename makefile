
.PHONY: help all

DEFAULT_GOAL: help

BIN := mytourofgo
BIN_DIR := ./bin
PACKAGE := github.com/workspace7/tourofgo

all: clean build

build: ### Build application binary
	@echo "Building application ${BIN} with sources  ${PACKAGE}" 
	@go build -o ${BIN_DIR}/${BIN} ${PACKAGE}
	@chmod +x ${BIN_DIR}/${BIN}

clean: 	### Clean binary and its directory
	@echo "Cleaning old binaries from ${BIN_DIR}"
	@rm -rf ${BIN_DIR}
	@mkdir -p ${BIN_DIR}

help: 
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


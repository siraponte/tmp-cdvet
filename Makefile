# ---------------------------------------------------------------------------- #
#                  Run make help to see all available commands                 #
# ---------------------------------------------------------------------------- #

.PHONY: help build run dev clean smell pre-commit fmt lint docker-build up down downv logs docs apidocs init githooks install-tools versions

export HELP_MESSAGE
help:
	@printf ${INFO}"$$HELP_MESSAGE"${RST}

# ---------------------------------------------------------------------------- #
#                                    DEVELOP                                   #
# ---------------------------------------------------------------------------- #

MAIN := ./cmd/main.go
OUT := ./bin/cdvet
AIR_CONFIG := ./config/.air.toml

build:
	@go build -o ${OUT} ${MAIN}

run: apidocs
	@go run ${MAIN}

dev:
	@air -c ${AIR_CONFIG}

clean:
	@rm -rf ./bin
	@rm -rf ./openapi

purge: clean downv

# ---------------------------------------------------------------------------- #
#                                  CODE SMELL                                  #
# ---------------------------------------------------------------------------- #

smell: fmt lint

pre-commit:
	@printf ${INFO}"Running pre-commit checks..."${END}
	@./githooks/pre-commit

fmt:
	@printf ${INFO}"Formatting code..."${END}
	@gofmt -w .

lint:
	@printf ${INFO}"Running linter..."${END}
	@golangci-lint run

# ---------------------------------------------------------------------------- #
#                                    DEPLOY                                    #
# ---------------------------------------------------------------------------- #

IMAGE_NAME := ko2dev.azurecr.io/cdvet-be:local
DOCKERFILE := ./deploy/Dockerfile
COMPOSE_CMD := docker compose
COMPOSE_FILE := ./deploy/docker-compose.yaml

docker-build:
	@printf ${INFO}"Building docker container..."${END}
	@docker build -t ${IMAGE_NAME} -f ${DOCKERFILE} .

up:
	@${COMPOSE_CMD} -f ${COMPOSE_FILE} up -d

down:
	@${COMPOSE_CMD} -f ${COMPOSE_FILE} down

downv:
	@${COMPOSE_CMD} -f ${COMPOSE_FILE} down -v

logs:
	@${COMPOSE_CMD} -f ${COMPOSE_FILE} logs -f

# ---------------------------------------------------------------------------- #
#                                 DOCUMENTATION                                #
# ---------------------------------------------------------------------------- #

GODOC_PORT := 6060
OPENAPI_OUT := ./openapi

docs:
	@printf ${SUCCESS}"Documentation available at http://localhost:${GODOC_PORT}/pkg/cdvet/?m=all"${END}
	@printf ${WARNING}"Code under folder 'internal' is visible only if query '?m=all' is appendend at the end of the url as shown in above example"${END}
	@godoc -http=:${GODOC_PORT}

apidocs:
	@printf ${INFO}"Generating API documentation..."${END}
	@swag i -d cmd,app/api -o ${OPENAPI_OUT}

# ---------------------------------------------------------------------------- #
#                                     INIT                                     #
# ---------------------------------------------------------------------------- #

GODOC_VERSION := v0.25.0
SWAGGO_VERSION := v1.16.3
GOLANGCI_LINT_VERSION := v1.61.0
AIR_VERSION := v1.52.3

init: githooks versions install-tools

githooks:
	@printf ${INFO}"Subscribing to githooks..."${END}
	@git config core.hooksPath .githooks
	@printf ${SUCCESS}"Subscribed to githooks"${END}

install-tools:
	@printf ${INFO}"Installing tools..."${END}
	@go install -v golang.org/x/tools/cmd/godoc@${GODOC_VERSION}
	@go install -v github.com/swaggo/swag/cmd/swag@$(SWAGGO_VERSION)
	@go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)
	@go install -v github.com/air-verse/air@$(AIR_VERSION)
	@printf ${SUCCESS}"Tools installed successfully"${END}
	@printf ${WARNING}"Make sure GOBIN (or GOPATH/bin) directory is in your PATH"${END}

versions:
	@printf ${CYN}"Versions of tools used in the project"${END}
	@printf ${TEXT}"godoc: ${GODOC_VERSION}"${END}
	@printf ${TEXT}"swaggo: ${SWAGGO_VERSION}"${END}
	@printf ${TEXT}"golangci-lint: ${GOLANGCI_LINT_VERSION}"${END}
	@printf ${TEXT}"air: ${AIR_VERSION}"${END}


# ---------------------------------------------------------------------------- #
#                                     HELP                                     #
# ---------------------------------------------------------------------------- #

define HELP_MESSAGE
Run "make init" if this is the first time
This will subscribe to githooks and install tools

Usage: make [command]

Available commands:
  Development:
    build         Build the project
    run           Run the project
    dev           Run the project in watch mode
  Code smell:
    smell         Run fmt and lint
    lint          Run linter
    fmt           Format the code
    pre-commit    Run pre-commit checks
  Deploy:
    docker-build  Build docker container
    up            Start docker container
    down          Stop docker container
    downv         Stop docker container and remove volumes
    logs          Show logs of docker container
  Documentation:
    docs          Generate code documentation and start a server
    apidocs       Generate API documentation
  Initialization:
    init          Initialize the project
    githooks      Subscribe to githooks
    install-tools Install tools in GOBIN directory
    versions      Show versions of tools used in the project
  Help:
    help          Show this help message
endef

# ---------------------------------------------------------------------------- #
#                                    COLORS                                    #
# ---------------------------------------------------------------------------- #

BLK="\033[30m"
RED="\033[31m"
GRN="\033[32m"
YLW="\033[33m"
BLU="\033[34m"
MAG="\033[35m"
CYN="\033[36m"
WHT="\033[37m"
RST="\033[0m"
END="\n"${RST}

TEXT=${WHT}
INFO=${CYN}
SUCCESS=${GRN}
ERROR=${RED}
WARNING=${YLW}

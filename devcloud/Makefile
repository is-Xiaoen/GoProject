PKG := "github.com/is-Xiaoen/GoProject/devcloud"
MOD_DIR := $(shell go env GOPATH)/pkg/mod
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/ | grep -v redis)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

GIT_TAG := $()
BUILD_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
BUILD_COMMIT := ${shell git rev-parse HEAD}
BUILD_TIME := ${shell date '+%Y-%m-%d %H:%M:%S'}
BUILD_GO_VERSION := $(shell go version | grep -o  'go[0-9].[0-9].*')
VERSION_PATH := "github.com/infraboard/mcube/v2/ioc/config/application"
OUTPUT_NAME := "devcloud-api"

.PHONY: all dep lint vet test test-coverage build clean

all: build

dep: ## Get the dependencies 依赖
	@go mod tidy

lint: ## Lint Golang files 风格
	@golint -set_exit_status ${PKG_LIST}

vet: ## Run go vet 静态分析
	@go vet ${PKG_LIST}

run: ## Run Devcloud  运行
	@go run main.go start

build: dep ## Build the binary file
	@go build -a -o dist/${OUTPUT_NAME} -ldflags "-s -w" -ldflags "-X '${VERSION_PATH}.GIT_BRANCH=${BUILD_BRANCH}' -X '${VERSION_PATH}.GIT_COMMIT=${BUILD_COMMIT}' -X '${VERSION_PATH}.BUILD_TIME=${BUILD_TIME}' -X '${VERSION_PATH}.GO_VERSION=${BUILD_GO_VERSION}' -X '${VERSION_PATH}.GIT_TAG=${GIT_TAG}'" ${MAIN_FILE}

linux: dep ## Build the linux binary file
	@GOOS=linux GOARCH=amd64 go build -a -o dist/${OUTPUT_NAME} -ldflags "-s -w" -ldflags "-X '${VERSION_PATH}.GIT_BRANCH=${BUILD_BRANCH}' -X '${VERSION_PATH}.GIT_COMMIT=${BUILD_COMMIT}' -X '${VERSION_PATH}.BUILD_TIME=${BUILD_TIME}' -X '${VERSION_PATH}.GO_VERSION=${BUILD_GO_VERSION}' -X '${VERSION_PATH}.GIT_TAG=${GIT_TAG}'" ${MAIN_FILE}

image: dep ## Build the docker image
	docker build -t ${IMAGE_VERSION} -f Dockerfile .

test: ## Run unittests
	@go test -short ${PKG_LIST}

test-coverage: ## Run tests with coverage 覆盖率
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST}
	@cat cover.out >> coverage.txt

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
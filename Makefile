.PHONY: usage build test get-linter lint staticcheck

OK_COLOR=\033[32;01m
NO_COLOR=\033[0m

# Build Flags
BUILD_DATE = $(shell date -u --rfc-3339=seconds)
BUILD_HASH ?= $(shell git rev-parse --short HEAD)
APP_VERSION ?= undefined
BUILD_NUMBER ?= dev

GO := go
GO_LINTER := golint

ECHOFLAGS ?=
ENVFLAGS ?= CGO_ENABLED=0
GOFLAGS ?=
BUILDFLAGS ?= -a -installsuffix cgo $(GOFLAGS) $(GO_LINKER_FLAGS)
EXTLDFLAGS ?= -extldflags "-lm -lstdc++ -static"

BIN_STREAMING_USER := strm-user 

ROOT_DIR := $(realpath .)
CREATE_LOCAL_ENV := $(shell if [ ! -f "$(ROOT_DIR)/local.env" ]; then cp $(ROOT_DIR)/local.env.sample $(ROOT_DIR)/local.env; fi)
LOCAL_VARIABLES ?= $(shell while read -r line; do printf "$$line" | sed 's/ /\\ /g' | awk '{print}'; done < $(ROOT_DIR)/local.env)

PKGS = $(shell $(GO) list ./...)

GO_LINKER_FLAGS ?= --ldflags \
	'$(EXTLDFLAGS) -s -w -X "github.com/streaming-user.BuildNumber=$(BUILD_NUMBER)" \
	-X "github.com/streaming-user.BuildDate=$(BUILD_DATE)" \
	-X "github.com/streaming-user.BuildHash=$(BUILD_HASH)" \
	-X "github.com/streaming-user.AppVersion=$(APP_VERSION)"'

usage: Makefile
	@echo $(ECHOFLAGS) "to use make call:"
	@echo $(ECHOFLAGS) "    make <action>"
	@echo $(ECHOFLAGS) ""
	@echo $(ECHOFLAGS) "list of available actions:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'

## build: build all.
build: lint staticcheck test
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Building binary (linux/amd64/$(BIN_STREAMING_USER))...$(NO_COLOR)"
	@echo $(ECHOFLAGS) $(ENVFLAGS) GOOS=linux GOARCH=amd64 $(GO) build -v $(BUILDFLAGS) -o bin/linux_amd64/$(BIN_STREAMING_USER) ./cmd
	@$(ENVFLAGS) GOOS=linux GOARCH=amd64 $(GO) build -v $(BUILDFLAGS) -o bin/linux_amd64/$(BIN_STREAMING_USER) ./cmd

## test: run unit tests
test:
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Running tests with envs:[$(LOCAL_VARIABLES)]...$(NO_COLOR)"
	@$(ENVFLAGS) $(GO) test $(GOFLAGS) $(PKGS)

## get-linter: install linter
get-linter:
	@echo $(ECHOFLAGS) "$(OK_COLOR")==> Getting linter...$(NO_COLOR)"
	@$(GO) get -v -u golang.org/x/lint/golint

## lint: lint package
lint: get-linter
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Running linter...$(NO_COLOR)"
	@$(GO_LINTER) -set_exit_status $(PKGS)

##staticcheck: run staticcheck on packages
staticcheck:
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Running staticcheck...$(NO_COLOR)"
	@$(GO) get -v honnef.co/go/tools/cmd/staticcheck
	@$(ENVFLAGS) staticcheck $(PKGS)
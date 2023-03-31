APP_NAME = streampay
DAEMON_NAME = streampayd
LEDGER_ENABLED ?= true

PACKAGES=$(shell go list ./... | grep -v '/simulation')
VERSION := $(shell echo $(shell git describe --tags --always) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
COSMOS_SDK := $(shell grep -i cosmos-sdk go.mod | awk '{print $$2}')

build_tags = netgo,
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags+=ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags+=ledger
      endif
    endif
  endif
endif
build_tags := $(strip $(build_tags))

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=${APP_NAME} \
	-X github.com/cosmos/cosmos-sdk/version.AppName=${DAEMON_NAME} \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
	-X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags),cosmos-sdk $(COSMOS_SDK)"

BUILD_FLAGS := -ldflags '$(ldflags)'

all: go.sum install

install: go.sum
		go install $(BUILD_FLAGS) ./cmd/streampayd/
build:
		go build $(BUILD_FLAGS) -o ${GOPATH}/bin/${DAEMON_NAME} ./cmd/streampayd/

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify

reset-and-start-test-chain:
	rm -rf ~/.streampay/config/*
	streampayd unsafe-reset-all
	streampayd init sp-node  --chain-id "sp-test-1"
	streampayd keys add validator --keyring-backend test
	streampayd add-genesis-account `streampayd keys show validator -a --keyring-backend test` 100000000stake
	streampayd gentx validator 1000000stake --moniker "validator-1" --chain-id "sp-test-1" --keyring-backend test
	streampayd collect-gentxs
	streampayd validate-genesis
	streampayd start

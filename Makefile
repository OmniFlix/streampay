APP_NAME = streampay
DAEMON_NAME = streampayd
LEDGER_ENABLED ?= true

BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')

# don't override user values
ifeq (,$(VERSION))
  VERSION := $(shell git describe --exact-match 2>/dev/null)
  # if VERSION is empty, then populate it with branch's name and raw commit hash
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

PACKAGES_SIMTEST=$(shell go list ./... | grep -v '/simulation')
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
BFT_VERSION := $(shell go list -m github.com/cometbft/cometbft | sed 's:.* ::')
DOCKER := $(shell which docker)
BUILDDIR ?= $(CURDIR)/build

GO_SYSTEM_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1-2)
REQUIRE_GO_VERSION = 1.20

export GO111MODULE = on

# process build tags
build_tags = netgo
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

ifeq (cleveldb,$(findstring cleveldb,$(STREAMPAY_BUILD_OPTIONS)))
  build_tags += gcc cleveldb
else ifeq (rocksdb,$(findstring rocksdb,$(STREAMPAY_BUILD_OPTIONS)))
  build_tags += gcc rocksdb
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace := $(whitespace) $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=${APP_NAME} \
	-X github.com/cosmos/cosmos-sdk/version.AppName=${DAEMON_NAME} \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
	-X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
	-X github.com/cometbft/cometbft/version.TMCoreSemVer=$(BFT_VERSION)

ifeq (cleveldb,$(findstring cleveldb,$(STREAMPAY_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
else ifeq (rocksdb,$(findstring rocksdb,$(STREAMPAY_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb
endif
ifeq (,$(findstring nostrip,$(STREAMPAY_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

ifeq (,$(findstring nostrip,$(STREAMPAY_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

all: install
	@echo "--> project root: go mod tidy"
	@go mod tidy
	@echo "--> project root: linting --fix"
	@GOGC=1 golangci-lint run --fix --timeout=8m

install: go.sum
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/streampayd
build:
		go build $(BUILD_FLAGS) -o ${BUILDDIR}/${DAEMON_NAME} ./cmd/streampayd

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify


lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify

reset-and-start-test-chain:
	rm -rf ~/.streampay/config/*
	streampayd tendermint unsafe-reset-all
	streampayd init sp-node  --chain-id "sp-test-1"
	streampayd keys add validator --keyring-backend test
	streampayd genesis add-genesis-account `streampayd keys show validator -a --keyring-backend test` 100000000stake
	streampayd genesis gentx validator 1000000stake --moniker "validator-1" --chain-id "sp-test-1" --keyring-backend test
	streampayd genesis collect-gentxs
	streampayd genesis validate-genesis
	streampayd start

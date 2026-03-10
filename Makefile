VERSION := $(shell echo $(shell git describe --always) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
BUILDDIR ?= $(CURDIR)/build
DOCKER := $(shell which docker)

export GO111MODULE = on

all: lint test-unit

###############################################################################
###                                Build flags                              ###
###############################################################################

# These lines here are essential to include the muslc library for static linking of libraries
# (which is needed for the wasmvm one) available during the build. Without them, the build will fail.
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

# Process linker flags
ldflags = -extldflags "-W"
	

ifeq ($(LINK_STATICALLY),true)
  ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

###############################################################################
###                                  Build                                  ###
###############################################################################
BUILD_TARGETS := build install

build: BUILD_ARGS=-o $(BUILDDIR)

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

$(BUILD_TARGETS): go.sum $(BUILDDIR)/
	GOPRIVATE=github.com/dao-portal go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./...

.PHONY: build install

###############################################################################
###                          Tools & Dependencies                           ###
###############################################################################

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify
	GOPRIVATE=github.com/dao-portal go mod tidy

clean:
	rm -rf $(BUILDDIR)/

.PHONY: go-mod-cache go.sum clean

###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

mockgen_cmd=go.uber.org/mock/mockgen
sql_container_name=dao-portal-test-db
sql_user=dao-portal
sql_password=password
sql_db_name=dao-portal

test-unit:
	@echo "Executing unit tests..."
	@go test -mod=readonly $(BUILD_FLAGS) -v -coverprofile coverage.txt ./...

coverage:
	@echo "Viewing test coverage..."
	@go tool cover --html=coverage.out

stop-test-db:
	@echo "Stopping test database..."
	@sudo docker stop $(sql_container_name) || true && sudo docker rm $(sql_container_name) || true

start-test-db: stop-test-db
	@echo "Starting test database..."
	@sudo docker run --rm --name $(sql_container_name) -e POSTGRES_USER=$(sql_user) -e POSTGRES_PASSWORD=$(sql_password) -e POSTGRES_DB=$(sql_db_name) -d -p 6432:5432 postgres

connect-test-db:
	@echo "Connecting to test database..."
	@sudo docker exec -it $(sql_container_name) psql -U $(sql_user) -d $(sql_db_name)

.PHONY: coverage test-unit

###############################################################################
###                                Codegen                                  ###
###############################################################################
goethereum_abigen_cmd=github.com/ethereum/go-ethereum/cmd/abigen

abigen-governor-bravo:
	@echo "--> Generating governor-bravo contracts bindings"
	@go run $(goethereum_abigen_cmd) \
		--abi ./modules/governor_bravo/contracts/governor_bravo.abi \
		--pkg contracts \
		--type GovernorBravo \
		--out ./modules/governor_bravo/contracts/governor_bravo.go \
		--v2
	@go run $(goethereum_abigen_cmd) \
		--abi ./modules/governor_bravo/contracts/timelock.abi \
		--pkg contracts \
		--type Timelock \
		--out ./modules/governor_bravo/contracts/timelock.go \
		--v2

abigen-oz-governor:
	@echo "--> Generating oz-governor contracts bindings"
	@go run $(goethereum_abigen_cmd) \
		--abi ./modules/oz_governor/contracts/oz_governor.abi \
		--pkg contracts \
		--type OZGovernor \
		--out ./modules/oz_governor/contracts/oz_governor.go \
		--v2

abigen: abigen-governor-bravo abigen-oz-governor

###############################################################################
###                                Linting                                  ###
###############################################################################
golangci_lint_cmd=github.com/golangci/golangci-lint/cmd/golangci-lint
goimports_cmd=golang.org/x/tools/cmd/goimports

lint:
	@echo "--> Running linter"
	@go run $(golangci_lint_cmd) run --timeout=10m

lint-fix:
	@echo "--> Running linter"
	@go run $(golangci_lint_cmd) run --fix --out-format=tab --issues-exit-code=0

format:
	find . -name '*.go' -type f -not -path "*.git*" -not -name '*.pb.go' | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "*.git*" -not -name '*.pb.go' | xargs misspell -w
	find . -name '*.go' -type f -not -path "*.git*" -not -name '*.pb.go' | xargs go run $(goimports_cmd) -w -local github.com/dao-portal/extractor
.PHONY: format

.PHONY: lint lint-fix format

TEST?=./openrtb/...  # vast can't pass go vet and golint right now
GOFMT_FILES?=$$(find . -not -path "./vendor/*" -type f -name '*.go')
PKG_PREFIX=github.com/Vungle/vungo
PKGS=$(PKG_PREFIX)/openrtb $(PKG_PREFIX)/vast

export GO111MODULE=on

default: test

# build
build:
	go build $(PKGS)

# test runs the unit tests
test:
	go test $(TESTARGS) ./...

# generate runs `go generate` to build the dynamically generated source files.
generate:
	find . -name '*_easyjson.go' -delete
	go generate ./...

fmt:
	gofmt -w $(GOFMT_FILES)

lint:
	@warnings=$$(golint $(TEST)); \
	if [[ "$$warnings" ]]; then \
	  echo "golint reports warnings:"; \
	  echo "$$warnings"; \
	  exit 1; \
	else \
	  echo "golint pass"; \
	fi

vet:
	go vet -printfuncs="Verbose,Verbosef,Info,Infof,Warning,Warningf,Error,Errorf,Fatal,Fatalf" $(TEST)

imports:
	@warnings=$$(find . -name "*.go" -not -name "*_easyjson.go" | xargs goimports -l -w); \
	if [[ "$$warnings" ]]; then \
	  echo "goimports reports warnings:"; \
	  echo "$$warnings"; \
	  exit 1; \
	else \
	  echo "goimports no change"; \
	fi

precommit: imports vet lint

# disallow any parallelism (-j) for Make.
.NOTPARALLEL:

.PHONY: build fmt generate imports lint precommit test vet

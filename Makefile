TEST?=./...
GOFMT_FILES?=$$(find . -not -path "./vendor/*" -type f -name '*.go')
PKG_PREFIX=github.com/Vungle/vungo
PKGS=$(PKG_PREFIX)/openrtb $(PKG_PREFIX)/vast

default: test

# build
build:
	go build $(PKGS)

# test runs the unit tests
test:
	go test $(TESTARGS) $(TEST)

# generate runs `go generate` to build the dynamically generated source files.
generate:
	find . -name '*_easyjson.go' -delete
	go generate $(TEST)

fmt:
	gofmt -w $(GOFMT_FILES)

lint:
	golint

# disallow any parallelism (-j) for Make.
.NOTPARALLEL:

.PHONY: build fmt generate lint test

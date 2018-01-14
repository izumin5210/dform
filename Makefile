.DEFAULT_GOAL := all

NAME := dform
VERSION := 0.0.1
REVISION := $(shell git describe --always)
LDFLAGS := -ldflags="-s -w -X \"main.Name=$(NAME)\" -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""
GO_TEST_FLAGS := -v -p=1
COVER_FILE := coverage.txt
COVER_MODE := atomic

GO_PKGS := $(shell go list ./... | grep -v vendor)
SRC_FILES := $(shell git ls-files | grep -E "\.go$$")

BIN := bin/$(NAME)
XC_ARCH := 386 amd64
XC_OS := darwin linux windows


#  App
#-----------------------------------------------
$(BIN): $(SRC_FILES)
	@echo "Build $@"
	@go build $(LDFLAGS) -o $(BIN) main.go

vendor/%: Gopkg.toml Gopkg.lock
	@dep ensure -v -vendor-only


#  Commands
#-----------------------------------------------
.PHONY: all
all: $(BIN)

.PHONY: clean
clean:
	rm -rf bin pkg

.PHONY: clobber
clobber: clean
	rm -rf vendor

.PHONY: dep
dep:
	dep ensure -v -vendor-only

.PHONY: test
test:
	@go test $(GO_TEST_FLAGS) ./...

.PHONY: cover
cover:
	@echo "mode: $(COVER_MODE)" > $(COVER_FILE)
	@set -e; \
	for pkg in $(GO_PKGS); do \
		tmp=/tmp/ro-coverage.out; \
		go test $(GO_TEST_FLAGS) -coverprofile=$$tmp -covermode=$(COVER_MODE) $$pkg; \
		if [ -f $$tmp ]; then \
			cat $$tmp | tail -n +2 >> $(COVER_FILE); \
			rm $$tmp; \
		fi \
	done

.PHONY: lint
lint:
	@gofmt -e -d -s $(SRC_FILES) | awk '{ e = 1; print $0 } END { if (e) exit(1) }'
	@echo $(SRC_FILES) | xargs -n1 golint -set_exit_status
	@go vet ./...

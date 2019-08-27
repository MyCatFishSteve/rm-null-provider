# Program constants
GO = go

GO_BUILD = $(GO) build
GO_BUILD_ARGS = -buildmode=plugin -v
GO_TEST = $(GO) test
GO_TEST_ARGS = -v

BINPATH = bin
BIN = rm-null-resource.so

INSTALLPATH = ${HOME}/.resource-manager.d/plugins

all: test build

test:
	$(GO_TEST) $(GO_TEST_ARGS) ./...

build:
	$(GO_BUILD) $(GO_BUILD_ARGS) -o $(BINPATH)/$(BIN)

install:
	mv "$(BINPATH)/$(BIN)" "$(INSTALLPATH)/"

clean:
	rm -r bin/

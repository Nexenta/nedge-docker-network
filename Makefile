NDNET_EXE = ndnet
FLAGS = -v

NEDGE_DEST = /opt/nedge/sbin
NEDGE_ETC = /opt/nedge/etc/ccow

GO_VERSION = 1.6
GO_INSTALL = /usr/lib/go-$(GO_VERSION)
GO = $(GO_INSTALL)/bin/go

all: $(NDNET_EXE)

GO_FILES = src/github.com/Nexenta/nedge-docker-network/ndnet/ndnet.go \
	src/github.com/Nexenta/nedge-docker-network/ndnet/ndnetcli/ndnetcli.go \
	src/github.com/Nexenta/nedge-docker-network/ndnet/ndnetcli/daemoncli.go \
	src/github.com/Nexenta/nedge-docker-network/ndnet/ndnetcli/volumecli.go \
	src/github.com/Nexenta/nedge-docker-network/ndnet/daemon/daemon.go \
	src/github.com/Nexenta/nedge-docker-network/ndnet/daemon/driver.go \
	src/github.com/Nexenta/nedge-docker-network/ndnet/ndnetapi/ndnetapi.go

$(GO_FILES): setup

deps: setup
	GOPATH=$(shell pwd) GOROOT=$(GO_INSTALL) $(GO) get github.com/docker/go-plugins-helpers/volume
	GOPATH=$(shell pwd) GOROOT=$(GO_INSTALL) $(GO) get github.com/docker/go-plugins-helpers/network
	GOPATH=$(shell pwd) GOROOT=$(GO_INSTALL) $(GO) get github.com/codegangsta/cli
	GOPATH=$(shell pwd) GOROOT=$(GO_INSTALL) $(GO) get github.com/Sirupsen/logrus
	GOPATH=$(shell pwd) GOROOT=$(GO_INSTALL) $(GO) get github.com/coreos/go-systemd/util
	GOPATH=$(shell pwd) GOROOT=$(GO_INSTALL) $(GO) get github.com/opencontainers/runc/libcontainer/user
	GOPATH=$(shell pwd) GOROOT=$(GO_INSTALL) $(GO) get golang.org/x/net/proxy


$(NDNET_EXE): $(GO_FILES)
	GOPATH=$(shell pwd) GOROOT=$(GO_INSTALL) $(GO) install github.com/Nexenta/nedge-docker-network/ndnet

install: $(NDNET_EXE)
	cp -f bin/$(NDNET_EXE) $(NEDGE_DEST)
	cp -f src/github.com/Nexenta/nedge-docker-network/ndnet/daemon/ndnet.json $(NEDGE_ETC)
uninstall:
	rm -f $(NEDGE_DEST)/$(NDNET_EXE)
	rm -f $(NEDGE_ETC)/ndnet.json
build:
	GOPATH=$(shell pwd) GOROOT=$(GO_INSTALL) $(GO) build $(FLAGS) github.com/Nexenta/nedge-docker-network/ndnet

setup: 
	mkdir -p src/github.com/Nexenta/nedge-docker-network/ 
	cp -R ndnet/ src/github.com/Nexenta/nedge-docker-network/ndnet 

lint:
	GOPATH=$(shell pwd) GOROOT=$(GO_INSTALL) $(GO) get -v github.com/golang/lint/golint
	for file in $$(find . -name '*.go' | grep -v vendor | grep -v '\.pb\.go' | grep -v '\.pb\.gw\.go'); do \
		golint $${file}; \
		if [ -n "$$(golint $${file})" ]; then \
			exit 1; \
		fi; \
	done

clean:
	GOPATH=$(shell pwd) GOROOT=$(GO_INSTALL) $(GO) clean

clobber:
	rm -rf src/github.com/Nexenta/nedge-docker-network
	rm -rf bin/ pkg/


NDNET_EXE = ndnet
FLAGS = -v

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
	GOPATH=$(shell pwd) go get github.com/docker/go-plugins-helpers/volume
	GOPATH=$(shell pwd) go get github.com/docker/go-plugins-helpers/network
	GOPATH=$(shell pwd) go get github.com/codegangsta/cli
	GOPATH=$(shell pwd) go get github.com/Sirupsen/logrus
	GOPATH=$(shell pwd) go get github.com/coreos/go-systemd/util
	GOPATH=$(shell pwd) go get github.com/opencontainers/runc/libcontainer/user
	GOPATH=$(shell pwd) go get golang.org/x/net/proxy


$(NDNET_EXE): $(GO_FILES)
	GOPATH=$(shell pwd) go install github.com/Nexenta/nedge-docker-network/ndnet

build:
	GOPATH=$(shell pwd) go build $(FLAGS) github.com/Nexenta/nedge-docker-network/ndnet

setup: 
	mkdir -p src/github.com/Nexenta/nedge-docker-network/ 
	cp -R ndnet/ src/github.com/Nexenta/nedge-docker-network/ndnet 

lint:
	GOPATH=$(shell pwd) go get -v github.com/golang/lint/golint
	for file in $$(find . -name '*.go' | grep -v vendor | grep -v '\.pb\.go' | grep -v '\.pb\.gw\.go'); do \
		golint $${file}; \
		if [ -n "$$(golint $${file})" ]; then \
			exit 1; \
		fi; \
	done

clean:
	GOPATH=$(shell pwd) go clean


clobber:
	rm -rf src/github.com/Nexenta/nedge-docker-network
	rm -rf bin/ pkg/


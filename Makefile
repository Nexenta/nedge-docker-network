NEDGE_DEST = $(DESTDIR)/opt/nedge/sbin
NEDGE_ETC = $(DESTDIR)/opt/nedge/etc/ccow
NDNET_EXE = ndnet

build: 
	GOPATH=$(shell pwd) go get -v github.com/docker/go-plugins-helpers/network
	cd src/github.com/docker/go-plugins-helpers/network; git checkout d7fc7d0
	cd src/github.com/docker/go-connections; git checkout acbe915
	GOPATH=$(shell pwd) go get -v github.com/Nexenta/nedge-docker-network/...

lint:
	GOPATH=$(shell pwd) GOROOT=$(GO_INSTALL) $(GO) get -v github.com/golang/lint/golint
	for file in $$(find . -name '*.go' | grep -v vendor | grep -v '\.pb\.go' | grep -v '\.pb\.gw\.go'); do \
		golint $${file}; \
		if [ -n "$$(golint $${file})" ]; then \
			exit 1; \
		fi; \
	done

install:
	cp -n ndnet/daemon/ndnet.json $(NEDGE_ETC)
	cp -f bin/$(NDNET_EXE) $(NEDGE_DEST)

uninstall:
	rm -f $(NEDGE_ETC)/ndnet.json
	rm -f $(NEDGE_DEST)/ndnet

clean:
	go clean github.com/Nexenta/nedge-docker-network

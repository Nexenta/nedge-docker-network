NEDGE_DEST = $(DESTDIR)/opt/nedge/sbin
NEDGE_ETC = $(DESTDIR)/opt/nedge/etc/ccow
NDNET_EXE = ndnet

build: 
	GOPATH=$(shell pwd) go get -v github.com/docker/go-plugins-helpers/network
	cd src/github.com/docker/go-plugins-helpers/network; git checkout a9ef19c479cb60e751efa55f7f2b265776af1abf
	cd src/github.com/docker/go-connections; git checkout 3ede32e2033de7505e6500d6c868c2b9ed9f169d
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

NEDGE_DEST = $(DESTDIR)/opt/nedge/sbin
NEDGE_ETC = $(DESTDIR)/opt/nedge/etc/ccow
NDNET_EXE = ndnet

build: 
	GOPATH=$(shell pwd) go get -v github.com/Nexenta/nedge-docker-network/...
	cp ndnet/daemon/ndnet.json $(NEDGE_ETC)
	cp -f bin/$(NDNET_EXE) $(NEDGE_DEST)

lint:
	GOPATH=$(shell pwd) GOROOT=$(GO_INSTALL) $(GO) get -v github.com/golang/lint/golint
	for file in $$(find . -name '*.go' | grep -v vendor | grep -v '\.pb\.go' | grep -v '\.pb\.gw\.go'); do \
		golint $${file}; \
		if [ -n "$$(golint $${file})" ]; then \
			exit 1; \
		fi; \
	done

install:
	cp -f bin/$(NDNET_EXE) $(NEDGE_DEST)

uninstall:
	rm -f $(NEDGE_ETC)/ndnet.json
	rm -f $(NEDGE_DEST)/ndnet

clean:
	go clean github.com/Nexenta/nedge-docker-network

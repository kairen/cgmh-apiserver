VERSION_MAJOR ?= 0
VERSION_MINOR ?= 3
VERSION_BUILD ?= 0
VERSION ?= v$(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_BUILD)

GOOS ?= $(shell go env GOOS)

OWNER := inwinstack
GROUP := cgmh
REPOPATH ?= $(OWNER)/$(GROUP)/apiserver

$(shell mkdir -p ./out)

.PHONY: build
build: out/server

.PHONY: out/server
out/server:
	GOOS=$(GOOS) go build \
	  -ldflags="-X $(REPOPATH)/pkg/version.version=$(VERSION)" \
	  -a -o $@ cmd/main.go

.PHONY: dep 
dep:
	dep ensure

.PHONY: test
test:
	./hack/test-go.sh

.PHONY: build_image
build_image:
	docker build -t registry.gitlab.com/$(REPOPATH):$(VERSION) .

.PHONY: push_image
push_image:
	docker push registry.gitlab.com/$(REPOPATH):$(VERSION)

.PHONY: start_server
start_server: 
	./hack/start-server.sh

.PHONY: clean
clean:
	rm -rf out/


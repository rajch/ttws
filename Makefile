# Bump these on release, and for now update the deployment files
VERSION_MAJOR ?= 0
VERSION_MINOR ?= 1
BUILD_NUMBER  ?= 0

IMAGE_TAG ?= $(VERSION_MAJOR).$(VERSION_MINOR).$(BUILD_NUMBER)
REGISTRY_USER ?= rajchaudhuri

.PHONY: all
all: tinyws dockerimage

out/tinyws: cmd/main.go internal/pkg/webserver/*.go
	CGO_ENABLED=0 go build -o $@ $<

.PHONY: tinyws
tinyws: out/tinyws

.PHONY: dockerimage
dockerimage: build/package/singlestage.Dockerfile out/tinyws
	docker image build -f build/package/singlestage.Dockerfile -t $(REGISTRY_USER)/tinyws:$(IMAGE_TAG) .

.PHONY: rmi
rmi:
	docker image rm $(REGISTRY_USER)/tinyws:$(IMAGE_TAG)

.PHONY: clean
clean:
	rm -r out/
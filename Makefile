# Bump these on release, and for now update the deployment files
VERSION_MAJOR ?= 0
VERSION_MINOR ?= 2
BUILD_NUMBER  ?= 1

IMAGE_TAG ?= $(VERSION_MAJOR).$(VERSION_MINOR).$(BUILD_NUMBER)
REGISTRY_USER ?= rajchaudhuri

# Macro for module meta-ruleset. Sample generated ruleset with input 'webserver': 
# .PHONY: webserver
# webserver: pkg/webserver/*.go
define MODULERULE
.PHONY: $M
$M: pkg/$(M)/*.go
endef

# Macro for command meta-ruleset. Sample generated ruleset with input 'ttws':
# .PHONY: ttws
# ttws: out/ttws
#
# out/ttws: cmd/ttws/main.go $(ttwsMODULES)
# 	CGO_ENABLED=0 go build -o out/ttws cmd/ttws/main.go
#
# .PHONY: ttwsimage
# ttwsimage: ttws out/ttwsDockerfile 
# 	docker image build -f out/ttwsDockerfile -t $(REGISTRY_USER)/ttws:$(IMAGE_TAG) out/
#
# out/ttwsDockerfile: build/package/ttws/singlestage.Dockerfile
# 	cp build/package/ttws/singlestage.Dockerfile out/ttwsDockerfile
#
# .PHONY: ttwsimagemultistage
#	docker image build -f build/package/ttws/multistage.Dockerfile -t $(REGISTRY_USER)/ttws:$(IMAGE_TAG) .
#
# .PHONY: ttwsrmi
# ttwsrmi:
# 	docker image rm $(REGISTRY_USER)/ttws:$(IMAGE_TAG)
define CMDRULE
.PHONY: $C
$C: out/$(C)

out/$C: cmd/$C/main.go $($(C)MODULES)
	CGO_ENABLED=0 go build -o out/$(C) -ldflags "-X 'github.com/rajch/ttws/pkg/webserver.version=v$(IMAGE_TAG)'" cmd/$C/main.go

.PHONY: $(C)image
$(C)image: $C out/$(C)Dockerfile $(C)webassets
	docker image build -f out/$(C)Dockerfile \
					   -t $(REGISTRY_USER)/$(C):$(IMAGE_TAG) \
					   out/

out/$(C)Dockerfile: build/package/$(C)/singlestage.Dockerfile
	cp build/package/$(C)/singlestage.Dockerfile out/$(C)Dockerfile

.PHONY: $(C)webassets
$(C)webassets:
	@if [ -d web/$(C) ]; then cp -r web/$(C)/* out/; else echo No web assets for $(C);fi

.PHONY: $(C)imagemultistage
$(C)imagemultistage: build/package/$(C)/multistage.Dockerfile
	docker image build -f build/package/$(C)/multistage.Dockerfile \
					   -t $(REGISTRY_USER)/$(C):$(IMAGE_TAG) \
					   --build-arg IMAGE_TAG=$(IMAGE_TAG) \
					   .

.PHONY: rmi$(C)
rmi$(C):
	docker image rm $(REGISTRY_USER)/$(C):$(IMAGE_TAG)

.PHONY: $(C)imagemultiarch
$(C)imagemultiarch: build/package/$(C)/multistage.Dockerfile
	docker buildx build -f build/package/$(C)/multistage.Dockerfile \
						-t $(REGISTRY_USER)/$(C):$(IMAGE_TAG) \
						--build-arg IMAGE_TAG=$(IMAGE_TAG) \
						--platform linux/amd64,linux/386,linux/arm64,linux/ppc64le,linux/arm/v7 \
						--push \
						.

endef

ALLMODULES = webserver cpuload ipaddresses envvars filesystem probes static
ALLCMDS = ttws ics ldgen probestest

ttwsMODULES = $(ALLMODULES)
icsMODULES = webserver ipaddresses envvars filesystem
ldgenMODULES= webserver cpuload
probestestMODULES= webserver probes

# .PHONY: all
all: list

# Apply command rules
$(foreach C,$(ALLCMDS),$(eval $(CMDRULE)))

# Apply module rules
$(foreach M,$(ALLMODULES),$(eval $(MODULERULE)))

.PHONY: clean
clean:
	rm -r out/

.PHONY: list
list:
	@echo Available targets: $(ALLCMDS)

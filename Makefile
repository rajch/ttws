# Bump these on release, and for now update the deployment files
VERSION_MAJOR ?= 0
VERSION_MINOR ?= 1
BUILD_NUMBER  ?= 0

IMAGE_TAG ?= $(VERSION_MAJOR).$(VERSION_MINOR).$(BUILD_NUMBER)
REGISTRY_USER ?= rajchaudhuri

# Macro for meta-rule
define MODULERULE
.PHONY: $M
$M: pkg/$(M)/*.go
endef

ALLMODULES = webserver cpuload ipaddresses envvars

.PHONY: all
all: ttws

.PHONY: ttws
ttws: out/ttws

out/ttws: cmd/ttws/main.go $(ALLMODULES)
	CGO_ENABLED=0 go build -o $@ $<

# The following meta-rule will generate rules like:
# .PHONY: webserver
# webserver: pkg/webserver/*.go
# .PHONY: cpuload
# cpuload: pkg/cpuload/*.go
# .PHONY: ipaddresses
# ipaddresses: pkg/ipaddresses/*.go
# .PHONY: envvars
# envvars: pkg/envvars/*.go
$(foreach M,$(ALLMODULES),$(eval $(MODULERULE)))

.PHONY: ttwsimage
ttwsimage: ttws out/ttwsDockerfile 
	docker image build -f out/ttwsDockerfile -t $(REGISTRY_USER)/ttws:$(IMAGE_TAG) out/

out/ttwsDockerfile: build/package/ttws/singlestage.Dockerfile
	cp $< $@

.PHONY: rmi
rmi:
	docker image rm $(REGISTRY_USER)/ttws:$(IMAGE_TAG)

.PHONY: clean
clean:
	rm -r out/
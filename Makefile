
GO ?= go

VERSION ?= $(shell cat ./VERSION)

LDFLAGS_COMMON = -X main.htmx_version=$(VERSION)
GO_BUILD := $(GO) build $(EXTRA_FLAGS) -ldflags "$(LDFLAGS_COMMON)"

.DEFAULT: all

.PHONY: all
%/main: %/main.go
	$(GO_BUILD) -o $@ $@.go

.PHONY: all
all: \
	click-to-edit/main click-to-load/main

.PHONY: all
clean:
	rm -rf */main

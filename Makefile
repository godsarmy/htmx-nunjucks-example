
GO ?= go

HTMX_VERSION ?= $(shell cat ./HTMX_VERSION)
NUNJUCKS_VERSION ?= $(shell cat ./NUNJUCKS_VERSION)

LDFLAGS_COMMON = -X main.htmx_version=$(HTMX_VERSION) -X main.nunjucks_version=$(NUNJUCKS_VERSION)
GO_BUILD := $(GO) build $(EXTRA_FLAGS) -ldflags "$(LDFLAGS_COMMON)"

.DEFAULT: all

.PHONY: all
%/main: %/main.go
	$(GO_BUILD) -o $@ $@.go

.PHONY: all
all: \
	active-search/main \
	animations/main \
	bulk-update/main \
	click-to-edit/main \
	click-to-load/main \
	delete-row/main

.PHONY: all
clean:
	rm -rf */main

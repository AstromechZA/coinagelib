.DEFAULT: help

# Display a list of the documented make targets
.PHONY: help
help:
	@echo Documented Make targets:
	@perl -e 'undef $$/; while (<>) { while ($$_ =~ /## (.*?)(?:\n# .*)*\n.PHONY:\s+(\S+).*/mg) { printf "\033[36m%-30s\033[0m %s\n", $$2, $$1 } }' $(MAKEFILE_LIST) | sort

GO_PACKAGES := $(shell go list ./...)
GO_PACKAGE_PATHS := $(shell go list -f '{{.Dir}}' $(GO_PACKAGES))
GO_FILES := $(shell find $(GO_PACKAGE_PATHS) -type f -name *.go)

.PHONY: test
test:
	@echo $(shell date): Running go test..
	@go test -v -vet= ./...
	@echo

	@echo $(shell date): Running go fmt to check formatting..
	@b=$(shell gofmt -l $(GO_FILES)); if [ ! -z "$$b" ]; then \
		echo $$(date): found $$(echo $$b | wc -l) issues in the following files:; \
		echo $$b; \
		echo; \
		false; \
	fi

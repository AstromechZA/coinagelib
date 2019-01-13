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

coverage.out: $(GO_FILES)
	@echo $(shell date): Analysing test coverage..
	@go test -cover -covermode=count -coverprofile=$@ ./...

.PHONY: coverage
coverage: coverage.out

.PHONY: coveralls
coveralls: coverage.out
	@if [ -z "$$(which goveralls)" ]; then \
		echo $$(date): Installing goveralls because we couldnt find it on the path..; \
		GO111MODULE=off go get github.com/mattn/goveralls; \
	fi
	@if [ -z "$(COVERALLS_TOKEN)" ]; then \
		echo You must set COVERALLS_TOKEN variable!; \
		false; \
	fi
	goveralls -coverprofile=coverage.out -service=circleci -repotoken $(COVERALLS_TOKEN)

.PHONY: clean
clean: PATTERNS = coverage.out
clean:
	@echo $(shell date): Removed $(shell rm -rfv $(PATTERNS) | wc -l) files..

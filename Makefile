BINARIES :=
include base.Makefile

## Send coveraalls report
.PHONY: coveralls
coveralls: artifacts/coverage.out
	@if [ -z "$$(which goveralls)" ]; then \
		echo $$(date): Installing goveralls because we couldnt find it on the path..; \
		GO111MODULE=off go get github.com/mattn/goveralls; \
	fi
	@if [ -z "$(COVERALLS_TOKEN)" ]; then \
		echo You must set COVERALLS_TOKEN variable!; \
		false; \
	fi
	goveralls -coverprofile=coverage.out -service=circleci -repotoken $(COVERALLS_TOKEN)

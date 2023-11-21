.PHONY: test

GOLANGLINT_INSTALLED_VERSION := $(shell golangci-lint version 2>/dev/null | sed -ne 's/.*version\ \([0-9]*\.[0-9]*\.[0-9]*\).*/\1/p')
GOLANG_LINT_VERSION := 1.55.2

REPO := github.com/pedrokiefer/awspol

COMMIT_ID := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(date -u +"%Y-%m-%dT%H:%M:%SZ")
COMMIT_TAG := $(shell git describe --tags --always --abbrev=0 --match="v[0-9]*.[0-9]*.[0-9]*" 2> /dev/null)
VERSION := $(shell echo "${COMMIT_TAG}" | sed 's/^.//')

GO_LDFLAGS := -ldflags "-X '${REPO}/cmd.Version=${VERSION}' -X '${REPO}/cmd.CommitID=${COMMIT_ID}' -X '${REPO}/cmd.BuildTime=${BUILD_TIME}'"

test:
	go test ./... -cover

lint:
ifneq (${GOLANG_LINT_VERSION}, ${GOLANGLINT_INSTALLED_VERSION})
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v${GOLANG_LINT_VERSION}
endif
	golangci-lint run

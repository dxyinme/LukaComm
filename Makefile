PROJECT=LukaComm_test
GOPATH ?= $(shell go env GOPATH)
CURDIR := $(shell pwd)

FAIL_ON_STDOUT := awk '{ print } END { if (NR > 0) { exit 1 } }'

GO        := GO111MODULE=on go
GOBUILD   := $(GO) build
GOTEST    := $(GO) test -p 4

FILES     := $$(find . -name "*.go")

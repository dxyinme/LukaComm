PROJECT=LukaComm_test
GOPATH ?= $(shell go env GOPATH)
CURDIR := $(shell pwd)

FAIL_ON_STDOUT := awk '{ print } END { if (NR > 0) { exit 1 } }'

GO        := GO111MODULE=on go
GOBUILD   := $(GO) build
GOTEST    := $(GO) test -p 4

FILES     := $$(find . -name "*.go")

fmt:
	@echo "gofmt (simplify)"
	@gofmt -s -l -w $(FILES) 2>&1 | $(FAIL_ON_STDOUT)

cynicu-test:
	@echo "go CynicU-test generate"
	$(GOBUILD) -o bin/CynicU_server main/CynicUServer.go
	$(GOBUILD) -o bin/CynicU_client main/CynicUClient.go
	$(GOBUILD) -o bin/CynicU_send_server main/CynicU-sendServer.go
	$(GOBUILD) -o bin/CynicU_send_client main/CynicU-sendClient.go
	@cp -r test_script/CynicU-test bin/

clear:
	@rm -rf bin/
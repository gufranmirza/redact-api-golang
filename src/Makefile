export PROJECT = github.com/gufranmirza/redact-api-golang
export BINARY_NAME = "redact-api"
export CONFIG_FILE = "app-config.json"

# global command
.PHONY: all
all: dependencies build test

.PHONY: dependencies
dependencies:
	go mod download

.PHONY: build
build:
	env GOOS=linux GOARCH=amd64 go build -o ../bin/${BINARY_NAME} $(PROJECT)/src
	chmod +x ../bin/${BINARY_NAME}
	cp -r ${CONFIG_FILE} ../bin

.PHONY: run
run:
	go run ./main.go

.PHONY: start
start:
	../bin/${BINARY_NAME}

.PHONY: mock
mock:	
	mockgen -package servermock github.com/gufranmirza/redact-api-golang/src/web/server Server > web/server/servermock/mock.go
	mockgen -package routermock github.com/gufranmirza/redact-api-golang/src/web/router Router > web/router/routermock/mock.go
	mockgen -package redactmock github.com/gufranmirza/redact-api-golang/src/web/services/redact Redact > web/services/redact/redactmock/mock.go
	mockgen -package healthmock github.com/gufranmirza/redact-api-golang/src/web/services/health Health > web/services/health/healthmock/mock.go

packages = \
	./config\
	./web/server\
	./web/router\
	./web/services/health\
	./web/services/redact\

.PHONY: test
test:
	@$(foreach package,$(packages), \
    		set -e; \
    		go test  $(package);)

export PROJECT = github.com/gufranmirza/redact-api-golang
export BINARY_NAME = "redact-api"

build:
	env GOOS=linux GOARCH=amd64 go build -o ../bin/${BINARY_NAME} $(PROJECT)/src
	chmod +x ../bin/${BINARY_NAME}

run:
	go run ./main.go

start:
	../bin/${BINARY_NAME}

mock:	
	mockgen -package healthmock github.com/gufranmirza/redact-api-golang/src/web/services/health Health > web/services/health/healthmock/mock.go
	mockgen -package servermock github.com/gufranmirza/redact-api-golang/src/web/server Server > web/server/servermock/mock.go
	mockgen -package routermock github.com/gufranmirza/redact-api-golang/src/web/router Router > web/router/routermock/mock.go

packages = \
	./config\
	./web/server\
	./web/router\
	./web/services/health\

test:
	@$(foreach package,$(packages), \
    		set -e; \
    		go test  $(package);)
